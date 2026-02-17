package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

// name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var (
	// awgo specific variable
	wf *aw.Workflow // Our Workflow object

	// base variables
	query         string
	baseURL       = "https://www.dnddeutsch.de/tools/json.php?apiv=0.7&o=monster&q="
	repo          = "Wayneoween/alfred-dndbeyond-monster-workflow"
	helpURL       = "https://github.com/" + repo
	maxResults    = 20
	doTranslateDE bool

	// possible commandline flags
	doCheck     bool
	doTranslate bool

	// cache variables
	cacheName   = "cache.json"              // Filename of cached repo list
	maxCacheAge = 7 * 24 * 60 * time.Minute // Cache each query for 14 days
)

func init() {
	// create a new *Workflow using default configuration
	// (workflow settings are read from the environment variables set by Alfred)
	wf = aw.New(
		aw.HelpURL(helpURL),
		aw.MaxResults(maxResults),
		update.GitHub(repo),
	)

	// add a commandline flag to the binary so that updateCheck can call it
	flag.BoolVar(&doCheck, "check", false, "check for a new version")
	// add a commandline flag to set the translation toggle
	flag.BoolVar(&doTranslate, "translate", false, "toggle german translation")
}

func toggleTranslate() {
	doTranslateDE = wf.Config.GetBool("translate", false)
	log.Println("DEBUG: doTranslateDE=" + strconv.FormatBool(doTranslateDE))

	if doTranslate {
		log.Printf("Toggled translate from %s to %s.", strconv.FormatBool(doTranslateDE), strconv.FormatBool(!doTranslateDE))
		wf.Configure(aw.TextErrors(true))
		wf.NewItem(fmt.Sprintf("Set %s to “%s”", "translate", strconv.FormatBool(doTranslateDE))).
			Subtitle("↩ to save").
			Arg(strconv.FormatBool(doTranslateDE)).
			Valid(true).
			Var("value", strconv.FormatBool(doTranslateDE)).
			Var("varname", "translate")

		if err := wf.Config.Set("translate", strconv.FormatBool(!doTranslateDE), false, wf.BundleID()).Do(); err != nil {
			wf.FatalError(err)
		}

		return
	}
}

func run() {
	log.Println("DEBUG: Function 'run'!")

	// call to handle any magic actions
	wf.Args()
	flag.Parse()

	// handle the translation setting
	toggleTranslate()

	// collect the first word as only argument
	if args := wf.Args(); len(args) > 0 {
		query = args[0]
	}

	log.Printf("[main] query=%s", query)
	monsters := []*Monster{}

	// try to load cached monsters from $query_$cachename.json
	if wf.Cache.Exists(query + "_" + cacheName) {
		log.Println("DEBUG: data is being loaded from cache.")
		if err := wf.Cache.LoadJSON(query+"_"+cacheName, &monsters); err != nil {
			wf.FatalError(err)
		}
		log.Println("DEBUG: monsters loaded: ", len(monsters))
	}

	if wf.Cache.Expired(strings.Replace(query, " ", "-", -1)+"_"+cacheName, maxCacheAge) {
		log.Println("DEBUG: data is being loaded from website.")

		var resultSet D3ResultSet

		log.Println("DEBUG: loading data from " + baseURL + query)
		response, err := http.Get(baseURL + query)
		if err != nil {
			fmt.Print(err.Error())
		}

		defer response.Body.Close()

		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("DEBUG: responseData: %s", responseData)

		// unmarshall the JSON response into resultSet
		json.Unmarshal(responseData, &resultSet)

		log.Printf("DEBUG: unmarshalled data:")
		log.Println(resultSet)

		// if there are results in the monsterarray for our query
		if len(resultSet.Monster) > 0 {
			log.Printf("DEBUG: %d Monsters were found!", len(resultSet.Monster))

			// range over the array and create entries for every one of them
			log.Println("DEBUG: printing each monster:")
			for _, result := range resultSet.Monster {

				name := result.NameEN
				if len(name) == 0 {
					name = result.NameDE
				}

				if len(result.Src) > 0 {
					if containsAny(excludeSrc, result.Src) {
						log.Printf("DEBUG: skipped monster %s because of excluded source %s\n", name, result.Src)
						continue
					}
				}

				if len(result.Size) > 0 && len(result.Type) > 0 {
					// add the result fields to temp Monster
					temp := result

					// only add the monster if the query string is contained in the name
					// the API also returns entries where the query string is contained
					// within the monster description and that is mixed german and english
					if !strings.Contains(strings.ToLower(name), strings.ToLower(query)) {
						log.Println("DEBUG: Monster name not containing search string. Skipping.")
						log.Println(result)
						continue
					}

					log.Println("MonsterName:   ", name)
					log.Println("MonsterType:   ", temp.Type)
					log.Println("MonsterSize:   ", temp.Size)
					log.Println("MonsterCR:     ", temp.Cr)
					log.Println("-------------------------------------------------")

					monsters = append(monsters, &temp)
				} else {
					log.Println("DEBUG: not a real monster entry. Skipping.")
					log.Println(result)
				}
			}
		}

		// write cache only if we have at least one monster
		if len(monsters) != 0 {
			log.Println("DEBUG: more than 1 monsters found. Caching...")
			wf.Configure(aw.TextErrors(true))
			if err := wf.Cache.StoreJSON(strings.Replace(query, " ", "-", -1)+"_"+cacheName, monsters); err != nil {
				wf.FatalError(err)
			}
		}
	}

	// if there are no monsters just send the warning.
	if len(monsters) == 0 {
		if doTranslateDE {
			wf.WarnEmpty("Leider nichts gefunden.", "Versuche einen anderen Namen.")
		} else {
			wf.WarnEmpty("Nothing found.", "Try another name.")
		}
	} else {
		// no matter if via internet or from the cache, add all monsters as items for alfred
		for _, temp := range monsters {

			if doTranslateDE {
				var titleDE string
				var subtitleDE string
				// if DE name and EN name are the same OR if DE name does not exist, use EN name
				if temp.NameDE == temp.NameEN || len(temp.NameDE) == 0 {
					titleDE = temp.NameEN
				} else {
					titleDE = fmt.Sprintf("%s - %s", temp.NameDE, temp.NameEN)
				}
				if temp.PageDE == "0" {
					subtitleDE = fmt.Sprintf("CR %s - %s - %s - %s %s(en)", temp.Cr, temp.Size, temp.Type, temp.Src, temp.PageEN)
				} else {
					subtitleDE = fmt.Sprintf("CR %s - %s - %s - %s %s(de) %s(en)", temp.Cr, temp.Size, temp.Type, temp.Src, temp.PageDE, temp.PageEN)
				}
				wf.NewItem(titleDE).
					Subtitle(subtitleDE).
					Icon(getIconForType(temp.Type)).
					Arg("https://www.dndbeyond.com/monsters/" + strings.Replace(strings.ToLower(temp.NameEN), " ", "-", -1)).
					UID(temp.NameEN + temp.SingleLine).
					Valid(true)
			} else {
				// But Ogrillon or Kyton as well as "in lair" is still a problem.
				// TODO: a function that removes special cases here.
				tmpTitle := temp.NameEN
				tmpTitle = strings.ReplaceAll(tmpTitle, "(", "")
				tmpTitle = strings.ReplaceAll(tmpTitle, ")", "")
				titleEN := tmpTitle
				subtitleEN := fmt.Sprintf("CR %s - %s - %s - %s %s", temp.Cr, temp.Size, temp.Type, temp.Src, temp.PageEN)
				wf.NewItem(titleEN).
					Subtitle(subtitleEN).
					Icon(getIconForType(temp.Type)).
					Arg("https://www.dndbeyond.com/monsters/" + strings.Replace(strings.ToLower(temp.NameEN), " ", "-", -1)).
					UID(temp.NameEN + temp.SingleLine).
					Valid(true)
			}
		}
	}

	if doCheck {
		wf.Configure(aw.TextErrors(true))
		log.Println("Checking for updates...")
		if err := wf.CheckForUpdate(); err != nil {
			wf.FatalError(err)
		}

		return
	}

	if wf.UpdateCheckDue() && !wf.IsRunning(updateJobName) {
		log.Println("Running update check in background...")

		cmd := exec.Command(os.Args[0], "-check")
		if err := wf.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("Error starting update check: %s", err)
		}
	}

	// Only show update status if query is empty.
	if query == "" && wf.UpdateAvailable() {
		// Turn off UIDs to force this item to the top.
		// If UIDs are enabled, Alfred will apply its "knowledge"
		// to order the results based on your past usage.
		wf.Configure(aw.SuppressUIDs(true))

		// Notify user of update. As this item is invalid (Valid(false)),
		// actioning it expands the query to the Autocomplete value.
		// "workflow:update" triggers the updater Magic Action that
		// is automatically registered when you configure Workflow with
		// an Updater.
		//
		// If executed, the Magic Action downloads the latest version
		// of the workflow and asks Alfred to install it.
		wf.NewItem("Update available!").
			Subtitle("↩ to install").
			Autocomplete("workflow:update").
			Valid(false).
			Icon(UpdateAvailable)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
	fmt.Println("")
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

// Name of the background job that checks for updates
const updateJobName = "checkForUpdate"

// D3ResultSet mirrors the response from the dnddeutsch.de API
type D3ResultSet struct {
	O       string    `json:"o"`
	V       float64   `json:"v"`
	Monster []Monster `json:"monster"`
}

// Monster mirrors the response from the dnddeutsch.de API
type Monster struct {
	NameDE        string   `json:"name_de"`
	NameDEUlisses string   `json:"name_de_ulisses"`
	NameEN        string   `json:"name_en"`
	PageDE        string   `json:"page_de"`
	PageEN        string   `json:"page_en"`
	Src           []string `json:"src"`
	SrdName       string   `json:"srdname"`
	Size          string   `json:"size"`
	Type          string   `json:"type"`
	Tags          string   `json:"tags"`
	Alignment     string   `json:"alignment"`
	Cr            string   `json:"cr"`
	Xp            string   `json:"xp"`
	SingleLine    string   `json:"singleline"`
}

var (
	// awgo specific variable
	wf *aw.Workflow // Our Workflow object

	// base variables
	query         string
	finalURL      strings.Builder
	baseURL       = "https://www.dnddeutsch.de/tools/json.php?apiv=0.7&o=monster"
	helpURL       = "https://github.com/" + repo
	maxResults    = 20
	doTranslateDE bool
	includeSrc    = []string{
		"AI",
		"BGDiA",
		"CM",
		"CoS",
		"CotN",
		"DoIP",
		"EBERRON",
		"EGtW",
		"FToD",
		"GGtR",
		"GoS",
		"HotDQ",
		"IDRotF",
		"LMoP",
		"MC1",
		"MM",
		"MMM",
		"MOoT",
		"MToF",
		"OotA",
		"PotA",
		"RoT",
		"SCC",
		"SKT",
		"SRD",
		"TalDorei",
		"TDR",
		"ToA",
		"TYP",
		"VGM",
		"VRGtR",
		"WbtW",
		"WDH",
		"WDMM",
	}
	excludeSrc = []string{
		"AiME-BRF",
		"AiME-Eria",
		"AiME-RIV",
		"AiME-RRF",
		"AiME-SLH",
		"AiME-WdD",
		"AiME-Wild",
		"AVENT-M",
		"AVENT-W",
		"CC",
		"CTH-GHOUL",
		"CTHULHU",
		"D3",
		"MARGREVE",
		"MTGAFR",
		"Myth-AdDM",
		"Myth-Held",
		"Myth-Saga",
		"RAGNAROK",
		"STRANGE",
		"ToB",
		"ToB2",
	}

	// commandline flags
	doCheck     bool
	doTranslate bool
	// updateCheck target
	repo = "Wayneoween/alfred-dndbeyond-monster-workflow" // GitHub repo

	// cache variables
	cacheName   = "monster_cache.json"       // Filename of cached repo list
	maxCacheAge = 14 * 24 * 60 * time.Minute // Cache each query for 14 days
)

func init() {
	// start building the API URL to access
	finalURL.WriteString(baseURL)

	for _, src := range excludeSrc {
		// add all the sources
		finalURL.WriteString("&xsrc[]=")
		finalURL.WriteString(src)
	}
	finalURL.WriteString("&q=")

	// Create a new *Workflow using default configuration
	// (workflow settings are read from the environment variables
	// set by Alfred)
	wf = aw.New(
		aw.HelpURL(helpURL),
		aw.MaxResults(maxResults),
		update.GitHub(repo),
	)

	// Add a commandline flag to the binary so that the updateCheck can call it
	flag.BoolVar(&doCheck, "check", false, "check for a new version")
	// Add a commandline flag to set the translation
	flag.BoolVar(&doTranslate, "translate", false, "toggle german translation")
}

func run() {
	log.Println("DEBUG: Function 'run'!")

	wf.Args() // call to handle magic actions
	flag.Parse()

	// handle the translation setting
	doTranslateDE := wf.Config.GetBool("translate", false)
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

	if args := wf.Args(); len(args) > 0 {
		query = args[0]
	}

	log.Printf("[main] query=%s", query)
	monsters := []*Monster{}

	// Try to load cached monsters from $query_$cachename.json
	if wf.Cache.Exists(query + "_" + cacheName) {
		log.Println("Data is being loaded from cache.")
		if err := wf.Cache.LoadJSON(query+"_"+cacheName, &monsters); err != nil {
			wf.FatalError(err)
		}
		log.Println("monsters after cache load: ", monsters)
	}

	if wf.Cache.Expired(strings.Replace(query, " ", "-", -1)+"_"+cacheName, maxCacheAge) {
		log.Println("Data is being loaded from website.")

		var resultSet D3ResultSet

		log.Println("Loading data from " + finalURL.String() + query)
		response, err := http.Get(finalURL.String() + query)
		if err != nil {
			fmt.Print(err.Error())
		}

		defer response.Body.Close()

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("responseData: %s", responseData)

		// unmarshall the JSON response into resultSet
		json.Unmarshal(responseData, &resultSet)

		log.Printf("DEBUG: unmarshalled data:")
		log.Println(resultSet)

		// if there are results in the monsterarray for our query
		if len(resultSet.Monster) > 0 {
			log.Printf("DEBUG: %d Monsters were found!", len(resultSet.Monster))

			// range over the array and create entries for every one of them
			log.Println("DEBUG: Printing each monster:")
			for _, result := range resultSet.Monster {

				if len(result.Size) > 0 && len(result.Type) > 0 {
					// add the result fields to temp Monster
					temp := result

					log.Println("MonsterCR:     ", temp.Cr)
					log.Println("MonsterName:   ", temp.NameDE)
					log.Println("MonsterType:   ", temp.Type)
					log.Println("MonsterSize:   ", temp.Size)
					log.Println("-------------------------------------------------")

					monsters = append(monsters, &temp)
				} else {
					log.Println("Not a real monster entry. Skipping.")
					log.Println(result)
				}
			}
		}

		// print the monster array
		log.Println(monsters)

		// write cache only if we have at least one monster
		if len(monsters) != 0 {
			log.Println("More than 1 monsters found. Caching...")
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
					titleDE = fmt.Sprintf("%s", temp.NameEN)
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
				// I need to remove the ( and ) from things like "Vampire (Warrior)"
				// But Ogrillon or Kyton as well as "in lair" is still a problem.
				// TODO: a function that removes special cases here.
				tmpTitle := fmt.Sprintf("%s", temp.NameEN)
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

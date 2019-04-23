// Copyright (c) 2019 Marius Schuller <code@marius-schuller.de>
// MIT Licence - http://opensource.org/licenses/MIT

/*
This Alfred Workflow queries the dndbeyond general search with the monster filter
and displaoys the first 10 matches.
Pressing Enter on each entry uses the default browser to open the monster page on dndbeyond.com
*/
package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// Name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var (
	// base variables
	wf         *aw.Workflow // Our Workflow object
	baseurl    = "https://www.dndbeyond.com"
	helpURL    = "https://marius-schuller.de"
	maxResults = 20
	url        = baseurl + "/monsters?filter-search=" // ddb search url

	// updatecheck variable
	doCheck bool
	repo    = "Wayneoween/alfred-dndbeyond-monster-workflow" // GitHub repo

	// cache variables
	cacheName   = "monsters.json"           // Filename of cached repo list
	maxCacheAge = 7 * 24 * 60 * time.Minute // Cache each query for 7 days

	// Icons
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}
	// monster Type Icons
	monsterIconDefault       = &aw.Icon{Value: "icons/dnd/default.png"}
	monsterIconAbberations   = &aw.Icon{Value: "icons/dnd/aberration.jpg"}
	monsterIconBeasts        = &aw.Icon{Value: "icons/dnd/beast.jpg"}
	monsterIconCelestials    = &aw.Icon{Value: "icons/dnd/celestial.jpg"}
	monsterIconConstructs    = &aw.Icon{Value: "icons/dnd/construct.jpg"}
	monsterIconDragons       = &aw.Icon{Value: "icons/dnd/dragon.jpg"}
	monsterIconElementals    = &aw.Icon{Value: "icons/dnd/elemental.jpg"}
	monsterIconFey           = &aw.Icon{Value: "icons/dnd/fey.jpg"}
	monsterIconFiends        = &aw.Icon{Value: "icons/dnd/fiend.jpg"}
	monsterIconGiants        = &aw.Icon{Value: "icons/dnd/giant.jpg"}
	monsterIconHumanoids     = &aw.Icon{Value: "icons/dnd/humanoid.jpg"}
	monsterIconMonstrosities = &aw.Icon{Value: "icons/dnd/monstrosity.jpg"}
	monsterIconOozes         = &aw.Icon{Value: "icons/dnd/ooze.jpg"}
	monsterIconPlants        = &aw.Icon{Value: "icons/dnd/plant.jpg"}
	monsterIconUndead        = &aw.Icon{Value: "icons/dnd/undead.jpg"}
)

type monster struct {
	MonsterCR   string
	MonsterIcon *aw.Icon
	MonsterName string
	MonsterSize string
	MonsterType string
	MonsterURL  string
}

func init() {
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
}

func run() {
	var query string
	monsters := []*monster{}

	wf.Args() // call to handle magic actions
	flag.Parse()

	// Use wf.Args() to enable Magic Actions
	if args := wf.Args(); len(args) > 0 {
		query = args[0]
	}

	// Try to load cached monsters
	if wf.Cache.Exists(query + "_" + cacheName) {
		log.Println("Data is being loaded from cache.")
		log.Println("monsters before cache load: ", monsters)
		if err := wf.Cache.LoadJSON(query+"_"+cacheName, &monsters); err != nil {
			wf.FatalError(err)
		}
		log.Println("monsters after cache load: ", monsters)
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
			Icon(updateAvailable)
	}

	log.Printf("[main] query=%s", query)

	if wf.Cache.Expired(query+"_"+cacheName, maxCacheAge) {
		log.Println("Data is being loaded from website.")

		// Instantiate default colly collector
		c := colly.NewCollector(
			// Visit only domains: old.reddit.com
			colly.AllowURLRevisit(),
			colly.Async(true),
		)

		// randomize the user agent colly uses
		extensions.RandomUserAgent(c)

		// on every node with class="info"
		c.OnHTML(".info", func(e *colly.HTMLElement) {
			temp := new(monster)
			temp.MonsterCR = e.ChildText(".monster-challenge")
			temp.MonsterType = e.ChildText(".type")
			// for now we use a generic icon for the monster type
			switch strings.ToLower(temp.MonsterType) {
			case "aberration":
				temp.MonsterIcon = monsterIconAbberations
			case "beast":
				temp.MonsterIcon = monsterIconBeasts
			case "celestial":
				temp.MonsterIcon = monsterIconCelestials
			case "construct":
				temp.MonsterIcon = monsterIconConstructs
			case "dragon":
				temp.MonsterIcon = monsterIconDragons
			case "elemental":
				temp.MonsterIcon = monsterIconElementals
			case "fey":
				temp.MonsterIcon = monsterIconFey
			case "fiend":
				temp.MonsterIcon = monsterIconFiends
			case "giant":
				temp.MonsterIcon = monsterIconGiants
			case "humanoid":
				temp.MonsterIcon = monsterIconHumanoids
			case "monstrosity":
				temp.MonsterIcon = monsterIconMonstrosities
			case "ooze":
				temp.MonsterIcon = monsterIconOozes
			case "plant":
				temp.MonsterIcon = monsterIconPlants
			case "undead":
				temp.MonsterIcon = monsterIconUndead
			default:
				temp.MonsterIcon = monsterIconDefault
			}
			temp.MonsterName = e.ChildText(".name")
			temp.MonsterSize = e.ChildText(".monster-size")
			temp.MonsterURL = e.ChildAttr(".name .link", "href")
			monsters = append(monsters, temp)

			log.Println("MonsterCR:   ", temp.MonsterCR)
			log.Println("MonsterIcon: ", temp.MonsterIcon)
			log.Println("MonsterName: ", temp.MonsterName)
			log.Println("MonsterType: ", temp.MonsterType)
			log.Println("MonsterSize: ", temp.MonsterSize)
			log.Println("MonsterURL:  ", baseurl+temp.MonsterURL)
			log.Println("-------------------------------------------------")
		})

		log.Println("Visiting ", url+query)
		// load the website
		c.Visit(url + query)
		// wait until the callbacks finished working
		c.Wait()

		// print the monster array
		log.Println(monsters)

		// write cache only if we have at least one monster
		if len(monsters) != 0 {
			wf.Configure(aw.TextErrors(true))
			if err := wf.Cache.StoreJSON(query+"_"+cacheName, monsters); err != nil {
				wf.FatalError(err)
			}
		}
	}

	// if there are no monsters just send the warning.
	if len(monsters) == 0 {
		wf.WarnEmpty("Nothing found.", "Try another name.")
	} else {
		// no matter if via internet or from the cache, add all monsters as items for alfred
		for _, temp := range monsters {
			wf.NewItem(temp.MonsterName).
				Subtitle("CR " + temp.MonsterCR + " - " + temp.MonsterSize + " - " + temp.MonsterType).
				Icon(temp.MonsterIcon).
				Arg(baseurl + temp.MonsterURL).
				UID(temp.MonsterName + temp.MonsterURL).
				Valid(true)
		}
	}

	// And send the results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}

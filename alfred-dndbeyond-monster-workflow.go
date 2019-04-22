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

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// Name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var (
	doCheck       bool
	baseurl       = "https://www.dndbeyond.com"
	helpURL       = "https://marius-schuller.de"
	maxResults    = 10
	url           = baseurl + "/search?f=monsters&c=monsters&q=" // ddb search url
	iconAvailable = &aw.Icon{Value: "update-available.png"}
	repo          = "Wayneoween/alfred-dndbeyond-monster-workflow" // GitHub repo
	wf            *aw.Workflow                                     // Our Workflow object
)

type monster struct {
	MonsterName string
	MonsterURL  string
}

func init() {
	flag.BoolVar(&doCheck, "check", false, "check for a new version")
	// Create a new *Workflow using default configuration
	// (workflow settings are read from the environment variables
	// set by Alfred)
	wf = aw.New(aw.HelpURL(helpURL), aw.MaxResults(maxResults), update.GitHub(repo))
}

func run() {
	var query string
	monsters := []monster{}
	wf.Args() // call to handle magic actions
	flag.Parse()

	// Use wf.Args() to enable Magic Actions
	if args := wf.Args(); len(args) > 0 {
		query = args[0]
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
			Icon(iconAvailable)
	}

	log.Printf("[main] query=%s", query)

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: old.reddit.com
		colly.AllowURLRevisit(),
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)

	c.OnHTML(".ddb-search-results-listing-item-header-primary-text", func(e *colly.HTMLElement) {
		temp := monster{}
		temp.MonsterName = e.ChildText("a")
		temp.MonsterURL = e.ChildAttr("a", "href")
		monsters = append(monsters, temp)
		wf.NewItem(temp.MonsterName).Subtitle(baseurl + temp.MonsterURL).Arg(baseurl + temp.MonsterURL).UID(temp.MonsterName + temp.MonsterURL).Valid(true)
		log.Println("MonsterName ", temp.MonsterName)
		log.Println("MonsterURL ", baseurl+temp.MonsterURL)
	})

	c.Visit(url + query)

	c.Wait()

	log.Println(monsters)

	if len(monsters) == 0 {
		wf.WarnEmpty("Nothing found.", "Try another name.")
	}

	// And send the results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}

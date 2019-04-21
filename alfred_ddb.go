// Copyright (c) 2019 Marius Schuller <code@marius-schuller.de>
// MIT Licence - http://opensource.org/licenses/MIT

/*
This Alfred Workflow queries the dndbeyond general search with the monster filter
and displaoys the first 10 matches.
Pressing Enter on each entry uses the default browser to open the monster page on dndbeyond.com
*/
package main

import (
	"log"

	aw "github.com/deanishe/awgo"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var (
	baseurl    = "https://www.dndbeyond.com"
	helpURL    = "https://marius-schuller.de"
	maxResults = 10
	url        = baseurl + "/search?f=monsters&c=monsters&q=" // ddb search url
	wf         *aw.Workflow                                   // Our Workflow object
)

type monster struct {
	MonsterName string
	MonsterURL  string
}

func init() {
	// Create a new *Workflow using default configuration
	// (workflow settings are read from the environment variables
	// set by Alfred)
	wf = aw.New(aw.HelpURL(helpURL), aw.MaxResults(maxResults))
}

func run() {
	var query string

	// Use wf.Args() to enable Magic Actions
	if args := wf.Args(); len(args) > 0 {
		query = args[0]
	}

	log.Printf("[main] query=%s", query)

	monsters := []monster{}

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

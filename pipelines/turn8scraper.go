package pipelines

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func Scrape() {
	// Create a new collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.turn8racing.com"),
	)

	// On every <div> element with id="wix-events-widget" call the callback
	c.OnHTML("#wix-events-widget", func(e *colly.HTMLElement) {
		// Split up container by events (RSVP button)
		events := strings.Split(e.Text, "RSVP")
		var lower string
		for i, part := range events {
			if part == "" {
				continue
			}
			lower = strings.ToLower(part)
			fmt.Printf("Part %d: \t%s\n", i+1, strings.TrimSpace(part))
			if strings.Contains(lower, "flash") {
				fmt.Println(Cyan + "FLASH event found!" + Reset)
			} else if strings.Contains(lower, "hpde") {
				fmt.Println(Green + "HPDE event found!" + Reset)
			} else {
				fmt.Println(Red + "Not a flash event" + Reset)
			}
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Print(Blue + "Visiting " + r.URL.String() + Reset + "\n")
	})

	// Start scraping on the URL
	err := c.Visit("https://www.turn8racing.com/track-events")
	if err != nil {
		log.Fatal(err)
	}
}

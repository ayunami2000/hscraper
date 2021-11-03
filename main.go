package main

import (
	"fmt"
	"time"

	"github.com/cutest-design/hscraper/trollcord"
)

func main() {
	fmt.Printf("Failed to load %d tokens! (%d total)\n", trollcord.LoadTokens(Tokens), len(Tokens))
	for {
		tags, err := trollcord.Tags(Guild, Category)

		if err != nil {
			panic(err)
		}

		for ch, tagsraw := range tags {
			tags := normalize(tagsraw)
			fmt.Println(tagsraw, "\tGoroutine start")
			go scrape(ch, tags, tagsraw)
			time.Sleep(1 * time.Second)
		}

		time.Sleep(SleepTime)
	}
}

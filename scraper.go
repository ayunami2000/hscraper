package main

import (
	"fmt"
	"sync"

	"github.com/cutest-design/hscraper/mod"
	_ "github.com/cutest-design/hscraper/mod/mods"
	"github.com/cutest-design/hscraper/seen"
	"github.com/cutest-design/hscraper/trollcord"
)

var cnt map[string]int = make(map[string]int)
var cntmtx = sync.RWMutex{}

func addgetcnt(ch string) int {
	cntmtx.RLock()
	val, exists := cnt[ch]
	cntmtx.RUnlock()

	if !exists {
		cntmtx.Lock()
		cnt[ch] = 0
		cntmtx.Unlock()
		return 0
	} else {
		return val
	}
}

func scrape(ch, tags, tagsraw string) {
	posts := mod.ScrapeAll(addgetcnt(ch), tags)
	fmt.Println(tagsraw, "\tScrape done")
	for _, p := range posts {

		if seen.Seen(ch + ":" + p.MD5) {
			fmt.Println(tagsraw, "\tPost", p.ID, "ignore")
			continue // already seen this post, go to the next one
		}

		seen.Add(ch + ":" + p.MD5)

		trollcord.Send(ch,
			fmt.Sprintf(Template, p.ID, p.Score, p.Tags, p.MD5, p.Module, p.ImageURL),
		)
		fmt.Println(tagsraw, "\tPost", p.ID, "end")
	}
	fmt.Println(tagsraw, "\tGoroutine end")
}

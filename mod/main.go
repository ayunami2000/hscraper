package mod

import (
	"sync"
)

type Tags struct {
	General   []string `json:"general"`
	Species   []string `json:"species"`
	Character []string `json:"character"`
	Artist    []string `json:"artist"`
	Meta      []string `json:"meta"`
}

type Post struct {
	ID       uint32
	Score    int32
	Tags     Tags
	ImageURL string
	MD5      string
	Module   string
}

type Mod func(c int, t string) []Post

var Postch = make(chan Post, 5000)
var Mods []Mod

func ScrapeCh(c int, t string) {
	var wg sync.WaitGroup

	for _, m := range Mods {
		wg.Add(1)
		go func(m Mod) {
			defer wg.Done()
			for _, p := range m(c, t) {
				Postch <- p
			}
		}(m)
	}

	wg.Wait()
}

func ScrapeAll(c int, t string) []Post {
	var res []Post

	for _, m := range Mods {
		mr := m(c, t)
		if mr == nil {
			continue
		}

		res = append(res, mr...)
	}

	return res
}

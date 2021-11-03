package mods

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cutest-design/hscraper/mod"
)

type cureninjapage struct {
	Results []cureninjadata `json:"results"`
}

type cureninjadata struct {
	ID   uint32 `json:"id,string"`
	MD5  string `json:"md5"`
	URL  string `json:"url"`
	Tags string `json:"tags"`
	Site string `json:"site"`
}

func init() {
	mod.Mods = append(mod.Mods, cureninja)
}

func cureninja(c int, t string) []mod.Post {
	req, err := http.NewRequest("GET", "https://cure.ninja/booru/api/json/"+strconv.Itoa(c)+"&q=tag:"+t, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("User-Agent", "BadScraper/1.0.0 (lunawasflaggedagain)")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var resstruct cureninjapage
	err = json.NewDecoder(res.Body).Decode(&resstruct)

	if err != nil {
		fmt.Println(res.Status)
		fmt.Println(err)
		return nil
	}

	var posts []mod.Post

	for _, e := range resstruct.Results {
		posts = append(posts, mod.Post{
			ID:    e.ID,
			Score: 0, // FUCK YOU
			Tags: mod.Tags{
				General: strings.Split(e.Tags, " "),
			},
			ImageURL: e.URL,
			MD5:      e.MD5,
			Module:   e.Site,
		})
	}

	return posts
}

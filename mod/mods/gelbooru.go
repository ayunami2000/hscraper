package mods

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cutest-design/hscraper/mod"
)

type gelbooruDara struct {
	ID      uint32 `json:"id"`
	FileURL string `json:"file_url"`
	Hash    string `json:"hash"`
	Score   int    `json:"score"`
	Tags    string `json:"tags"`
}

func init() {
	mod.Mods = append(mod.Mods, gelbooru)
}

func gelbooru(c int, t string) []mod.Post {
	req, err := http.NewRequest("GET", "https://gelbooru.com/index.php?page=dapi&s=post&q=index&json=1&pid="+strconv.Itoa(c)+"&limit=100&tags="+t, nil)

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

	var resstruct []gelbooruDara
	err = json.NewDecoder(res.Body).Decode(&resstruct)

	if err != nil {
		fmt.Println(res.Status)
		fmt.Println(err)
		return nil
	}

	var posts []mod.Post

	for _, e := range resstruct {
		posts = append(posts, mod.Post{
			ID:    e.ID,
			Score: int32(e.Score),
			Tags: mod.Tags{
				General: strings.Split(e.Tags, " "),
			},
			ImageURL: e.FileURL,
			MD5:      e.Hash,
			Module:   "gelbooru.com",
		})
	}

	return posts
}

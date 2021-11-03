package mods

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cutest-design/hscraper/mod"
)

type e621page struct {
	Posts []e621data `json:"posts"`
}

type e621data struct {
	ID   uint32 `json:"id"`
	File struct {
		MD5 string `json:"md5"`
		URL string `json:"url"`
	} `json:"file"`
	Tags  mod.Tags `json:"tags"`
	Score struct {
		Total int `json:"total"`
	} `json:"score"`
}

func init() {
	mod.Mods = append(mod.Mods, e621)
}

func e621(c int, t string) []mod.Post {
	req, err := http.NewRequest("GET", "https://e621.net/posts.json?page="+strconv.Itoa(c)+"&limit=100&tags="+t, nil)

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

	var resstruct e621page
	err = json.NewDecoder(res.Body).Decode(&resstruct)

	if err != nil {
		fmt.Println(res.Status)
		fmt.Println(err)
		return nil
	}

	var posts []mod.Post

	for _, e := range resstruct.Posts {
		posts = append(posts, mod.Post{
			ID:       e.ID,
			Score:    int32(e.Score.Total),
			Tags:     e.Tags,
			ImageURL: e.File.URL,
			MD5:      e.File.MD5,
			Module:   "e621.net",
		})
	}

	return posts
}

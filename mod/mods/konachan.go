package mods

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cutest-design/hscraper/mod"
)

type konachanpage struct {
	Posts []konachandata `json:"posts"`
}

type konachandata struct {
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
	mod.Mods = append(mod.Mods, konachan)
}

func konachan(c int, t string) []mod.Post {
	req, err := http.NewRequest("GET", "https://konachan.com/post.json?page="+strconv.Itoa(c)+"&limit=100&tags="+t, nil)

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

	var resstruct konachanpage
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
			Module:   "konachan.com",
		})
	}

	return posts
}

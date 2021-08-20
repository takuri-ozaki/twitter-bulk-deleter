package tweet

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

type Entity struct {
	Tweet Tweet `json:"tweet"`
}

type Tweet struct {
	Id        string `json:"id_str"`
	CreatedAt string `json:"created_at"`
}

func Parse(start, end time.Time) ([]int64, int, error) {
	raw, err := ioutil.ReadFile("./tweet.js")
	if err != nil {
		return []int64{}, 0, err
	}

	var tweets []Entity
	prefix := "window.YTD.tweet.part0 = "
	err = json.Unmarshal(raw[len(prefix):], &tweets)
	if err != nil {
		return []int64{}, 0, err
	}

	var targets []int64
	errorCount := 0
	for _, v := range tweets {
		created, err := time.Parse(time.RubyDate, v.Tweet.CreatedAt)
		if err != nil {
			continue
		}
		if created.Before(end) && created.After(start) {
			id, err := strconv.ParseInt(v.Tweet.Id, 10, 64)
			if err != nil {
				errorCount++
				continue
			}
			targets = append(targets, id)
		}
	}

	return targets, errorCount, nil
}

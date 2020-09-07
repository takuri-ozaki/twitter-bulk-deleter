package tweet

import (
	"fmt"
	"twitter-bulk-deleter/internal/util"
)

type Deleter struct {
	api     util.Api
	deleted int
	failed  int
	execute bool
	message string
}

func NewDeleter(api util.Api, execute bool) *Deleter {
	message := " tweets deleted (dryrun)"
	if execute {
		message = " tweets deleted"
	}
	return &Deleter{api: api, execute: execute, message: message}
}

func (d *Deleter) Delete(targets []int64) {
	for _, v := range targets {
		if d.execute {
			_, err := d.api.DeleteTweet(v, true)
			if err != nil {
				d.failed++
				fmt.Println(fmt.Sprint(v) + " failed to delete")
				continue
			}
		}
		fmt.Println(fmt.Sprint(v) + " deleted")
		d.deleted++
	}
	fmt.Println(fmt.Sprint(d.deleted) + d.message)
	if d.failed != 0 {
		fmt.Println(fmt.Sprint(d.failed) + " tweets failed to delete")
	}
}

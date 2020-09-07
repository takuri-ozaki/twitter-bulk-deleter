package util

import (
	"github.com/ChimeraCoder/anaconda"
)

type Api interface {
	DeleteTweet(id int64, trimUser bool) (tweet anaconda.Tweet, err error)
}

func NewRealApi(accessToken string, accessTokenSecret string, consumerKey string, consumerSecret string) Api {
	return anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)
}

type DummyApi struct {
}

func NewDummyAPI() Api {
	return &DummyApi{}
}

func (d DummyApi) DeleteTweet(id int64, _ bool) (tweet anaconda.Tweet, err error) {
	return anaconda.Tweet{Id: id}, nil
}

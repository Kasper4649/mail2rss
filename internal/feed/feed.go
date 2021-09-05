package feed

import (
	"github.com/gorilla/feeds"
	"github.com/tidwall/gjson"
	"time"
)

func parseFeedItem(data string) ([]*feeds.Item, error) {
	var items []*feeds.Item
	emails := gjson.Get(data, "emails").Array()
	for _, e := range emails {
		items = append(items, &feeds.Item{
			Title:       e.Get("subject").String(),
			Description: e.Get("html").String(),
			Created:     time.UnixMilli(e.Get("timestamp").Int()),
			Author: &feeds.Author{
				Name: e.Get("from").String(),
			},
			Id: e.Get("id").String(),
			Link: &feeds.Link{
				Href: e.Get("downloadUrl").String(),
			},
		})
	}
	return items, nil
}

func MakeRSS(data, tag string) (string, error) {
	items, err := parseFeedItem(data)
	if err != nil {
		return "", err
	}
	feed := &feeds.Feed{
		Title: tag + " Mail RSS",
		Link: &feeds.Link{
			Href: "",
		},
		Description: tag + " Mail RSS",
		Copyright:   "https://github.com/Kasper4649/mail2rss",
		Created:     time.Now(),
	}
	for _, v := range items {
		feed.Add(v)
	}

	rss, err := feed.ToAtom()
	if err != nil {
		return "", err
	}
	return rss, nil
}

package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type hn_item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Id          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Url         string `json:"url"`
}

const (
	hn_api = "https://hacker-news.firebaseio.com/v0"
)

type Client struct {
	apibase string
}

func (c *Client) defaulty() {
	if c.apibase == "" {
		c.apibase = hn_api
	}
}

func (c *Client) FetchData() ([]int, error) {
	c.defaulty()

	res, err := http.Get(fmt.Sprintf("%s/topstories.json", c.apibase))

	if err != nil {
		return nil, err
	}
	var ids []int

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&ids)

	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	return ids, nil

}

func (c *Client) GetItem(id int) (hn_item, error) {

	var item hn_item

	res, err := http.Get(fmt.Sprintf("%s/item/%d.json", c.apibase, id))

	if err != nil {
		return item, err
	}

	dec := json.NewDecoder(res.Body)

	err = dec.Decode(&item)

	if err != nil {
		return item, err
	}

	defer res.Body.Close()

	return item, nil

}

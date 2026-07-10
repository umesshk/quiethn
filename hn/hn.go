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

func (c *Client) FetchData() {
	c.defaulty()

	res, err := http.Get(fmt.Sprintf("%s/topstories.json", c.apibase))

	if err != nil {
		fmt.Println(err)
		panic("Error Occured fetching ")
	}

	var ids []int

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&ids)

	if err != nil {
		fmt.Println("Error Occured", err)
	}

	fmt.Println(ids)

	defer res.Body.Close()
}

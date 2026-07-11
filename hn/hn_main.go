package hn

import (
	"flag"
	"fmt"
)

func StartServer() {

	var port, numStories int

	flag.IntVar(&port, "port", 3000, "provide the port number for the server ")
	flag.IntVar(&numStories, "num", 30, "provide the number of stories to fetch ")

	flag.Parse()

	var c Client

	ids, err := c.FetchData()

	if err != nil {
		return
	}

}

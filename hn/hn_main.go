package hn

import "fmt"

func StartServer() {

	var c Client

	ids, err := c.FetchData()

	if err != nil {
		return
	}

	item, err := c.GetItem(ids[0])

	fmt.Println(item)

}

package hn

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"
)

func StartServer() {

	var port, numStories int

	flag.IntVar(&port, "port", 3000, "provide the port number for the server ")
	flag.IntVar(&numStories, "num", 30, "provide the number of stories to fetch ")

	flag.Parse()

	tpl := template.Must(template.ParseFiles("hn/index.html"))

	http.HandleFunc("/", handler(numStories, tpl))

	portstr := fmt.Sprintf(":%d", port)

	log.Println("Server started at loalhost:", port)

	log.Fatal(http.ListenAndServe(portstr, nil))

}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		var c Client
		fmt.Println("Fetching the story id's.. ")
		ids, err := c.FetchData()

		if err != nil {
			http.Error(w, "Error fetching data", http.StatusInternalServerError)
			return
		}

		var stories []item

		fmt.Println("Getting the top stories..")
		for _, id := range ids {

			fmt.Println("Getting Story id : ", id)
			story_item, err := c.GetItem(id)

			if err != nil {
				http.Error(w, "Error fetching data", http.StatusInternalServerError)
				return

			}

			if isLinkStory(story_item) {
				item := parseHn(story_item)
				stories = append(stories, item)

				if len(stories) >= numStories {
					break
				}
			}

		}

		data := templateData{
			Stories: stories,
			Time:    time.Since(start),
		}

		fmt.Println("Writing to HTML template..")
		err = tpl.Execute(w, data)

		if err != nil {
			log.Fatal(err)
			http.Error(w, "Error Writing to Html ", http.StatusInternalServerError)
			return
		}

		fmt.Println("Page Rendered...", time.Since(start))

	}
}

func isLinkStory(item hn_item) bool {
	return item.Type == "story" && item.Url != ""
}

func parseHn(data hn_item) item {

	new_item := item{hn_item: data}

	url, err := url.Parse(new_item.Url)

	if err == nil {
		new_item.Host = strings.TrimPrefix(url.Host, "www.")
	}

	return new_item

}

type item struct {
	hn_item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}

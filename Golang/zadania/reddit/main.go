package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"reddit/fetcher"
)

func main() {
	var f fetcher.RedditFetcher // do not change
	var w io.Writer             // do not change

	f = &fetcher.HttpFetcher{URL: "https://www.reddit.com/r/golang.json"}

	err := f.Fetch()

	if err != nil {
		log.Println(fmt.Errorf("cannot fetch: %w", err))
		os.Exit(1)
	}

	w, err = os.Create("output.txt")

	err = f.Save(w)
	if err != nil {
		log.Println(fmt.Errorf("cannot save: %w", err))
		os.Exit(1)
	}
}

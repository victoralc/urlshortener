package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"urlshortener/handlers"
	"urlshortener/url"
)

var (
	port      *int
	enableLog *bool
	urlBase   string
)

func init() {
	port = flag.Int("p", 8888, "port")
	enableLog = flag.Bool("l", true, "log enabled/disabled")
	urlBase = fmt.Sprintf("http://localhost:%d", *port)
	flag.Parse()
	repo := url.NewInMemoryRepository()
	url.SetRepository(repo)
}

func main() {
	stats := make(chan string)
	defer close(stats)
	go url.RegisterStatistics(stats)

	http.Handle("/api/short", &handlers.Shortener{UrlBase: urlBase})
	http.Handle("/r/", &handlers.Redirect{Stats: stats})
	http.HandleFunc("/api/stats/", handlers.UrlStatsHandler)

	log.Printf("Initializing server on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

package handlers

import (
	"net/http"
	"urlshortener/url"
)

type Redirect struct {
	Stats chan string
}

func (r *Redirect) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	url.FindUrlAndExecute(w, req, func(url *url.URL) {
		http.Redirect(w, req, url.Destiny, http.StatusMovedPermanently)
		r.Stats <- url.ID
	})
}

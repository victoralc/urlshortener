package handlers

import (
	json2 "encoding/json"
	"net/http"
	"urlshortener/http/response"
	"urlshortener/url"
)

func UrlStatsHandler(w http.ResponseWriter, r *http.Request) {
	url.FindUrlAndExecute(w, r, func(url *url.URL) {
		json, err := json2.Marshal(url.Stats())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response.WithJson(w, string(json))
	})
}

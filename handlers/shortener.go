package handlers

import (
	"fmt"
	"log"
	"net/http"
	"urlshortener/http/response"
	"urlshortener/url"
)

type Shortener struct {
	UrlBase string
}

func (s *Shortener) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		response.With(w, http.StatusMethodNotAllowed, response.Headers{
			"Allow": "POST",
		})
		return
	}
	url, newUrl, err := url.FindOrCreateNewURL(url.ExtractUrl(r))
	if err != nil {
		response.With(w, http.StatusBadRequest, nil)
		return
	}
	var status int
	if newUrl {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}
	response.With(w, status, response.Headers{
		"Location": fmt.Sprintf(s.UrlBase+"/r/%s", url.ID),
		"Link":     fmt.Sprintf("<%s/api/stats/%s>; rel=\"stats\"", s.UrlBase, url.ID),
	})

	log.Printf("URL %s shortened successfully for %s", url.Destiny, url)
}

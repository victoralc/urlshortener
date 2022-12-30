package url

import (
	"log"
	"math/rand"
	"net/http"
	url2 "net/url"
	"strings"
	"time"
)

const (
	size    = 5
	symbols = "abcdefghijklmnopqr...STUVWXYZ1234567890_-+"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type URL struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Destiny   string    `json:"destiny"`
}

type Stats struct {
	Url    *URL `json:"url"`
	Clicks int  `json:"clicks"`
}

func ExtractUrl(r *http.Request) string {
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)
}

func FindOrCreateNewURL(destiny string) (u *URL, newUrl bool, err error) {
	if u := repo.FindByURL(destiny); u != nil {
		return u, false, nil
	}
	if _, err = url2.ParseRequestURI(destiny); err != nil {
		return nil, false, err
	}
	url := URL{generateId(), time.Now(), destiny}
	repo.Save(url)
	return &url, true, nil
}

func Find(id string) *URL {
	return repo.FindByID(id)
}

func RegisterClick(id string) {
	repo.RegisterClick(id)
}

func (u *URL) Stats() *Stats {
	clicks := repo.FindClicks(u.ID)
	return &Stats{Url: u, Clicks: clicks}
}

func RegisterStatistics(ids <-chan string) {
	for id := range ids {
		RegisterClick(id)
		log.Printf("Click registered successfully for %s. \n", id)
	}
}

func generateId() string {
	newId := func() string {
		id := make([]byte, size, size)
		for i := range id {
			id[i] = symbols[rand.Intn(len(symbols))]
		}
		return string(id)
	}

	for {
		if id := newId(); !repo.ExistID(id) {
			return id
		}
	}
}

func FindUrlAndExecute(w http.ResponseWriter, r *http.Request, executor func(url *URL)) {
	//Split in tokens separated by '/'
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]
	if url := Find(id); url != nil {
		executor(url)
	} else {
		http.NotFound(w, r)
	}
}

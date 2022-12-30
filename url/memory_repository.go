package url

type Repository interface {
	ExistID(id string) bool
	FindByID(id string) *URL
	FindByURL(url string) *URL
	Save(url URL) error
	RegisterClick(id string)
	FindClicks(id string) int
}

type inMemoryRepository struct {
	urls   map[string]*URL
	clicks map[string]int
}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{make(map[string]*URL), make(map[string]int)}
}

func (r *inMemoryRepository) ExistID(id string) bool {
	_, exist := r.urls[id]
	return exist
}

func (r *inMemoryRepository) FindByID(id string) *URL {
	return r.urls[id]
}

func (r *inMemoryRepository) FindByURL(url string) *URL {
	for _, u := range r.urls {
		if u.Destiny == url {
			return u
		}
	}
	return nil
}

func (r *inMemoryRepository) Save(url URL) error {
	r.urls[url.ID] = &url
	return nil
}

func (r *inMemoryRepository) RegisterClick(id string) {
	r.clicks[id] += 1
}

func (r *inMemoryRepository) FindClicks(id string) int {
	return r.clicks[id]
}

var repo Repository

func SetRepository(r Repository) {
	repo = r
}

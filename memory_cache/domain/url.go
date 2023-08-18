package domain

type Url struct {
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
}

func (url *Url) Shorten() {
	url.Shortened = "http://localhost:8080/shortened"
}

package exercises

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (exe *Exercises) unmarshalCharactersResponse(data []byte) (CharactersResponse, error) {
	var r CharactersResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

type CharactersResponse struct {
	Info    Info     `json:"info"`
	Results []Result `json:"results"`
}

type Info struct {
	Count int64       `json:"count"`
	Pages int64       `json:"pages"`
	Next  string      `json:"next"`
	Prev  interface{} `json:"prev"`
}

type Result struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Status   Status   `json:"status"`
	Species  Species  `json:"species"`
	Type     string   `json:"type"`
	Gender   Gender   `json:"gender"`
	Origin   Location `json:"origin"`
	Location Location `json:"location"`
	Image    string   `json:"image"`
	Episode  []string `json:"episode"`
	URL      string   `json:"url"`
	Created  string   `json:"created"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Gender string

const (
	Female        Gender = "Female"
	GenderUnknown Gender = "unknown"
	Male          Gender = "Male"
)

type Species string

const (
	Alien Species = "Alien"
	Human Species = "Human"
)

type Status string

const (
	Alive         Status = "Alive"
	Dead          Status = "Dead"
	StatusUnknown Status = "unknown"
)

func (exe *Exercises) getRickAndMortyCharacters() (*CharactersResponse, error) {
	client := &http.Client{}
	url := "https://rickandmortyapi.com/api/character"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	queryParams := req.URL.Query()
	queryParams.Add("count", "20")
	queryParams.Add("page", "1")

	req.URL.RawQuery = queryParams.Encode()

	fmt.Println(req.URL)
	res, err := client.Do(req)

	if res != nil && res.StatusCode != 200 {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	characters, err := exe.unmarshalCharactersResponse(data)

	if err != nil {
		return nil, err
	}
	return &characters, nil
}

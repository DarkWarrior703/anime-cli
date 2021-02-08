package anime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	nsfw  = true
	limit = -1
)

// Temporal class for getAnimeList
type tmp struct {
	Results []struct {
		Title    string `json:"title"`
		Airing   bool   `json:"airing"`
		Type     string `json:"type"`
		Episodes int    `json:"episodes"`
		Rated    string `json:"rated"`
	} `json:"results"`
}

// Anime struct stands for a class containing details about animes.
type Anime struct {
	Title    string
	Status   string
	Type     string
	Episodes int
	Rated    string
}

func getAnimeList(query string) (*tmp, error) {
	url := "https://api.jikan.moe/v3/search/anime?" + parameters(query)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	tmpdata := &tmp{}
	err = json.Unmarshal(body, tmpdata)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return tmpdata, nil
}

func parameters(query string) string {
	str := fmt.Sprintf("q=%s", query)
	if !nsfw {
		str += "&genre=12&genre_exclude=0"
	}
	if limit > 0 {
		str += fmt.Sprintf("&limit=%d", limit)
	}
	return str
}

//RetrieveAnimeData fetches jikan.moe API for retrieving anime data.
//Takes a query as the argument
func RetrieveAnimeData(query string) ([]Anime, error) {
	data := []Anime{}
	animeList, err := getAnimeList(query)
	if err != nil {
		return nil, err
	}
	for _, anime := range animeList.Results {
		status := ""
		if anime.Airing {
			status = "Outgoing"
		} else {
			status = "Finished"
		}
		data = append(data, Anime{
			Title:    anime.Title,
			Status:   status,
			Type:     anime.Type,
			Episodes: anime.Episodes,
			Rated:    anime.Rated,
		})
	}
	return data, nil
}

// SetNSFW takes an argument, if true RetrieveAnimeData returns NSFW animes, otherwise no.
func SetNSFW(state bool) {
	nsfw = state
}

// SetLimit takes an argument the limit.
func SetLimit(arg int) {
	limit = arg
}

package manga

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

type tmp struct {
	Results []struct {
		Title    string `json:"title"`
		Synopsis string `json:"synopsis"`
		Type     string `json:"type"`
		Chapters int    `json:"chapters"`
		Volumes  int    `json:"volumes"`
	} `json:"results"`
}

// Manga struct stands for a class containing details about mangas.
type Manga struct {
	Title    string
	Synopsis string
	Type     string
	Chapters int
	Volumes  int
}

func getMangaList(query string) (*tmp, error) {
	url := "https://api.jikan.moe/v3/search/manga?" + parameters(query)
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

//RetrieveMangaData fetches jikan.moe API for retrieving manga data.
//Takes a string as the argument
func RetrieveMangaData(query string) ([]Manga, error) {
	data := []Manga{}
	mangaList, err := getMangaList(query)
	if err != nil {
		return nil, err
	}
	for _, manga := range mangaList.Results {
		data = append(data, Manga{
			Title:    manga.Title,
			Synopsis: manga.Synopsis,
			Type:     manga.Type,
			Chapters: manga.Chapters,
			Volumes:  manga.Volumes,
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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Anime struct {
	Results []struct {
		Title    string      `json:"title"`
		Airing   interface{} `json:"airing"`
		Type     string      `json:"type"`
		Episodes int         `json:"episodes"`
		Rated    string      `json:"rated"`
	} `json:"results"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need an argument.")
	} else if os.Args[1] == "help" {
		fmt.Println("You need to add the name of the anime.")
	}
	url := "https://api.jikan.moe/v3/search/anime?q=" + strings.Join(os.Args, "+")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(resp.Body)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	animeList := &Anime{}
	err = json.Unmarshal(body, animeList)
	if err != nil {
		log.Fatal(err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "Episodes", "Type", "Status", "Rating"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiBlueColor},
		tablewriter.Colors{tablewriter.FgHiMagentaColor},
		tablewriter.Colors{tablewriter.FgHiCyanColor})
	table.SetColumnColor(tablewriter.Colors{tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.FgCyanColor})
	for _, anime := range animeList.Results {
		status := ""
		if anime.Airing == false {
			status = "Finished"
		} else if anime.Airing == true {
			status = "Ongoing"
		}
		table.Append([]string{anime.Title, fmt.Sprint(anime.Episodes), anime.Type, status, anime.Rated})
	}
	table.Render()
}

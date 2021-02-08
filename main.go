package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/DarkWarrior703/anime-cli/anime"
	"github.com/DarkWarrior703/anime-cli/manga"
	"github.com/olekukonko/tablewriter"
)

func main() {
	table := tablewriter.NewWriter(os.Stdout)
	if len(os.Args) < 2 {
		log.Fatal("You need an argument.")
	} else if os.Args[1] == "help" {
		fmt.Println("Usage: anime-cli <command> query\n<command> is either 'anime' or 'manga'")
		os.Exit(0)
	} else if os.Args[1] == "anime" {
		animeList, err := anime.RetrieveAnimeData(strings.Join(os.Args[2:], "+"))
		if err != nil {
			log.Fatal(err)
		}
		table.SetHeader([]string{"Title", "Episodes", "Type", "Status", "Rating"})
		for _, anime := range animeList {
			table.Append([]string{anime.Title, fmt.Sprint(anime.Episodes), anime.Type, anime.Status, anime.Rated})
		}
	} else if os.Args[1] == "manga" {
		mangaList, err := manga.RetrieveMangaData(strings.Join(os.Args[2:], "+"))
		if err != nil {
			log.Fatal(err)
		}
		table.SetHeader([]string{"Title", "Synopsis", "Type", "Chapters", "Volumes"})
		for _, manga := range mangaList {
			table.Append([]string{manga.Title, manga.Synopsis, manga.Type, fmt.Sprint(manga.Chapters), fmt.Sprint(manga.Volumes)})
		}
	}
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
	table.Render()
}

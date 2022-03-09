package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	_ "github.com/mattn/go-sqlite3"
)

var (
	artistName string
	albumName  string
	songName   string
	plays      string
)

func topSongs(database *sql.DB) string {
	rows, err := database.Query(`SELECT artists.name, songs.name, COUNT(*) as count FROM songs LEFT JOIN albums ON songs.album = albums.id LEFT JOIN artists ON albums.artist = artists.id WHERE date>=date('2022-01-01') GROUP BY songs.name ORDER BY count DESC`)
	if err != nil {
		log.Fatal(err)
	}

	table := fmt.Sprintf("| Artist | Song | > |\n| --- | --- | --- |\n")
	for i := 1; rows.Next(); i++ {
		rows.Scan(&artistName, &songName, &plays)
		table += fmt.Sprintf("| %s | %s | %s |\n", artistName, songName, plays)
		if i == 10 {
			break
		}
	}
	rows.Close()
	return table
}

func topArtists(database *sql.DB) string {
	rows, err := database.Query(`SELECT artists.name, COUNT(*) as count FROM songs LEFT JOIN albums ON songs.album = albums.id LEFT JOIN artists ON albums.artist = artists.id WHERE date>=date('2022-01-01') GROUP BY artists.name ORDER BY count DESC`)
	if err != nil {
		log.Fatal(err)
	}

	table := fmt.Sprintf("| Artist | > |\n| --- | --- |\n")
	for i := 1; rows.Next(); i++ {
		rows.Scan(&artistName, &plays)
		table += fmt.Sprintf("| %s | %s |\n", artistName, plays)
		if i == 10 {
			break
		}
	}
	rows.Close()
	return table
}

func topAlbums(database *sql.DB) string {
	rows, err := database.Query(`SELECT albums.name, artists.name, COUNT(*) as count FROM songs LEFT JOIN albums ON songs.album = albums.id LEFT JOIN artists ON albums.artist = artists.id WHERE date>=date('2022-01-01') GROUP BY albums.id ORDER BY count DESC
`)
	if err != nil {
		log.Fatal(err)
	}

	table := fmt.Sprintf("| Album | Artist | > |\n| --- | --- | --- |\n")
	for i := 1; rows.Next(); i++ {
		rows.Scan(&albumName, &artistName, &plays)
		table += fmt.Sprintf("| %s | %s | %s |\n", albumName, artistName, plays)
		if i == 10 {
			break
		}
	}
	rows.Close()
	return table
}

func main() {
	var (
		dbPath string
		table  string
		out    string
	)
	if len(os.Args) > 1 {
		dbPath = os.Args[1]
	} else {
		dirname, err := os.UserConfigDir()
		if err != nil {
			log.Fatal(err)
		}
		dbPath = dirname + "/musyca/database.db"
	}

	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// table = topArtists(database)

	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(60),
	)

	out, err = r.Render(table)
	fmt.Print(out)
	database.Close()

	ui()
}

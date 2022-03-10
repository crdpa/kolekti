package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/charmbracelet/glamour"
	_ "github.com/mattn/go-sqlite3"
)

var (
	artistName string
	albumName  string
	songName   string
	plays      string
	sqlString  string
)

// cut long names
func wordWrap(word string) string {
	if len(word) > 26 {
		return word[:26] + "…"
	}
	return word
}

func topSongs(database *sql.DB, dateString string) string {
	sqlQuery := `SELECT artists.name, songs.name, COUNT(*) as count FROM songs LEFT JOIN albums ON songs.album = albums.id LEFT JOIN artists ON albums.artist = artists.id ` + dateString + ` GROUP BY songs.name ORDER BY count DESC`
	rows, err := database.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	table := fmt.Sprintf("| Artist | Song | > |\n| --- | --- | --- |\n")
	for i := 1; rows.Next(); i++ {
		rows.Scan(&artistName, &songName, &plays)
		artistName = wordWrap(artistName)
		songName = wordWrap(songName)
		table += fmt.Sprintf("| %s | %s | %s |\n", artistName, songName, plays)
		if i == 10 {
			break
		}
	}
	rows.Close()
	return table
}

func topArtists(database *sql.DB, dateString string) string {
	sqlQuery := `SELECT artists.name, COUNT(*) as count FROM songs LEFT JOIN albums ON songs.album = albums.id LEFT JOIN artists ON albums.artist = artists.id ` + dateString + ` GROUP BY artists.name ORDER BY count DESC`
	rows, err := database.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	table := fmt.Sprintf("| Artist | > |\n| --- | --- |\n")
	for i := 1; rows.Next(); i++ {
		rows.Scan(&artistName, &plays)
		artistName = wordWrap(artistName)
		table += fmt.Sprintf("| %s | %s |\n", artistName, plays)
		if i == 10 {
			break
		}
	}
	rows.Close()
	return table
}

func topAlbums(database *sql.DB, dateString string) string {
	sqlQuery := `SELECT albums.name, artists.name, COUNT(*) as count FROM songs LEFT JOIN albums ON songs.album = albums.id LEFT JOIN artists ON albums.artist = artists.id ` + dateString + ` GROUP BY albums.id ORDER BY count DESC`
	rows, err := database.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	table := fmt.Sprintf("| Album | Artist | > |\n| --- | --- | --- |\n")
	for i := 1; rows.Next(); i++ {
		rows.Scan(&albumName, &artistName, &plays)
		artistName = wordWrap(artistName)
		albumName = wordWrap(songName)
		table += fmt.Sprintf("| %s | %s | %s |\n", albumName, artistName, plays)
		if i == 10 {
			break
		}
	}
	rows.Close()
	return table
}

// Parse dates. If not a valid date, use '*'.
func parseDate(startDate string, endDate string) (string, string) {
	const layout = "2006-01-02"

	_, err := time.Parse(layout, startDate)
	if err != nil {
		startDate = "*"
	}
	_, err = time.Parse(layout, endDate)
	if err != nil {
		endDate = "*"
	}
	return startDate, endDate
}

// convert dates to sql language
func dateToSql(startDate string, endDate string) string {
	if startDate == "*" {
		if endDate == "*" {
			sqlString = ""
		} else {
			sqlString = `WHERE date<=('` + endDate + `')`
		}
	} else {
		if endDate == "*" {
			sqlString = `WHERE date >=('` + startDate + `')`
		} else {
			sqlString = `WHERE date BETWEEN '` + startDate + `' AND '` + endDate + `'`
		}
	}
	return sqlString
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

	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		// glamour.WithWordWrap(60),
	)

	ui()
	startDate, endDate = parseDate(startDate, endDate)
	dateString := dateToSql(startDate, endDate)
	table = topArtists(database, dateString)
	out, err = r.Render(table)
	fmt.Print(out)
	database.Close()
}

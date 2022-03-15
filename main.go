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

func wordWrap(word string) string {
	if len(word) > 26 {
		return word[:26] + "…"
	}
	return word
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

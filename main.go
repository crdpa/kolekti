package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nathan-fiscaletti/consolesize-go"
)

var (
	artistName       string
	albumName        string
	songName         string
	plays            string
	sqlString        string
	colSize, lastCol int
)

func adjustTable(columns int) int {
	const lastCol = 6
	cols, _ := consolesize.GetConsoleSize()
	colSize := (cols - lastCol) / (columns - 1)
	return colSize
}

func wordWrap(word string, colSize int) string {
	if len(word) > colSize {
		return word[:colSize-1] + "…"
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

	ui()
	startDate, endDate = parseDate(startDate, endDate)
	dateString := dateToSql(startDate, endDate)
	table = topSongs(database, dateString)
	fmt.Print(table)
	database.Close()
}

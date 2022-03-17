package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/muesli/termenv"
)

var (
	artistName       string
	albumName        string
	songName         string
	plays            string
	sqlString        string
	colSize, lastCol int
)

func main() {
	var (
		dbPath string
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
	limitInt := checkLimit(limit)
	table := getData(database, dateString, "songs", limitInt)
	p := termenv.ColorProfile()

	for i := range table {
		s := termenv.String(table[i])
		if table[i] == "" {
			break
		}
		if i == 0 {
			fmt.Println(s.Bold().Underline().Foreground(p.Color("2")))
			continue
		} else if i%2 == 0 {
			fmt.Println(s.Faint())
		} else {
			fmt.Println(s)
		}
	}
	database.Close()
}

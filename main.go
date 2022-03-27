package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/muesli/termenv"
)

var (
	artistName       string
	albumName        string
	songName         string
	plays            string
	sqlString        string
	limit            string
	colSize, lastCol int
)

func main() {
	var (
		dbPath string
	)

	dirname, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	dbPath = dirname + "/musyca/database.db"

	limitFlag := flag.String("l", "10", "Number of results to display.")
	startDateFlag := flag.String("s", "2000-12-30", "Start date")
	endDateFlag := flag.String("e", time.Now().Local().Format("2006-01-02"), "End date")
	dbFlag := flag.String("db", dbPath, "Database path")
	dataFlag := flag.String("data", "songs", "Songs, artists or albums.")
	flag.Parse()

	d := strings.ToLower(*dataFlag)

	if d != "songs" && d != "artists" && d != "albums" {
		d = "songs"
	}

	database, err := sql.Open("sqlite3", *dbFlag)
	if err != nil {
		log.Fatal(err)
	}

	limit = checkLimit(*limitFlag)
	dateString := dateToSql(*startDateFlag, *endDateFlag)
	table := getData(database, dateString, d)
	p := termenv.ColorProfile()

	fmt.Println()
	for i := range table {
		s := termenv.String(table[i])
		if table[i] == "" {
			break
		}
		if i == 0 {
			fmt.Println(s.Bold().Underline().
				Foreground(p.Color("2")))
			continue
		} else if i%2 == 0 {
			fmt.Println(s.Faint())
		} else {
			fmt.Println(s)
		}
	}
	database.Close()
}

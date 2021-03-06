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
	dataFlag := flag.String("data", "songs", "songs, artists or albums.")
	expFlag := flag.String("export", "", "Write data to a file.\n-export=/home/user/data.txt")
	flag.Parse()

	d := strings.ToLower(*dataFlag)

	if d != "songs" && d != "artists" && d != "albums" {
		d = "songs"
	}

	database, err := sql.Open("sqlite3", *dbFlag)
	if err != nil {
		log.Fatal(err)
	}

	limit := checkLimit(*limitFlag)
	dateString := dateToSql(*startDateFlag, *endDateFlag)
	table := getData(database, dateString, d, limit)
	p := termenv.ColorProfile()

	if *expFlag == "" {
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
	} else {
		f, err := os.Create(*expFlag)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		for _, v := range table {
			_, err2 := f.WriteString(v + "\n")
			if err2 != nil {
				log.Fatal(err)
			}
		}

		fmt.Printf("Data exported to %s\n", *expFlag)
	}
	database.Close()
}

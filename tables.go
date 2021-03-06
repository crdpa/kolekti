package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/nathan-fiscaletti/consolesize-go"
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

func checkLimit(limit string) string {
	if _, err := strconv.Atoi(limit); err != nil {
		return "10"
	}

	return limit
}

// convert dates to sql language
func dateToSql(startDate string, endDate string) string {
	const layout = "2006-01-02"
	var sqlString string

	_, err := time.Parse(layout, startDate)
	if err != nil {
		startDate = "*"
	}
	_, err = time.Parse(layout, endDate)
	if err != nil {
		endDate = "*"
	}

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
			sqlString = `WHERE date BETWEEN '` + startDate +
				`' AND '` + endDate + `'`
		}
	}

	return sqlString
}

func getData(database *sql.DB, dateString, what, limit string) []string {
	var (
		artistName string
		albumName  string
		songName   string
		plays      string
		table      []string
	)

	switch what {
	case "songs":
		sqlQuery := `SELECT artists.name, songs.name, COUNT(*)
		as count FROM songs LEFT JOIN albums ON songs.album =
		albums.id LEFT JOIN artists ON albums.artist = artists.id ` +
			dateString + ` GROUP BY songs.name ORDER BY count DESC LIMIT ` + limit

		rows, err := database.Query(sqlQuery)
		if err != nil {
			log.Fatal(err)
		}

		colSize = adjustTable(3)

		table = append(table, fmt.Sprintf("%-*s %-*s %-s", colSize, "Song", colSize, "Artist", ">"))

		for i := 1; rows.Next(); i++ {
			rows.Scan(&artistName, &songName, &plays)
			songName = strconv.Itoa(i) + ". " + songName
			songName = wordWrap(songName, colSize)
			artistName = wordWrap(artistName, colSize)
			table = append(table, fmt.Sprintf("%-*s %-*s %-s", colSize, songName,
				colSize, artistName, plays))
		}
		rows.Close()

	case "artists":
		sqlQuery := `SELECT artists.name, COUNT(*) as count
		FROM songs LEFT JOIN albums ON songs.album = albums.id
		LEFT JOIN artists ON albums.artist = artists.id ` +
			dateString + ` GROUP BY artists.name ORDER BY
			count DESC LIMIT ` + limit

		rows, err := database.Query(sqlQuery)
		if err != nil {
			log.Fatal(err)
		}

		colSize = adjustTable(2)

		table = append(table, fmt.Sprintf("%-*s %-s", colSize, "Artist", ">"))

		for i := 1; rows.Next(); i++ {
			rows.Scan(&artistName, &plays)
			artistName = strconv.Itoa(i) + ". " + artistName
			artistName = wordWrap(artistName, colSize)
			table = append(table, fmt.Sprintf("%-*s %-s", colSize, artistName, plays))
		}
		rows.Close()

	case "albums":
		sqlQuery := `SELECT albums.name, artists.name, COUNT(*) as count
		FROM songs LEFT JOIN albums ON songs.album = albums.id LEFT JOIN
		artists ON albums.artist = artists.id ` + dateString +
			` GROUP BY albums.id ORDER BY count DESC LIMIT ` + limit

		rows, err := database.Query(sqlQuery)
		if err != nil {
			log.Fatal(err)
		}

		colSize := adjustTable(3)

		table = append(table, fmt.Sprintf("%-*s %-*s %-s", colSize, "Album", colSize, "Artist", ">"))
		for i := 1; rows.Next(); i++ {
			rows.Scan(&albumName, &artistName, &plays)
			albumName = strconv.Itoa(i) + ". " + albumName
			albumName = wordWrap(albumName, colSize)
			artistName = wordWrap(artistName, colSize)
			table = append(table, fmt.Sprintf("%-*s %-*s %-s", colSize,
				albumName, colSize, artistName, plays))
		}
		rows.Close()
	}

	return table
}

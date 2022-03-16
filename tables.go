package main

import (
	"database/sql"
	"fmt"
	"log"
)

func topSongs(database *sql.DB, dateString string) string {
	sqlQuery := `SELECT artists.name, songs.name, COUNT(*) as count FROM songs LEFT JOIN albums ON songs.album = albums.id LEFT JOIN artists ON albums.artist = artists.id ` + dateString + ` GROUP BY songs.name ORDER BY count DESC`
	rows, err := database.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	colSize = adjustTable(3)

	table := fmt.Sprintf("%-*s %-*s %-s\n", colSize, "Artist", colSize, "Song", ">")
	for i := 1; rows.Next(); i++ {
		rows.Scan(&artistName, &songName, &plays)
		artistName = wordWrap(artistName, colSize)
		songName = wordWrap(songName, colSize)
		table += fmt.Sprintf("%*s %-*s %-s\n", colSize, artistName, colSize, songName, plays)
		if i == 10 {
			break
		}
	}
	rows.Close()
	fmt.Println(colSize)
	return table
}

func topArtists(database *sql.DB, dateString string) string {
	sqlQuery := `SELECT artists.name, COUNT(*) as count FROM songs LEFT JOIN albums ON songs.album = albums.id LEFT JOIN artists ON albums.artist = artists.id ` + dateString + ` GROUP BY artists.name ORDER BY count DESC`
	rows, err := database.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	colSize := adjustTable(2)

	table := fmt.Sprintf("%-*s %-s\n", colSize, "Artist", ">")
	for i := 1; rows.Next(); i++ {
		rows.Scan(&artistName, &plays)
		artistName = wordWrap(artistName, colSize)
		table += fmt.Sprintf("%-*s %-s\n", colSize, artistName, plays)
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

	colSize := adjustTable(3)

	table := fmt.Sprintf("%-*s %-*s %-s\n", colSize, "Album", colSize, "Artist", ">")
	for i := 1; rows.Next(); i++ {
		rows.Scan(&albumName, &artistName, &plays)
		artistName = wordWrap(artistName, colSize)
		albumName = wordWrap(albumName, colSize)
		table += fmt.Sprintf("%-*s %-*s %-s\n", colSize, albumName, colSize, artistName, plays)
		if i == 10 {
			break
		}
	}
	rows.Close()
	return table
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

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

<div align="center">
  <h1 align="center">KOLEKTI</h1>

  <p align="center">
    Show your music listening statistics.<br><br>
    <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/crdpa/kolekti?style=for-the-badge">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/crdpa/kolekti?style=for-the-badge">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/crdpa/kolekti?style=for-the-badge"><br>
  </p>
</div>

## ABOUT

Kolekti displays info about [Musyca's database](https://github.com/crdpa/musyca/) in the terminal.

![Kolekti demo](assets/kolekti_demo.gif)

## USAGE

```
kolekti -data artists -s 2020-02-01 -e 2022-13-10 -l 15 -db "path to database" --export "path to file"

  -data string
    songs, artists or albums. (default "songs")
  -db string
    Database path (default "$HOME/.config/musyca/database.db")
  -s string
    Start date (default "2000-12-30")
  -e string
    End date (default "2022-03-17")
  -l int
    Number of results to display. (default 10)
  -export string
    File path to export (default none)
```

Any other format or text used for the dates will be discarded. If you type "March" in 'Start Date' and '2022-03-01' in 'End Date', it will retrieve all the data from the beginning to '2022-03-01'.

## TODO

- [x] Show top artists, top songs and top albums
- [x] Implement start date and end date to retrieve data
- [x] Implement limit
- [x] Add option to export data to a file

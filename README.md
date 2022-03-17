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

For now it is just showing the top played songs.

![Kolekti demo](assets/kolekti_demo.gif)

Just type the dates in YYYY-MM-DD format, type the limit and see the output.

Any other format or text will be discarded. If you type "March" in 'Start Date' and '2022-03-01' in 'End Date', it will retrieve all the data from the beginning to '2022-03-01'. If you type anything other than numbers in the 'limit' field, it will default to top 10.

## TODO

- [x] Show top artists, top songs and top albums
- [x] Implement start date and end date to retrieve data
- [x] Implement limit
- [ ] Redesign the UI

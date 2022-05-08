# Entertainment API
#### API for data about movies, TV shows, books, and comics.

### Built with [Gin.](https://gin-gonic.com/)![](https://www.chetu.com/img/technology-logo/gin-gonic.png)
---
### Requirements
You will need Go / Golang version 1.14 or higher to run this application. (Preferably 1.18 or latest).
All dependencies included in [go.mod](https://github.com/oranges0da/entertainment-api/blob/main/go.mod) file.
___

## Endpoints

### Movies
### Search for movie info by title or IMDb id.
```
GET /movie/?title=2001+space+odyssey
GET /movie/?imdb=tt0120915
```
### Response Format:
```json
{
    "data": {
        "Title": "2001: A Space Odyssey",
        "Year": "1968",
        "Rated": "G",
        "Runtime": "149 min",
        "Genre": "Sci-Fi",
        "Director": "Stanley Kubrick",
        "Writer": "Stanley Kubrick",
        "Actors": "Keir Dullea,Gary Lockwood,William Sylvester,Daniel Richter",
        "Plot": "2001 is a story of evolution. Sometime in the distant past, someone or something nudged evolution by placing a monolith on Earth (presumably elsewhere throughout the universe as well). Evolution then enabled humankind to reach the moon's surface, where yet another monolith is found, one that signals the monolith placers that humankind has evolved that far. Now a race begins between computers (HAL) and human (Bowman) to reach the monolith placers. The winner will achieve the next step in evolution, whatever that may be.",
        "Awards": "Won 1 Oscar. 16 wins & 12 nominations total",
        "Poster": "https://m.media-amazon.com/images/M/MV5BMmNlYzRiNDctZWNhMi00MzI4LThkZTctMTUzMmZkMmFmNThmXkEyXkFqcGdeQXVyNzkwMjQ5NzM@._V1_SX300.jpg",
        "Ratings": [
            {"Source":"Internet Movie Database","Value":"8.3/10"},
            {"Source":"Rotten Tomatoes","Value":"92%"},
            {"Source":"Metacritic","Value":"84/100"}
        ],
        "Metascore": "84",
        "imdbRating": "8.3",
        "imdbID": "tt0062622",
        "DVD": "23 Oct 2007",
        "BoxOffice": "$26,617,553",
    },
    "errors": [],
}
```
# Entertainment API
#### API for data about movies, TV shows, books, and comics.

### Built with [Gin.](https://gin-gonic.com/)![](https://www.chetu.com/img/technology-logo/gin-gonic.png)
---
### Requirements
You will need Go / Golang version 1.14 or higher to run this application. (Preferably 1.18 or latest).
___

## Endpoints
___
### Movies
### Search for movie info by title or IMDb id.
```
GET /movie/?title=2001+space+odyssey
GET /movie/?imdb=tt0120915
```
### Response Format:
```json
{
    "title": "2001: A Space Odyssey",
    "year": "1968",
    "age_rating": "G",
    "runtime": "149 min",
    "genre": "Sci-Fi",
    "director": "Stanley Kubrick",
    "writer": "Stanley Kubrick",
    "cast": [
        "Keir Dullea",
        "Gary Lockwood",
        "William Sylvester",
        "Daniel Richter",
    ],
    "plot": "2001 is a story of evolution. Sometime in the distant past, someone or something nudged evolution by placing a monolith on Earth (presumably elsewhere throughout the universe as well). Evolution then enabled humankind to reach the moon's surface, where yet another monolith is found, one that signals the monolith placers that humankind has evolved that far. Now a race begins between computers (HAL) and human (Bowman) to reach the monolith placers. The winner will achieve the next step in evolution, whatever that may be.",
    "awards": "Won 1 Oscar. 16 wins & 12 nominations total",
    "poster": "https://m.media-amazon.com/images/M/MV5BMmNlYzRiNDctZWNhMi00MzI4LThkZTctMTUzMmZkMmFmNThmXkEyXkFqcGdeQXVyNzkwMjQ5NzM@._V1_SX300.jpg",
    "ratings": [
        {"Source":"Internet Movie Database","Value":"8.3/10"},
        {"Source":"Rotten Tomatoes","Value":"92%"},
        {"Source":"Metacritic","Value":"84/100"}
    ],
    "metascore": "84",
    "imdb_rating": "8.3",
}
```
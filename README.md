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
### TV Shows
### Search for a TV show info by title or IMDb id.
```
GET /tv/?title=breaking+bad
GET /tv/?imdb=tt0903747
```
### Response Format:
```json
{
    "data": {
        "Title": "Breaking Bad",
        "Year": "2008–2013",
        "Rated": "TV-MA",
        "Runtime": "49 min",
        "Genre": "Crime, Drama, Thriller",
        "Director": "N/A",
        "Writer": "Vince Gilligan",
        "Actors": "Bryan Cranston, Aaron Paul, Anna Gunn",
        "Plot": "A high school chemistry teacher diagnosed with inoperable lung cancer turns to manufacturing and selling methamphetamine in order to secure his family's future.",
        "Awards": "Won 16 Primetime Emmys. 154 wins & 247 nominations total",
        "Poster": "https://m.media-amazon.com/images/M/MV5BZTE2YWRlMmYtOGFkYy00MjcxLWJkNmQtNTJmNTZkZjVhZGE1XkEyXkFqcGdeQXVyMTMzNDExODE5._V1_SX300.jpg",
        "Ratings": [
            {
                "Source": "Internet Movie Database",
                "Value": "9.5/10"
            },
            {
                "Source": "Rotten Tomatoes",
                "Value": "96%"
            }
        ],
        "Metascore": "N/A",
        "imdbRating": "9.5",
        "imdbID": "tt0903747",
        "DVD": "",
        "BoxOffice": "",
        "totalSeasons": "5"
    },
    "errors": []
}
```
### Books
### Search for a book info by title.
```
GET /book/?title=dune
```
### Response Format:
```json
{
    "data": {
        "totalItems": 2003,
        "items": [
            {
                "volumeInfo": {
                    "title": "Dune",
                    "authors": [
                        "Frank Herbert"
                    ],
                    "publisher": "Penguin",
                    "publishedDate": "2005",
                    "description": "Follows the adventures of Paul Atreides, the son of a betrayed duke given up for dead on a treacherous desert planet and adopted by its fierce, nomadic people, who help him unravel his most unexpected destiny.",
                    "pageCount": 528,
                    "printType": "BOOK",
                    "averageRating": 4.5,
                    "maturityRating": "NOT_MATURE",
                    "imageLinks": {
                        "smallThumbnail": "http://books.google.com/books/content?id=B1hSG45JCX4C&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
                        "thumbnail": "http://books.google.com/books/content?id=B1hSG45JCX4C&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
                    },
                    "language": "en"
                }
            },
        ]
    "errors": [],
}
```
# Entertainment API
#### API for data about movies, TV shows, books, and comics.

### Built with [Gin.](https://gin-gonic.com/)  ![](https://www.chetu.com/img/technology-logo/gin-gonic.png)
---
### Requirements
You will need Go / Golang version 1.14 or higher to run this application. (Preferably 1.18 or latest).
___
## Movies
___
### Endpoints
/movie/:name (example: /movie/t?=2001+space+odyssey)
```json
{
    "title": "2001: A Space Odyssey",
    "year": "1968",
    "age_rating": "G",
    "runtime": "149 min",
    "genre": "Sci-Fi",
    "director": "Stanley Kubrick",
    "writer": "Stanley Kubrick",
    "actors": [
        "Keir Dullea",
        "Gary Lockwood",
        "William Sylvester",
        "Daniel Richter",
        "John Hurt",
        "David Bowman",
    ],

}
```
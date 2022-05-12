package handler

import (
	"github.com/gin-gonic/gin"
	book "github.com/oranges0da/entertainment-api/service/book"
	movie "github.com/oranges0da/entertainment-api/service/movie"
	music "github.com/oranges0da/entertainment-api/service/music"
	tv "github.com/oranges0da/entertainment-api/service/tv"
)

type Config struct {
	R *gin.Engine // for configurating root router
}

func NewHandler(c *Config) {
	movieGroup := c.R.Group("/movie") // movie endpoint
	{
		movieGroup.GET("/title", movie.GetMovie())
	}

	TVGroup := c.R.Group("/tv") // tv endpoint
	{
		TVGroup.GET("/title", tv.GetTV())
	}

	BookGroup := c.R.Group("/book") // book endpoint
	{
		BookGroup.GET("/title", book.GetBook())
		BookGroup.GET("/author", book.GetAuthor())
	}

	MusicGroup := c.R.Group("/music") // music endpoint
	{
		MusicGroup.GET("/song", music.GetSong())
		MusicGroup.GET("/artist", music.GetArtist())
		MusicGroup.GET("/album", music.GetAlbum())
	}
}

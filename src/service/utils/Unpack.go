package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	model "github.com/oranges0da/entertainment-api/model"
)

func Unpack[M model.Movie | model.TV | model.Music | model.Book](res *http.Response, model M) M {
	body, err := ioutil.ReadAll(res.Body) // read res body

	if err != nil {
		log.Fatal(err)
	}

	var data M

	err = json.Unmarshal(body, &data) // unmarshal body to data

	if err != nil {
		log.Fatal(err)
	}

	return data
}

package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/oranges0da/entertainment-api/model"
)

func UnpackTrack(res *http.Response) model.Music {
	body, err := ioutil.ReadAll(res.Body) // read res body

	if err != nil {
		log.Fatal(err)
	}

	var data model.Music

	err = json.Unmarshal(body, &data) // unmarshal body to data

	if err != nil {
		log.Fatal(err)
	}

	return data
}

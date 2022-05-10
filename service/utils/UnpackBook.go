package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/oranges0da/entertainment-api/model"
)

func UnpackBook(res *http.Response) model.Book {
	body, err := ioutil.ReadAll(res.Body) // read res body
	log.Printf("JSON body: " + string(body))

	if err != nil {
		log.Printf("Error reading body: " + err.Error())
	}

	var data model.Book

	err = json.Unmarshal(body, &data) // unmarshal body to data

	if err != nil {
		log.Fatal(err)
	}

	return data
}

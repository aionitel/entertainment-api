package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/oranges0da/entertainment-api/model"
)

func UnpackTV(res *http.Response) model.TV {
	body, err := ioutil.ReadAll(res.Body) // read res body

	if err != nil {
		panic(err)
	}

	var data model.TV

	err = json.Unmarshal(body, &data) // unmarshal body to data

	if err != nil {
		panic(err)
	}

	return data
}

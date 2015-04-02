package structs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Region struct {
	Id   int64  `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Type string `json:"type" xml:"type"`
}

type Regions struct {
	Regions []Region `json:"region" xml:"region"`
}

func (re Regions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(re)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
}

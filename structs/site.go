package structs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Sites struct {
	Sites []Site `json:"site" xml:"site"`
}

type Site struct {
	Id     int64  `json:"id" xml:"id"`
	Type   string `json:"type" xml:"type"`
	Name   string `json:"name" xml:"name"`
	CoOrds string `json:"coords" xml:"coords"`
}

func (s Sites) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(s)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
}

package structs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WorldConstructions struct {
	WorldConstructions []WorldConstruction `json:"world_construction" xml:"world_construction"`
}

type WorldConstruction struct {
	Id int64 `json:"id" xml:"id"`
}

func (wc WorldConstructions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(wc)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
}

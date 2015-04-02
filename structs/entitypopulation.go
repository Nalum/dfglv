package structs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EntityPopulations struct {
	EntityPopulations []EntityPopulation `json:"entity_population" xml:"entity_population"`
}

type EntityPopulation struct {
	Id int64 `json:"id" xml:"id"`
}

func (e EntityPopulations) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(e)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
}

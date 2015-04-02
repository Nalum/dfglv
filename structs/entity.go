package structs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Entities struct {
	Entities []Entity `json:"entity" xml:"entity"`
}

type Entity struct {
	Id   int64  `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

func (e Entities) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(e)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
}

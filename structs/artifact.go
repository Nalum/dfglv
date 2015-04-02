package structs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Artifact struct {
	Id   int64  `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Item string `json:"item" xml:"item"`
}

type Artifacts struct {
	Artifacts []Artifact `json:"artifact" xml:"artifact"`
}

func (a Artifacts) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(a)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
}

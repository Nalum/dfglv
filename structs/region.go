package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	path := r.URL.Path
	var data []byte
	var err error

	if path[strings.LastIndex(path, "/"):] == "/" {
		data, err = json.Marshal(re.Regions)
	} else {
		index, err := strconv.ParseInt(path[strings.LastIndex(path, "/")+1:], 10, 64)

		if err != nil {
			panic(err)
		}

		data, err = json.Marshal(re.Regions[index])
	}

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
	log.Println(r.URL)
}

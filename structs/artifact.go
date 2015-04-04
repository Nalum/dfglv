package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	path := r.URL.Path
	var data []byte
	var err error

	if path[strings.LastIndex(path, "/"):] == "/" {
		data, err = json.Marshal(a.Artifacts)
	} else {
		index, err := strconv.ParseInt(path[strings.LastIndex(path, "/")+1:], 10, 64)

		if err != nil {
			panic(err)
		}

		data, err = json.Marshal(a.Artifacts[index])
	}

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))

	log.Println(r.URL)
}

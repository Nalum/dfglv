package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	path := r.URL.Path
	var data []byte
	var err error

	if path[strings.LastIndex(path, "/"):] == "/" {
		data, err = json.Marshal(s.Sites)
	} else {
		index, err := strconv.ParseInt(path[strings.LastIndex(path, "/")+1:], 10, 64)

		if err != nil {
			panic(err)
		}

		data, err = json.Marshal(s.Sites[index])
	}

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
	log.Println(r.URL)
}

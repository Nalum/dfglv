package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UndergroundRegions struct {
	UndergroundRegions []UndergroundRegion `json:"underground_region" xml:"underground_region"`
}

type UndergroundRegion struct {
	Id    int64  `json:"id" xml:"id"`
	Type  string `json:"type" xml:"type"`
	Depth int    `json:"depth" xml:"depth"`
}

func (u UndergroundRegions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var data []byte
	var err error

	if path[strings.LastIndex(path, "/"):] == "/" {
		data, err = json.Marshal(u.UndergroundRegions)
	} else {
		index, err := strconv.ParseInt(path[strings.LastIndex(path, "/")+1:], 10, 64)

		if err != nil {
			panic(err)
		}

		data, err = json.Marshal(u.UndergroundRegions[index])
	}

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
	log.Println(r.URL)
}

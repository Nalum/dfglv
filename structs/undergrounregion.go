package structs

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	data, err := json.Marshal(u)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
}

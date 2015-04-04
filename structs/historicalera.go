package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type HistoricalEras struct {
	HistoricalEras []HistoricalEra `json:"historical_era" xml:"historical_era"`
}

type HistoricalEra struct {
	Name      string `json:"name" xml:"name"`
	StartYear int    `json:"start_year" xml:"start_year"`
}

func (h HistoricalEras) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var data []byte
	var err error

	if path[strings.LastIndex(path, "/"):] == "/" {
		data, err = json.Marshal(h.HistoricalEras)
	} else {
		index, err := strconv.ParseInt(path[strings.LastIndex(path, "/")+1:], 10, 64)

		if err != nil {
			panic(err)
		}

		data, err = json.Marshal(h.HistoricalEras[index])
	}

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
	log.Println(r.URL)
}

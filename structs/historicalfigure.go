package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type HistoricalFigures struct {
	HistoricalFigures []HistoricalFigure `json:"historical_figure" xml:"historical_figure"`
}

type HistoricalFigure struct {
	Id             int64        `json:"id" xml:"id"`
	Name           string       `json:"name" xml:"name"`
	Race           string       `json:"race" xml:"race"`
	Caste          string       `json:"caste" xml:"caste"`
	Appeared       int          `json:"appeared" xml:"appeared"`
	BirthYear      int          `json:"birth_year" xml:"birth_year"`
	DeathYear      int          `json:"death_year" xml:"death_year"`
	AssociatedType string       `json:"associated_type" xml:"associated_type"`
	EntityLink     []EntityLink `json:"entity_link" xml:"entity_link"`
	Spheres        []string     `json:"sphere" xml:"sphere"`
}

type EntityLink struct {
	LinkType string `json:"link_type" xml:"link_type"`
	EntityId int64  `json:"entity_id" xml:"entity_id"`
}

func (h HistoricalFigures) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var data []byte
	var err error

	if path[strings.LastIndex(path, "/"):] == "/" {
		data, err = json.Marshal(h.HistoricalFigures)
	} else {
		index, err := strconv.ParseInt(path[strings.LastIndex(path, "/")+1:], 10, 64)

		if err != nil {
			panic(err)
		}

		data, err = json.Marshal(h.HistoricalFigures[index])
	}

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
	log.Println(r.URL)
}

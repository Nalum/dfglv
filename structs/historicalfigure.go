// DFGLV - Dwarf Fortress GO Legends Viewer
// Copyright (C) 2015  Luke Mallon
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// HistoricalFigure is a struct containing information about a historical figure
type HistoricalFigure struct {
	ID             int64        `json:"id" xml:"id"`
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

// HistoricalFigures is a struct containing a list of HistoricalFigure
type HistoricalFigures struct {
	HistoricalFigures []HistoricalFigure `json:"historical_figure" xml:"historical_figure"`
}

// EntityLink is a struct containing a link to a specific Entity
type EntityLink struct {
	LinkType string `json:"link_type" xml:"link_type"`
	EntityID int64  `json:"entity_id" xml:"entity_id"`
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

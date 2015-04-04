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
)

// Legends is a struct which contains all information about a world
type Legends struct {
	Artifacts                  Artifacts                  `json:"artifacts" xml:"artifacts"`
	Entities                   Entities                   `json:"entities" xml:"entities"`
	EntityPopulations          EntityPopulations          `json:"entity_populations" xml:"entity_populations"`
	HistoricalEras             HistoricalEras             `json:"historical_eras" xml:"historical_eras"`
	HistoricalEvents           HistoricalEvents           `json:"historical_events" xml:"historical_events"`
	HistoricalEventCollections HistoricalEventCollections `json:"historical_event_collections" xml:"historical_event_collections"`
	HistoricalFigures          HistoricalFigures          `json:"historical_figures" xml:"historical_figures"`
	Regions                    Regions                    `json:"regions" xml:"regions"`
	Sites                      Sites                      `json:"sites" xml:"sites"`
	UndergroundRegions         UndergroundRegions         `json:"underground_regions" xml:"underground_regions"`
	WorldConstructions         WorldConstructions         `json:"world_constructions" xml:"world_constructions"`
}

func (l Legends) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(l)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
	log.Println(r.URL)
}

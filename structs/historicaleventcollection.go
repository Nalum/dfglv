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

// HistoricalEventCollection is a struct containing information about a collection
// of historical events
type HistoricalEventCollection struct {
	ID             int64   `json:"id" xml:"id"`
	StartYear      int     `json:"start_year" xml:"start_year"`
	StartSeconds72 int64   `json:"start_seconds72" xml:"start_seconds72"`
	EndYear        int     `json:"end_year" xml:"end_year"`
	EndSeconds72   int64   `json:"end_seconds72" xml:"end_seconds72"`
	Events         []int64 `json:"events" xml:"event"`
	Type           string  `json:"type" xml:"type"`
	ParentEventcol int64   `json:"parent_eventcol" xml:"parent_eventcol"`
	SubRegionID    int64   `json:"subregion_id" xml:"subregion_id"`
	FeatureLayerID int64   `json:"feature_layer_id" xml:"feature_layer_id"`
	SiteID         int64   `json:"site_id" xml:"site_id"`
	CoOrds         string  `json:"coords" xml:"coords"`
	DefendingEnID  int     `json:"defending_enid" xml:"defending_enid"`
	Ordinal        int     `json:"ordinal" xml:"ordinal"`
}

// HistoricalEventCollections is a struct containing a list of HistoricalEventCollection
type HistoricalEventCollections struct {
	HistoricalEventCollections []HistoricalEventCollection `json:"historical_event_collection" xml:"historical_event_collection"`
}

func (h HistoricalEventCollections) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var data []byte
	var err error

	if path[strings.LastIndex(path, "/"):] == "/" {
		data, err = json.Marshal(h.HistoricalEventCollections)
	} else {
		index, err := strconv.ParseInt(path[strings.LastIndex(path, "/")+1:], 10, 64)

		if err != nil {
			panic(err)
		}

		data, err = json.Marshal(h.HistoricalEventCollections[index])
	}

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
	log.Println(r.URL)
}

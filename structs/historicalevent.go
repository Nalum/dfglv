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

// HistoricalEvent is a struct containing information about historical events
type HistoricalEvent struct {
	ID             int64  `json:"id" xml:"id"`
	Year           int    `json:"year" xml:"year"`
	Seconds72      int64  `json:"seconds72" xml:"seconds72"`
	Type           string `json:"type" xml:"type"`
	HFID           int64  `json:"hfid" xml:"hfid"`
	State          string `json:"state" xml:"state"`
	SiteID         int64  `json:"site_id" xml:"site_id"`
	SubRegionID    int64  `json:"subregion_id" xml:"subregion_id"`
	FeatureLayerID int64  `json:"feature_layer_id" xml:"feature_layer_id"`
	CoOrds         string `json:"coords" xml:"coords"`
}

// HistoricalEvents is a struct containing a list of HistoricalEvent
type HistoricalEvents struct {
	HistoricalEvents []HistoricalEvent `json:"historical_event" xml:"historical_event"`
}

func (h HistoricalEvents) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var data []byte
	var err error

	if path[strings.LastIndex(path, "/"):] == "/" {
		data, err = json.Marshal(h.HistoricalEvents)
	} else {
		index, err := strconv.ParseInt(path[strings.LastIndex(path, "/")+1:], 10, 64)

		if err != nil {
			panic(err)
		}

		data, err = json.Marshal(h.HistoricalEvents[index])
	}

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
	log.Println(r.URL)
}

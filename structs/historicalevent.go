package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type HistoricalEvents struct {
	HistoricalEvents []HistoricalEvent `json:"historical_event" xml:"historical_event"`
}

type HistoricalEvent struct {
	Id             int64  `json:"id" xml:"id"`
	Year           int    `json:"year" xml:"year"`
	Seconds72      int64  `json:"seconds72" xml:"seconds72"`
	Type           string `json:"type" xml:"type"`
	HFId           int64  `json:"hfid" xml:"hfid"`
	State          string `json:"state" xml:"state"`
	SiteId         int64  `json:"site_id" xml:"site_id"`
	SubRegionId    int64  `json:"subregion_id" xml:"subregion_id"`
	FeatureLayerId int64  `json:"feature_layer_id" xml:"feature_layer_id"`
	CoOrds         string `json:"coords" xml:"coords"`
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

package structs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HistoricalEras struct {
	HistoricalEras []HistoricalEra `json:"historical_era" xml:"historical_era"`
}

type HistoricalEra struct {
	Name      string `json:"name" xml:"name"`
	StartYear int    `json:"start_year" xml:"start_year"`
}

func (h HistoricalEras) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(h)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(data))
}

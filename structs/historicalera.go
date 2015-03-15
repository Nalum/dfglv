package structs

type HistoricalEras struct {
	HistoricalEras []HistoricalEra `json:"historical_era" xml:"historical_era"`
}

type HistoricalEra struct {
	Name      string `json:"name" xml:"name"`
	StartYear int    `json:"start_year" xml:"start_year"`
	EndYear   int    `json:"end_year" xml:"end_year"`
}

package structs

type HistoricalEventCollections struct {
	HistoricalEventCollections []HistoricalEventCollection `json:"historical_event_collection" xml:"historical_event_collection"`
}

type HistoricalEventCollection struct {
	Id             int64   `json:"id" xml:"id"`
	StartYear      int     `json:"start_year" xml:"start_year"`
	StartSeconds72 int64   `json:"start_seconds72" xml:"start_seconds72"`
	EndYear        int     `json:"end_year" xml:"end_year"`
	EndSeconds72   int64   `json:"end_seconds72" xml:"end_seconds72"`
	Events         []int64 `json:"events" xml:"event"`
	Type           string  `json:"type" xml:"type"`
	ParentEventcol int64   `json:"parent_eventcol" xml:"parent_eventcol"`
	SubRegionId    int64   `json:"subregion_id" xml:"subregion_id"`
	FeatureLayerId int64   `json:"feature_layer_id" xml:"feature_layer_id"`
	SiteId         int64   `json:"site_id" xml:"site_id"`
	CoOrds         string  `json:"coords" xml:"coords"`
	DefendingEnid  int     `json:"defending_enid" xml:"defending_enid"`
	Ordinal        int     `json:"ordinal" xml:"ordinal"`
}

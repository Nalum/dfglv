package structs

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

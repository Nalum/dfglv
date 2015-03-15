package structs

type Legends struct {
	Artifacts         Artifacts         `json:"artifacts" xml:"artifacts"`
	Entities          Entities          `json:"entities" xml:"entities"`
	EntityPopulations EntityPopulations `json:"entity_populations" xml:"entity_populations"`
	HistoricalEras    HistoricalEras    `json:"historical_eras" xml:"historical_eras"`
	//HistoricalEvents           []HistoricalEvent
	//HistoricalEventCollections []HistoricalEventCollection
	//HistoricalFigures          []HistoricalFigure
	Regions Regions `json:"regions" xml:"regions"`
	Sites   Sites   `json:"sites" xml:"sites"`
	//UndergroundRegions         []UndergroundRegion
	//WorldConstructions         []WorldConstruction
}

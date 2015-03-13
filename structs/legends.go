package structs

type Legends struct {
	Artifacts []Artifact `json:"artifacts" xml:"artifacts"`
	//Entities                   []Entity
	//EntityPopulations          []EntityPopulation
	//HistoricalEras             []HistoricalEra
	//HistoricalEvents           []HistoricalEvent
	//HistoricalEventCollections []HistoricalEventCollection
	//HistoricalFigures          []HistoricalFigure
	Regions []Region `json:"regions" xml:"regions"`
	//Sites                      []Site
	//UndergroundRegions         []UndergroundRegion
	//WorldConstructions         []WorldConstruction
}

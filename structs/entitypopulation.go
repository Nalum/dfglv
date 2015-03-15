package structs

type EntityPopulations struct {
	EntityPopulations []EntityPopulation `json:"entity_population" xml:"entity_population"`
}

type EntityPopulation struct {
	Id uint64 `json:"id" xml:"id"`
}

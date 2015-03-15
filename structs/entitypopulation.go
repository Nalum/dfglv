package structs

type EntityPopulations struct {
	EntityPopulations []EntityPopulation `json:"entity_population" xml:"entity_population"`
}

type EntityPopulation struct {
	Id int64 `json:"id" xml:"id"`
}

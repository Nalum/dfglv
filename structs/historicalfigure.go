package structs

type HistoricalFigures struct {
	HistoricalFigures []HistoricalFigure `json:"historical_figure" xml:"historical_figure"`
}

type HistoricalFigure struct {
	Id             int64        `json:"id" xml:"id"`
	Name           string       `json:"name" xml:"name"`
	Race           string       `json:"race" xml:"race"`
	Caste          string       `json:"caste" xml:"caste"`
	Appeared       int          `json:"appeared" xml:"appeared"`
	BirthYear      int          `json:"birth_year" xml:"birth_year"`
	DeathYear      int          `json:"death_year" xml:"death_year"`
	AssociatedType string       `json:"associated_type" xml:"associated_type"`
	EntityLink     []EntityLink `json:"entity_link" xml:"entity_link"`
	Spheres        []string     `json:"sphere" xml:"sphere"`
}

type EntityLink struct {
	LinkType string `json:"link_type" xml:"link_type"`
	EntityId int64  `json:"entity_id" xml:"entity_id"`
}

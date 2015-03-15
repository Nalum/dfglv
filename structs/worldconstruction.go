package structs

type WorldConstructions struct {
	WorldConstructions []WorldConstruction `json:"world_construction" xml:"world_construction"`
}

type WorldConstruction struct {
	Id int64 `json:"id" xml:"id"`
}

package structs

type Entities struct {
	Entities []Entity `json:"entity" xml:"entity"`
}

type Entity struct {
	Id   int64  `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

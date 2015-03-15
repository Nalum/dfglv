package structs

type Sites struct {
	Sites []Site `json:"site" xml:"site"`
}

type Site struct {
	Id     int64  `json:"id" xml:"id"`
	Type   string `json:"type" xml:"type"`
	Name   string `json:"name" xml:"name"`
	CoOrds string `json:"coords" xml:"coords"`
}

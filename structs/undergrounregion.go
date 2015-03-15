package structs

type UndergroundRegions struct {
	UndergroundRegions []UndergroundRegion `json:"underground_region" xml:"underground_region"`
}

type UndergroundRegion struct {
	Id    int64  `json:"id" xml:"id"`
	Type  string `json:"type" xml:"type"`
	Depth int    `json:"depth" xml:"depth"`
}

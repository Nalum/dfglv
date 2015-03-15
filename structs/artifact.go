package structs

type Artifact struct {
	Id   int64  `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Item string `json:"item" xml:"item"`
}

type Artifacts struct {
	Artifacts []Artifact `json:"artifact" xml:"artifact"`
}

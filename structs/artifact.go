package structs

type Artifact struct {
	Id   uint64 `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Item string `json:"item" xml:"item"`
}

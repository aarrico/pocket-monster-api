package seed

type ApiData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ApiResp struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []ApiData `json:"results"`
}

type SubField struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type BaseStats struct {
	Value int      `json:"base_stat"`
	Stat  SubField `json:"stat"`
}

type Types struct {
	Slot int32    `json:"slot"`
	Type SubField `json:"type"`
}

type Pokemon struct {
	Name           string      `json:"name"`
	Height         int32       `json:"height"`
	Weight         int32       `json:"weight"`
	Species        SubField    `json:"species"`
	BaseExperience int32       `json:"base_experience"`
	IsDefault      bool        `json:"is_default"`
	SortOrder      int32       `json:"order"`
	BaseStats      []BaseStats `json:"stats"`
	Types          []Types     `json:"types"`
}

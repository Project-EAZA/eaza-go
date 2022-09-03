package course

type NameAndNumberRequest struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

type AbbrAndNumberRequest struct {
	Abbr   string `json:"abbr"`
	Number string `json:"number"`
}

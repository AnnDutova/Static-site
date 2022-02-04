package model

type Rewiew struct {
	Text           string `json:"text"`
	Grage          int    `json:"grade"`
	AuthorUsername string `json:"username"`
}

package seller

type MusicCard struct {
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Upload   string `json:"upload"`
	ID_genre int    `json:"id_genre"`
	ID_mus   int    `json:"id_mus"`
	ID_salon int    `json:"id_salon"`
	Count    int    `json:"count"`
	Price    int    `json:"price"`
	Sale     int    `json:"sale"`
}

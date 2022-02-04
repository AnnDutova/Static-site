package seller

type Seller struct {
	ID         int    `json:"id_seller"`
	ID_user    int    `json:"id_user"`
	ID_salon   int    `json:"id_salon"`
	Salon_name string `json:"title"`
}

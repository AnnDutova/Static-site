package apiserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/AnnDutova/static/internal/app/model/seller"
	model "github.com/AnnDutova/static/internal/app/model/user"
)

type ShopName struct {
	Name string
}

func (s *storeRouter) handleSellerPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/new_seller.html")
	data := ShopName{
		Name: s.seller.Salon_name,
	}
	tmpl.Execute(w, data)
}

func (s *storeRouter) handleCreateSellerAcc(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Username   string `json:"username"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		Salon_name string `json:"salon_name"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	u := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := s.store.User().Create(u); err != nil {
		log.Print(err)
		parseError(rw, r, http.StatusUnprocessableEntity, err)
		return
	}
	u.Sanitize()
	sel := &seller.Seller{
		ID_user:    u.ID,
		Salon_name: req.Salon_name,
	}
	if err := s.store.Seller().Create(sel); err != nil {
		parseError(rw, r, http.StatusUnprocessableEntity, err)
		return
	}
	respond(rw, r, http.StatusCreated, sel)
	s.seller = sel
}

func (s *storeRouter) logInSeller(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Salon_name string `json:"salon_name"`
		Password   string `json:"password"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	sel, u, err := s.store.Seller().FindBySalonName(req.Salon_name, req.Password)
	if err != nil || !u.ComparePassword(req.Password) {
		parseError(rw, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
		return
	}
	s.seller = sel
	respond(rw, r, http.StatusAccepted, sel)
}

type GenreArtistAnswer struct {
	Genre   []string
	Artists []string
}

func (s *storeRouter) handleInformationForSelector(rw http.ResponseWriter, r *http.Request) {
	log.Print("handleInformationForSelector")
	artist, err := s.store.Tools().GetAllArtists()
	if err != nil {
		log.Print(err)
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	genre, err := s.store.Tools().GetAllGenres()
	if err != nil {
		log.Print(err)
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	data := GenreArtistAnswer{
		Genre:   genre,
		Artists: artist,
	}
	respond(rw, r, http.StatusOK, data)
}

func (s *storeRouter) handleCreateMusicCard(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Title    string `json:"title"`
		Duration string `json:"duration"`
		Genre    string `json:"genre"`
		Artist   string `json:"artist"`
		Count    string `json:"count"`
		Price    string `json:"price"`
		Sale     string `json:"sale"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	log.Print("Decode new Music Card")
	mus_id, err := strconv.Atoi(req.Artist)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	mus_id++
	genre_id, err := strconv.Atoi(req.Genre)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	genre_id++
	count, err := strconv.Atoi(req.Count)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	price, err := strconv.Atoi(req.Price)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	sale, err := strconv.Atoi(req.Sale)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	card := &seller.MusicCard{
		Title:    req.Title,
		ID_mus:   mus_id,
		ID_genre: genre_id,
		Duration: req.Duration,
		ID_salon: s.seller.ID_salon,
		Count:    count,
		Price:    price,
		Sale:     sale,
	}
	if err := s.store.Seller().AddMusiCard(card); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	log.Print("Create new Music Card")
	respond(rw, r, http.StatusOK, card)
}

func (s *storeRouter) handleGenerateAllMusicCard(rw http.ResponseWriter, r *http.Request) {
	res, err := s.store.Seller().GetAllCompositions(s.seller)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	respond(rw, r, http.StatusOK, res)
}

func (s *storeRouter) handleCreateCollection(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Title    string `json:"title"`
		Artist   string `json:"artist"`
		ColTitle string `json:"collection_title"`
		ColPrice string `json:"collection_price"`
		ColSale  string `json:"collection_sale"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	price, err := strconv.Atoi(req.ColPrice)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	sale, err := strconv.Atoi(req.ColSale)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	songs := strings.Split(req.Title, ",")
	artists := strings.Split(req.Artist, ",")

	cInner := &model.CollectionInner{}
	mas := []model.CollectionCard{}
	for i := 0; i < len(songs)-1; i++ {
		colCard := &model.CollectionCard{}
		colCard.Author = artists[i]
		colCard.Song = songs[i]
		colCard.Salon = s.seller.Salon_name
		mas = append(mas, *colCard)
	}
	cInner.Collection = mas
	err = s.store.Seller().ComponateCollection(s.seller, cInner, price, sale, req.ColTitle)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	respond(rw, r, http.StatusOK, cInner)
}

func (s *storeRouter) handleCollectCollectionsSeller(rw http.ResponseWriter, r *http.Request) {
	res, err := s.store.Seller().GetAllCollections(s.seller)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	respond(rw, r, http.StatusOK, res)
}

func (s *storeRouter) openMusicCardSeller(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Author string `json:"author"`
		Song   string `json:"song"`
		Salon  string `json:"salon"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	card := &model.MusicCard{
		Author: req.Author,
		Song:   req.Song,
		Salon:  req.Salon,
	}
	rew, err := s.store.Rewiew().GetAllRewiewMusic(card)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	s.rewiew = rew
	respond(rw, r, http.StatusOK, rew)
}

type cardAnswerWithListenings struct {
	Author     string
	Title      string
	Salon      string
	Listenings int
	Sale       int
}

func (s *storeRouter) openCardSellerHTML(rw http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/composition_card_from_seller.html")
	count, err := s.store.Seller().GetCountOfListenings(s.rewiew.MusAuthor, s.rewiew.Song)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	sale, err := s.store.Seller().GetCurrentSale(s.rewiew.Song, s.seller.ID_salon)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	data := cardAnswerWithListenings{
		Author:     s.rewiew.MusAuthor,
		Title:      s.rewiew.Song,
		Salon:      s.rewiew.Salon,
		Listenings: count,
		Sale:       sale,
	}
	tmpl.Execute(rw, data)
}

func (s *storeRouter) handleCreateSale(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "static/create_sale_form.html")
}

func (s *storeRouter) handleAddSaleInDB(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Count string `json:"count"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	sale, err := strconv.Atoi(req.Count)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	if s.admin == nil {
		log.Print("Seller add sale")
		if err = s.store.Seller().EnterDiscount(s.rewiew.Song, s.seller.ID_salon, sale); err != nil {
			parseError(rw, r, http.StatusBadRequest, errInTransaction)
			return
		}
		respond(rw, r, http.StatusOK, s.seller)
	} else {
		log.Print("Admin add sale")
		if err = s.store.Administrator().SetSale(s.rewiew.MusAuthor, s.rewiew.Song, s.rewiew.Salon, sale); err != nil {
			parseError(rw, r, http.StatusBadRequest, err)
			return
		}
		respond(rw, r, http.StatusOK, s.admin)
	}
}

package apiserver

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	model "github.com/AnnDutova/static/internal/app/model/user"
)

func (s *storeRouter) handleCompositionPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/composition_card_for_admin.html")
	log.Print("handleCompositionPage", s.rewiew.MusAuthor, s.rewiew.Song)
	count, err := s.store.Tools().GetCountOfListenings(s.rewiew.MusAuthor, s.rewiew.Song)
	if err != nil {
		parseError(w, r, http.StatusBadRequest, err)
		return
	}
	sale, err := s.store.Tools().GetCurrentSale(s.rewiew.MusAuthor, s.rewiew.Song, s.rewiew.Salon)
	if err != nil {
		parseError(w, r, http.StatusBadRequest, err)
		return
	}
	data := cardAnswerWithListenings{
		Author:     s.rewiew.MusAuthor,
		Title:      s.rewiew.Song,
		Salon:      s.rewiew.Salon,
		Listenings: count,
		Sale:       sale,
	}
	tmpl.Execute(w, data)
}

func (s *storeRouter) handleCreateSaleFromShop(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "static/create_sale_form.html")
}

func (s *storeRouter) LogInAdministrator(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	admin, u, err := s.store.Administrator().Find(req.Username)
	if err != nil || !u.ComparePassword(req.Password) {
		parseError(rw, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
		return
	}
	s.admin = admin
	respond(rw, r, http.StatusAccepted, admin)
}

func (s *storeRouter) handleCompositionCard(rw http.ResponseWriter, r *http.Request) {
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
	log.Print(req.Author, req.Song, req.Salon)
	rew, err := s.store.Rewiew().GetAllRewiewMusic(card)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	s.rewiew = rew
	log.Print(s.rewiew)
	respond(rw, r, http.StatusOK, rew)
}

func (s *storeRouter) handleDeliteRewiew(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username"`
		Grade    int    `json:"grade"`
		Text     string `json:"text"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Print(err)
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	rew := &model.Rewiew{
		AuthorUsername: req.Username,
		Grage:          req.Grade,
		Text:           req.Text,
	}
	err := s.store.Rewiew().DeliteRewiew(rew, s.rewiew.Song, s.rewiew.MusAuthor, s.rewiew.Salon)
	if err != nil {
		log.Print(err)
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	respond(rw, r, http.StatusOK, nil)
}

func (s *storeRouter) handleBlockUser(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	u, err := s.store.User().FindByUsername(req.Username)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	if err = s.store.User().BlockUser(u); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	respond(rw, r, http.StatusOK, nil)
}

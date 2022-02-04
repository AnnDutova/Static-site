package apiserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"

	model "github.com/AnnDutova/static/internal/app/model/user"
)

func (s *storeRouter) handleCreateAcc(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
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
		parseError(rw, r, http.StatusUnprocessableEntity, err)
		return
	}
	u.Sanitize()
	if err := s.store.User().AddToCustomer(u); err != nil {
		parseError(rw, r, http.StatusInternalServerError, errNoAccount)
		return
	}
	if err := s.store.User().CreateCustomerStatus(u); err != nil {
		parseError(rw, r, http.StatusInternalServerError, errNoAccount)
		return
	}
	respond(rw, r, http.StatusCreated, u)
	s.user = u
}

func (s *storeRouter) logIn(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	u, err := s.store.User().FindByUsername(req.Username)
	if err != nil || !u.ComparePassword(req.Password) {
		parseError(rw, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
		return
	}
	res, err := s.store.User().AvailableProfile(u)
	if err != nil {
		parseError(rw, r, http.StatusInternalServerError, err)
		return
	}
	if res {
		s.user = u
		err = s.store.User().GetCurrentAccount(u)
		log.Print(err)
		if err != nil {
			parseError(rw, r, http.StatusInternalServerError, errNoAccount)
			return
		}
		respond(rw, r, http.StatusAccepted, u)
	} else {
		parseError(rw, r, http.StatusInternalServerError, errNoAvailableProfile)
		return
	}
}

type UserName struct {
	Name  string
	Money int
}

func (s *storeRouter) handleUserPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/new_user.html")

	data := UserName{
		Name:  s.user.Username,
		Money: s.user.Money,
	}
	tmpl.Execute(w, data)
}

type userView struct {
	Name       string
	Bucket     *model.Bucket
	Music      *model.MusicInner
	Collection *model.CollectionInner
	Total      int
}

func (s *storeRouter) handleCollectInformation(w http.ResponseWriter, r *http.Request) {
	bucket, total, err := s.store.User().FindBucketCondition(s.user)
	if err != nil {
		parseError(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	mus, err := s.store.User().FindMusicContainer(s.user)
	if err != nil {
		parseError(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	col, err := s.store.User().FindCollectionContainer(s.user)
	if err != nil {
		parseError(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	data := userView{
		Name:       s.user.Username,
		Bucket:     bucket,
		Music:      mus,
		Collection: col,
		Total:      total,
	}
	respond(w, r, http.StatusOK, data)
}

func (s *storeRouter) handleSendRewiew(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Rewiew string `json:"rewiew"`
		Author string `json:"author"`
		Song   string `json:"song"`
		Salon  string `json:"salon"`
		Grade  int    `json:"grade"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	log.Print("Decode rewiew")
	log.Print(req.Author, req.Song, req.Salon)
	card := &model.MusicCard{
		Author: req.Author,
		Song:   req.Song,
		Salon:  req.Salon,
	}
	if err := s.store.Rewiew().CreateMusicRewiew(card, s.user, req.Rewiew, req.Grade); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	log.Print("Create rewiew")
	respond(rw, r, http.StatusOK, nil)
}

func (s *storeRouter) replenishWallet(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Count string `json:"count"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	money, err := strconv.Atoi(req.Count)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	if err = s.store.User().GetTransaction(s.user, money); err != nil {
		parseError(rw, r, http.StatusBadRequest, errInTransaction)
		return
	}
	if err = s.store.User().GetCurrentAccount(s.user); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	respond(rw, r, http.StatusOK, s.user)
}

type cardAnswer struct {
	Author     string
	Title      string
	Salon      string
	Listenings int
	Price      int
}

func (s *storeRouter) handleCardPage(rw http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/composition_card.html")
	count, err := s.store.Tools().GetCountOfListenings(s.rewiew.MusAuthor, s.rewiew.Song)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	price, err := s.store.Tools().GetComposePrice(s.rewiew.MusAuthor, s.rewiew.Song, s.rewiew.Salon)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	data := cardAnswer{
		Author:     s.rewiew.MusAuthor,
		Title:      s.rewiew.Song,
		Salon:      s.rewiew.Salon,
		Listenings: count,
		Price:      price,
	}
	tmpl.Execute(rw, data)
}

func (s *storeRouter) handleCardPageFromMarket(rw http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/composition_card_from_market.html")
	count, err := s.store.Tools().GetCountOfListenings(s.rewiew.MusAuthor, s.rewiew.Song)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	price, err := s.store.Tools().GetComposePrice(s.rewiew.MusAuthor, s.rewiew.Song, s.rewiew.Salon)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	data := cardAnswer{
		Author:     s.rewiew.MusAuthor,
		Title:      s.rewiew.Song,
		Salon:      s.rewiew.Salon,
		Listenings: count,
		Price:      price,
	}
	tmpl.Execute(rw, data)
}

func (s *storeRouter) handleGetPreferences(rw http.ResponseWriter, r *http.Request) {
	preferences, err := s.store.User().GetPreferences(s.user)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	respond(rw, r, http.StatusOK, preferences)
}

func (s *storeRouter) handleAddPreference(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Preferences string `json:"preferences"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	log.Print(req.Preferences)
	err := s.store.User().AddPreferences(s.user, req.Preferences)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	respond(rw, r, http.StatusOK, nil)
}

func (s *storeRouter) handleGenerateRecomendation(rw http.ResponseWriter, r *http.Request) {
	log.Print("generateRecomendation in server ")
	list, err := s.store.User().GetPreferences(s.user)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	col, err := s.store.Tools().GetRecomendation(list)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	log.Print("Server return Recomendation")
	respond(rw, r, http.StatusOK, col)
}

func (s *storeRouter) handleMusicCompositionCard(rw http.ResponseWriter, r *http.Request) {
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

func (s *storeRouter) handleCollectionCompositionCard(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Title string `json:"title"`
		Value string `json:"value"`
		Salon string `json:"salon"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	card := &model.CollectionCard{
		Author: req.Value,
		Song:   req.Title,
		Salon:  req.Salon,
	}
	//rew, err := s.store.Rewiew().GetAllRewiewCollections(card)
	/*if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}*/
	respond(rw, r, http.StatusOK, card)
}

func (s *storeRouter) handleBucketCard(rw http.ResponseWriter, r *http.Request) {
	log.Print(json.NewDecoder(r.Body))
}

func (s *storeRouter) handleAddMusicCardToBucketFromCollection(rw http.ResponseWriter, r *http.Request) {
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
	mus := &model.MusicCard{
		Author: req.Author,
		Song:   req.Song,
		Salon:  req.Salon,
	}
	log.Print("Add to bucket ", mus.Author, " ", mus.Song, " ", mus.Salon)
	if err := s.store.User().AddToBucket(s.user, mus); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	log.Print("Added to bucket ", mus.Author, " ", mus.Song, " ", mus.Salon)
	respond(rw, r, http.StatusOK, nil)
}

func (s *storeRouter) handleAddAllCollectionToBucket(rw http.ResponseWriter, r *http.Request) {
	type request struct {
		Title string `json:"title"`
		Salon string `json:"salon"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	log.Print(req.Title, req.Salon, " Try to add collectionto bucket")
	mus := &model.CollectionCard{
		Song:  req.Title,
		Salon: req.Salon,
	}
	if err := s.store.User().AddCollectionToBucket(s.user, mus); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	log.Print(req.Title, req.Salon, " added collection to bucket")
	respond(rw, r, http.StatusOK, nil)
}

func (s *storeRouter) handleCollectBucket(rw http.ResponseWriter, r *http.Request) {
	bucket, _, err := s.store.User().FindBucketCondition(s.user)
	if err != nil {
		parseError(rw, r, http.StatusUnprocessableEntity, err)
		return
	}
	respond(rw, r, http.StatusOK, bucket)
}

func (s *storeRouter) handleDeliteFromBucket(rw http.ResponseWriter, r *http.Request) {
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
	mus := &model.MusicCard{
		Author: req.Author,
		Song:   req.Song,
		Salon:  req.Salon,
	}
	if err := s.store.User().DeliteFromBucket(s.user, mus); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	respond(rw, r, http.StatusOK, nil)
}

func (s *storeRouter) handleBuyFromBucketCard(rw http.ResponseWriter, r *http.Request) {
	_, total, err := s.store.User().FindBucketCondition(s.user)
	if err != nil {
		parseError(rw, r, http.StatusUnprocessableEntity, err)
		return
	}
	if s.user.Money >= total {
		log.Print("enought money")
		if err := s.store.User().BuyAllBucket(s.user); err != nil {
			parseError(rw, r, http.StatusBadRequest, err)
			return
		}
		if err := s.store.User().GetBuyTransaction(s.user, total); err != nil {
			parseError(rw, r, http.StatusBadRequest, err)
			return
		}
		if err := s.store.User().GetCurrentAccount(s.user); err != nil {
			parseError(rw, r, http.StatusBadRequest, err)
			return
		}
		respond(rw, r, http.StatusOK, err)
	} else {
		log.Print("No money")
		parseError(rw, r, http.StatusUnprocessableEntity, errNoMoney)
		return
	}
}

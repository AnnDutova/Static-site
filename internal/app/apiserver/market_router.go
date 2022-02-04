package apiserver

import (
	"net/http"

	"github.com/AnnDutova/static/internal/app/model/seller"
	model "github.com/AnnDutova/static/internal/app/model/user"
)

type MarketAnswer struct {
	IsAuthorised *model.IsAuthorized
	Genres       []string
	MaxPrice     int
	MinPrice     int
	Collections  *seller.Collection
	Tracks       *model.MusicInner
}

func (s *storeRouter) handleGenerateMarketPage(rw http.ResponseWriter, r *http.Request) {
	isAuth := &model.IsAuthorized{}
	if s.user != nil {
		who := &model.WhoIs{
			IsAuth: true,
			ID:     s.user.ID,
		}
		isAuth.IsCustomer = who
	} else if s.seller != nil {
		who := &model.WhoIs{
			IsAuth: true,
			ID:     s.seller.ID,
		}
		isAuth.IsSeller = who
	} else if s.admin != nil {
		who := &model.WhoIs{
			IsAuth: true,
			ID:     s.admin.ID,
		}
		isAuth.IsAdministrator = who

	} else {
		isAuth = nil
	}
	genres, err := s.store.Tools().GetAllGenres()
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, errInTransaction)
		return
	}
	max, min, err := s.store.Tools().GetAllMaxAndMinPrice()
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, errInTransaction)
		return
	}
	col, err := s.store.Tools().GetAllCollections()
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, errInTransaction)
		return
	}
	songs, err := s.store.Tools().GetAllCompositions()
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, errInTransaction)
		return
	}
	data := &MarketAnswer{
		IsAuthorised: isAuth,
		Genres:       genres,
		MaxPrice:     max,
		MinPrice:     min,
		Collections:  col,
		Tracks:       songs,
	}
	respond(rw, r, http.StatusOK, data)
}

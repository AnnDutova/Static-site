package apiserver

import (
	"net/http"

	tools "github.com/AnnDutova/static/internal/app/model"
	"github.com/AnnDutova/static/internal/app/model/seller"
	model "github.com/AnnDutova/static/internal/app/model/user"
)

type GalleryAnswer struct {
	IsAuthorised *model.IsAuthorized
	Collections  *seller.Collection
	Tracks       *model.MusicInner
	Values       *tools.CollectionValueInner
}

func (s *storeRouter) handleGenerateGalleryPage(rw http.ResponseWriter, r *http.Request) {
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
	values, err := s.store.Tools().GetAllValuesCollections()
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, errInTransaction)
		return
	}
	data := &GalleryAnswer{
		IsAuthorised: isAuth,
		Collections:  col,
		Tracks:       songs,
		Values:       values,
	}
	respond(rw, r, http.StatusOK, data)
}

package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AnnDutova/static/internal/app/model/administrator"
	"github.com/AnnDutova/static/internal/app/model/seller"
	model "github.com/AnnDutova/static/internal/app/model/user"
	"github.com/AnnDutova/static/internal/app/store"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("Not Authenticated")
	errNoAccount                = errors.New("Cant found your wallet")
	errInTransaction            = errors.New("Cant get a transaction operation")
	errNoAvailableProfile       = errors.New("Sorry, this account was blocked due to a violation of community rules")
	errNoMoney                  = errors.New("You dont have enough money")
)

type storeRouter struct {
	store  store.Store
	user   *model.User
	rewiew *model.RewiewInner
	seller *seller.Seller
	admin  *administrator.Administrator
}

func (s *storeRouter) initialize(store store.Store, cookieKey int) {
	s.store = store
}

func (s *storeRouter) handleExit(w http.ResponseWriter, r *http.Request) {
	s.user = nil
	s.admin = nil
	s.seller = nil
	http.ServeFile(w, r, "static/new_home.html")
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/new_home.html")
}
func handleAuthPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/new_sign.html")
}
func handleAuthSellerPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/new_sign_seller.html")
}
func handleLogInPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/new_login.html")
}
func handleLogInSellerPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/new_login_seller.html")
}

func handleLogInAdministratorPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/new_login_administrator.html")
}

func handleGalleryPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/new_gallery.html")
}
func handleMarketPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/new_market.html")
}
func (s *storeRouter) handleFullOfWallet(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "static/wallet_form.html")
}

func (s *storeRouter) handleAllRewiews(rw http.ResponseWriter, r *http.Request) {
	respond(rw, r, http.StatusOK, s.rewiew.Rewiews)
}

func configureRouters(store store.Store) {
	s := &storeRouter{
		store: store,
		user:  nil,
	}
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/auth", handleAuthPage)
	http.HandleFunc("/log", handleLogInPage)
	http.HandleFunc("/seller_auth", handleAuthSellerPage)
	http.HandleFunc("/seller_log", handleLogInSellerPage)
	http.HandleFunc("/admin_log", handleLogInAdministratorPage)
	http.HandleFunc("/gallery", handleGalleryPage)
	http.HandleFunc("/market", handleMarketPage)
	http.HandleFunc("/user", s.handleUserPage)
	http.HandleFunc("/seller", s.handleSellerPage)
	http.HandleFunc("/card", s.handleCardPage)
	http.HandleFunc("/card_", s.handleCardPageFromMarket)
	http.HandleFunc("/createAccount", s.handleCreateAcc)
	http.HandleFunc("/createSellerAccount", s.handleCreateSellerAcc)
	http.HandleFunc("/loginAccount", s.logIn)
	http.HandleFunc("/loginSellerAccount", s.logInSeller)
	http.HandleFunc("/loginAdministratorAccount", s.LogInAdministrator)
	http.HandleFunc("/sendRew", s.handleSendRewiew)
	http.HandleFunc("/collectInformation", s.handleCollectInformation)
	http.HandleFunc("/wallet", s.handleFullOfWallet)
	http.HandleFunc("/replenishWallet", s.replenishWallet)
	http.HandleFunc("/musicCard", s.handleMusicCompositionCard)
	http.HandleFunc("/collectionCard", s.handleCollectionCompositionCard)
	http.HandleFunc("/bucketCard", s.handleBucketCard)
	http.HandleFunc("/getAllRewiews", s.handleAllRewiews)
	http.HandleFunc("/getArtistGenres", s.handleInformationForSelector)
	http.HandleFunc("/createMusicCard", s.handleCreateMusicCard)
	http.HandleFunc("/collectMusicinSeller", s.handleGenerateAllMusicCard)
	http.HandleFunc("/createCollection", s.handleCreateCollection)
	http.HandleFunc("/collectCollectionsSeller", s.handleCollectCollectionsSeller)
	http.HandleFunc("/musicCardSeller", s.openMusicCardSeller)
	http.HandleFunc("/cardSeller", s.openCardSellerHTML)
	http.HandleFunc("/createSale", s.handleCreateSale)
	http.HandleFunc("/addSaleInDB", s.handleAddSaleInDB)
	http.HandleFunc("/generateGalleryPage", s.handleGenerateGalleryPage)
	http.HandleFunc("/generateMarketPage", s.handleGenerateMarketPage)
	http.HandleFunc("/addMusicCardToBucketFromCollection", s.handleAddMusicCardToBucketFromCollection)
	http.HandleFunc("/addAllCollectionToBucket", s.handleAddAllCollectionToBucket)
	http.HandleFunc("/collectBucket", s.handleCollectBucket)
	http.HandleFunc("/deliteFromBucket", s.handleDeliteFromBucket)
	http.HandleFunc("/getPreferences", s.handleGetPreferences)
	http.HandleFunc("/AddPreference", s.handleAddPreference)

	http.HandleFunc("/openCompositionPageFromAdministrator", s.handleCompositionCard)
	http.HandleFunc("/cardFromAdmin", s.handleCompositionPage)
	http.HandleFunc("/deliteRewiew", s.handleDeliteRewiew)
	http.HandleFunc("/blockUser", s.handleBlockUser)
	http.HandleFunc("/createSaleFromShop", s.handleCreateSaleFromShop)
	http.HandleFunc("/buyFromBucketCard", s.handleBuyFromBucketCard)
	http.HandleFunc("/generateRecomendation", s.handleGenerateRecomendation)

	http.HandleFunc("/exit", s.handleExit)
}

func parseError(w http.ResponseWriter, r *http.Request, code int, err error) {
	respond(w, r, code, map[string]string{"error": err.Error()})
}
func respond(w http.ResponseWriter, r *http.Request, code int, date interface{}) {
	w.WriteHeader(code)
	if date != nil {
		json.NewEncoder(w).Encode(date)
	}
}

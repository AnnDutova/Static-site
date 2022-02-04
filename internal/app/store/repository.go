package store

import (
	tools "github.com/AnnDutova/static/internal/app/model"
	"github.com/AnnDutova/static/internal/app/model/administrator"
	seller "github.com/AnnDutova/static/internal/app/model/seller"
	model "github.com/AnnDutova/static/internal/app/model/user"
)

type UserRepository interface {
	Create(*model.User) error
	FindByUsername(string) (*model.User, error)
	FindBucketCondition(*model.User) (*model.Bucket, int, error)
	FindMusicContainer(*model.User) (*model.MusicInner, error)
	FindCollectionContainer(*model.User) (*model.CollectionInner, error)
	GetCurrentAccount(*model.User) error
	GetBuyTransaction(*model.User, int) error
	GetTransaction(*model.User, int) error
	AddToBucket(*model.User, *model.MusicCard) error
	AddCollectionToBucket(*model.User, *model.CollectionCard) error
	DeliteFromBucket(*model.User, *model.MusicCard) error
	WhoAmI(*model.User) (*model.IsAuthorized, error)
	BlockUser(*model.User) error
	BuyAllBucket(u *model.User) error
	GetPreferences(*model.User) ([]string, error)
	AddPreferences(*model.User, string) error
	AvailableProfile(*model.User) (bool, error)
	AddToCustomer(u *model.User) error
	CreateCustomerStatus(u *model.User) error
}

type RewiewRepository interface {
	CreateMusicRewiew(*model.MusicCard, *model.User, string, int) error
	GetAllRewiewMusic(*model.MusicCard) (*model.RewiewInner, error)
	DeliteRewiew(*model.Rewiew, string, string, string) error
}
type SellerRepository interface {
	Create(*seller.Seller) error
	FindBySalonName(string, string) (*seller.Seller, *model.User, error)
	AddMusiCard(*seller.MusicCard) error
	GetAllCompositions(*seller.Seller) (*model.MusicInner, error)
	ComponateCollection(*seller.Seller, *model.CollectionInner, int, int, string) error
	GetAllCollections(*seller.Seller) (*seller.Collection, error)
	GetAllGenres() ([]string, error)
	GetAllArtists() ([]string, error)
	GetCountOfListenings(string, string) (int, error)
	EnterDiscount(string, int, int) error
	GetCurrentSale(string, int) (int, error)
}

type ToolsRepository interface {
	GetAllCompositions() (*model.MusicInner, error)
	GetAllCollections() (*seller.Collection, error)
	GetAllValuesCollections() (*tools.CollectionValueInner, error)
	GetAllGenres() ([]string, error)
	GetAllArtists() ([]string, error)
	GetAllMaxAndMinPrice() (int, int, error)
	GetCountOfListenings(string, string) (int, error)
	GetComposePrice(string, string, string) (int, error)
	GetRecomendation([]string) (*seller.Collection, error)
	GetCurrentSale(string, string, string) (int, error)
}

type AdministratorRepository interface {
	Find(string) (*administrator.Administrator, *model.User, error)
	SetSale(string, string, string, int) error
}

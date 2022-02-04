package store

type Store interface {
	User() UserRepository
	Rewiew() RewiewRepository
	Seller() SellerRepository
	Tools() ToolsRepository
	Administrator() AdministratorRepository
}

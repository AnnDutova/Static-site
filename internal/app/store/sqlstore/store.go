package sqlstore

import (
	"database/sql"

	"github.com/AnnDutova/static/internal/app/store"
)

// Store ...
type Store struct {
	db               *sql.DB
	userRepository   *UserRepository
	rewiewRepository *RewiewRepository
	sellerRepository *SellerRepository
	toolsRepository  *ToolsRepository
	adminRepository  *AdministratorRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
func (s *Store) Rewiew() store.RewiewRepository {
	if s.rewiewRepository != nil {
		return s.rewiewRepository
	}

	s.rewiewRepository = &RewiewRepository{
		store: s,
	}

	return s.rewiewRepository
}

func (s *Store) Seller() store.SellerRepository {
	if s.rewiewRepository != nil {
		return s.sellerRepository
	}

	s.sellerRepository = &SellerRepository{
		store: s,
	}

	return s.sellerRepository
}

func (s *Store) Tools() store.ToolsRepository {
	if s.toolsRepository != nil {
		return s.toolsRepository
	}

	s.toolsRepository = &ToolsRepository{
		store: s,
	}
	return s.toolsRepository
}
func (s *Store) Administrator() store.AdministratorRepository {
	if s.adminRepository != nil {
		return s.adminRepository
	}

	s.adminRepository = &AdministratorRepository{
		store: s,
	}
	return s.adminRepository
}

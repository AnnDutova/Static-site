package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/AnnDutova/static/internal/app/model/administrator"
	model "github.com/AnnDutova/static/internal/app/model/user"
	"github.com/AnnDutova/static/internal/app/store"
)

type AdministratorRepository struct {
	store *Store
}

func (r *AdministratorRepository) Find(username string) (*administrator.Administrator, *model.User, error) {
	u := &model.User{}
	sqlRequest := fmt.Sprintf(`SELECT id_user, email, encrypted_password FROM users WHERE username="%s"`, username)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, store.ErrRecordNotFound
		}
		return nil, nil, err
	}
	admin := &administrator.Administrator{}
	sqlRequest = fmt.Sprintf(`select id_admin from administrator where id_user='%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&admin.ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, store.ErrRecordNotFound
		}
		return nil, nil, err
	}
	admin.ID_user = u.ID
	return admin, u, nil
}
func (r *AdministratorRepository) SetSale(author, song, salon string, sale int) error {
	sqlStatment := fmt.Sprintf(`update records_in_salon set sale = '%d' where 
	id_rec = (select id_rec from records
	left join musician using(id_mus) where title = '%s' and mus_name='%s') 
	and id_salon = (select id_salon from salon where title = '%s')`, sale, song, author, salon)
	_, err := r.store.db.Exec(sqlStatment)
	if err != nil {
		return err
	}
	return nil
}

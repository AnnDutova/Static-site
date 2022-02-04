package sqlstore

import (
	"database/sql"
	"fmt"
	"log"

	model "github.com/AnnDutova/static/internal/app/model/user"
)

type RewiewRepository struct {
	store *Store
}

func (r *RewiewRepository) CreateMusicRewiew(c *model.MusicCard, u *model.User, rew string, grade int) error {
	var id_mus, id_rec, id_customer int
	sqlRequest := fmt.Sprintf(`select id_mus from musician where mus_name='%s'`, c.Author)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&id_mus); err != nil {
		log.Print(1, err)
		return err
	}
	sqlRequest = fmt.Sprintf(`select id_customer from customer where id_user='%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&id_customer); err != nil {
		log.Print(2, err)
		return err
	}
	sqlRequest = fmt.Sprintf(`select records.id_rec from records_in_salon 
	left join salon using(id_salon), records 
	where records.id_rec = records_in_salon.id_rec and id_mus='%d' and records.title='%s'`, id_mus, c.Song)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&id_rec); err != nil {
		log.Print(3, err)
		return err
	}
	sqlRequest = fmt.Sprintf(`insert into rewiews(text, grade, id_mus, id_rec, id_customer) value('%s', %d, %d, %d, %d)`,
		rew, grade, id_mus, id_rec, id_customer)
	if _, err := r.store.db.Exec(sqlRequest); err != nil {
		log.Print(4, err)
		return err
	}
	return nil
}

func (r *RewiewRepository) GetAllRewiewMusic(c *model.MusicCard) (*model.RewiewInner, error) {
	rInner := &model.RewiewInner{}
	mas := []model.Rewiew{}
	sqlRequest := fmt.Sprintf(`call find_rewiew_records('%s', '%s')`, c.Song, c.Author)
	rows, err := r.store.db.Query(sqlRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Print(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		el := &model.Rewiew{}
		if err := rows.Scan(&el.Text, &el.Grage, &el.AuthorUsername); err != nil {
			return nil, err
		}
		mas = append(mas, *el)
	}
	rInner.Rewiews = mas
	rInner.MusAuthor = c.Author
	rInner.Song = c.Song
	rInner.Salon = c.Salon
	return rInner, nil
}

func (r *RewiewRepository) DeliteRewiew(rew *model.Rewiew, song, author, salon string) error {
	sqlRequest := fmt.Sprintf(`Delete from rewiews where text='%s' and grade='%d' and id_mus=(
		select id_mus from musician where mus_name ="%s") and id_rec = (
		select id_rec from records where title='%s') and id_customer=(
		select id_customer from customer,users where users.id_user= customer.id_user and username ="%s")`,
		rew.Text, rew.Grage, author, song, rew.AuthorUsername)
	_, err := r.store.db.Exec(sqlRequest)
	if err != nil {
		return err
	}
	return nil
}

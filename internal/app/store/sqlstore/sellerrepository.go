package sqlstore

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/AnnDutova/static/internal/app/model/seller"
	model "github.com/AnnDutova/static/internal/app/model/user"
	"github.com/AnnDutova/static/internal/app/store"
)

type SellerRepository struct {
	store *Store
}

func (s *SellerRepository) Create(sel *seller.Seller) error {
	sqlStatment := `INSERT INTO salon(title) VALUE (?);`
	res, err := s.store.db.Exec(sqlStatment, sel.Salon_name)
	if err != nil {
		return err
	}
	log.Print(err)
	salon_id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	log.Print(err)
	sel.ID_salon = int(salon_id)
	sqlStatment = `INSERT INTO seller(id_user, id_salon) VALUE (?, ?);`
	res, err = s.store.db.Exec(sqlStatment, sel.ID_user, sel.ID_salon)
	if err != nil {
		return err
	}
	log.Print(err)
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	log.Print(err)
	sel.ID = int(id)
	return nil
}

func (s *SellerRepository) FindBySalonName(salon_name, password string) (*seller.Seller, *model.User, error) {
	sel := &seller.Seller{}
	sqlRequest := fmt.Sprintf(`SELECT id_seller, id_user, id_salon FROM salon left join seller using(id_salon) WHERE salon.title="%s"`, salon_name)
	if err := s.store.db.QueryRow(sqlRequest).Scan(&sel.ID, &sel.ID_user, &sel.ID_salon); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, store.ErrRecordNotFound
		}
		return nil, nil, err
	}
	sel.Salon_name = salon_name
	u := &model.User{}
	sqlRequest = fmt.Sprintf(`SELECT username, email, encrypted_password FROM users WHERE id_user="%d"`, sel.ID_user)
	if err := s.store.db.QueryRow(sqlRequest).Scan(&u.Username, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, store.ErrRecordNotFound
		}
		return nil, nil, err
	}
	return sel, u, nil
}

func (s *SellerRepository) GetAllGenres() ([]string, error) {
	var mas []string
	rows, err := s.store.db.Query(`select name from genre`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var el string
		if err := rows.Scan(&el); err != nil {
			return nil, err
		}
		mas = append(mas, el)
	}
	return mas, nil
}

func (s *SellerRepository) GetAllArtists() ([]string, error) {
	var mas []string
	rows, err := s.store.db.Query(`select mus_name from musician`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var el string
		if err := rows.Scan(&el); err != nil {
			return nil, err
		}
		mas = append(mas, el)
	}
	return mas, nil
}

func (s *SellerRepository) AddMusiCard(c *seller.MusicCard) error {
	sqlStatment := `INSERT INTO records(title, duration, upload, id_genre, id_mus, bought) VALUE (?, ?, ?, ?, ?, ?);`
	res, err := s.store.db.Exec(sqlStatment, c.Title, c.Duration, time.Now(), c.ID_genre, c.ID_mus, 0)
	if err != nil {
		return err
	}
	record_id, err := res.LastInsertId()
	log.Print("Add to records")
	if err != nil {
		return err
	}
	sqlStatment = `INSERT INTO records_in_salon(id_rec, id_salon, count_, price, sale) VALUE (?, ?, ?, ?, ?);`
	res, err = s.store.db.Exec(sqlStatment, record_id, c.ID_salon, c.Count, c.Price, c.Sale)
	if err != nil {
		return err
	}
	return nil
}
func (s *SellerRepository) ComponateCollection(sel *seller.Seller, cInner *model.CollectionInner, price int, sale int, title string) error {
	var count int
	var record_id []int
	for i := 0; i < len(cInner.Collection); i++ {
		var temp, record int
		sqlRequest := fmt.Sprintf(`select count_, id_rec from records_in_salon 
		left Join salon using(id_salon) 
		where salon.title = "%s" and id_rec in (Select id_rec from musician 
		left join records using(id_mus)
		where records.title = '%s' and musician.mus_name = '%s');`,
			sel.Salon_name, cInner.Collection[i].Song, cInner.Collection[i].Author)
		if err := s.store.db.QueryRow(sqlRequest).Scan(&temp, &record); err != nil {
			if err == sql.ErrNoRows {
				return store.ErrRecordNotFound
			}
			return err
		}
		record_id = append(record_id, record)
		if i == 0 {
			count = temp
		} else if temp < count {
			count = temp
		}
	}
	var value int
	rows, err := s.store.db.Query(`select * from collections_value Order By count_max`)
	if err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		var max_count string
		if err := rows.Scan(&id, &name, &max_count); err != nil {
			return err
		}
		max, err := strconv.Atoi(max_count)
		if err != nil {
			return err
		}
		if count <= max {
			value, err = strconv.Atoi(id)
			if err != nil {
				return err
			}
			break
		}
	}
	sqlStatment := `INSERT INTO collections(title, id_value) VALUE (?, ?);`
	res, err := s.store.db.Exec(sqlStatment, title, value)
	if err != nil {
		return err
	}
	id_col, err := res.LastInsertId()
	if err != nil {
		return err
	}
	for i := 0; i < len(record_id); i++ {
		sqlStatment = `INSERT INTO collection_details(id_collection, id_rec) VALUE (?, ?);`
		_, err := s.store.db.Exec(sqlStatment, id_col, record_id[i])
		if err != nil {
			return err
		}
	}
	sqlStatment = `INSERT INTO collections_in_salon(id_collection, id_salon, count_, price, sale) VALUE (?, ?, ?, ?,?);`
	_, err = s.store.db.Exec(sqlStatment, id_col, sel.ID_salon, count, price, sale)
	if err != nil {
		return err
	}
	return nil
}

func (s *SellerRepository) GetAllCompositions(sel *seller.Seller) (*model.MusicInner, error) {
	mas := []model.MusicCard{}
	mInner := &model.MusicInner{}

	sqlRequest := fmt.Sprintf(`Select musician.mus_name, title from records_in_salon, records, musician 
	where records_in_salon.id_rec = records.id_rec  and records.id_mus = musician.id_mus and id_salon="%d";`, sel.ID_salon)
	rows, err := s.store.db.Query(sqlRequest)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		musCard := &model.MusicCard{}
		if err := rows.Scan(&musCard.Author, &musCard.Song); err != nil {
			return nil, err
		}
		musCard.Salon = sel.Salon_name
		mas = append(mas, *musCard)
	}
	mInner.Music = mas
	return mInner, nil
}

func (s *SellerRepository) GetAllCollections(sel *seller.Seller) (*seller.Collection, error) {
	col := &seller.Collection{}
	cInnerMas := []seller.CollectionInner{}

	sqlRequest := fmt.Sprintf(`select id_collection, price, count_, sale, title from collections_in_salon
	left join collections using (id_collection)
	where id_salon = '%d';`, sel.ID_salon)
	rows, err := s.store.db.Query(sqlRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection_id, sale int
		colInner := &seller.CollectionInner{}
		if err := rows.Scan(&collection_id, &colInner.Price, &colInner.Count, &sale, &colInner.Title); err != nil {
			return nil, err
		}
		colInner.Owner = sel.Salon_name
		colInner.Price = colInner.Price - colInner.Price*sale/100
		colInnerMas := []seller.CollectionCard{}
		sqlRequest := fmt.Sprintf(`select title, mus_name, id_rec from collection_details
			left join records using(id_rec)
			left join musician using(id_mus) 
			where id_collection = '%d';`, collection_id)
		rows, err := s.store.db.Query(sqlRequest)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var record int
			colCard := &seller.CollectionCard{}
			if err := rows.Scan(&colCard.Title, &colCard.Artist, &record); err != nil {
				return nil, err
			}
			sqlRequest = fmt.Sprintf(`Select count_, price, sale from records_in_salon
			where id_salon="%d" and id_rec="%d";`, sel.ID_salon, record)
			if err := s.store.db.QueryRow(sqlRequest).Scan(&colCard.Count, &colCard.Price, &colCard.Sale); err != nil {
				if err == sql.ErrNoRows {
					return nil, store.ErrRecordNotFound
				}
				return nil, err
			}
			colInnerMas = append(colInnerMas, *colCard)
		}
		colInner.Collection = colInnerMas
		cInnerMas = append(cInnerMas, *colInner)
	}
	col.Collection = cInnerMas
	return col, nil
}

func (s *SellerRepository) GetCountOfListenings(author string, song string) (int, error) {
	var count int
	sqlRequest := fmt.Sprintf(`Select bought from records, musician where records.id_mus = musician.id_mus and mus_name = '%s' and title = '%s';`, author, song)
	if err := s.store.db.QueryRow(sqlRequest).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return 0, store.ErrRecordNotFound
		}
		return 0, err
	}
	return count, nil
}

func (s *SellerRepository) EnterDiscount(song string, id_s int, sale int) error {
	sqlStatment := fmt.Sprintf(`update records_in_salon set sale = '%d' where id_rec = (select id_rec from records where title = '%s') and id_salon = '%d'`, sale, song, id_s)
	_, err := s.store.db.Exec(sqlStatment)
	if err != nil {
		return err
	}
	return nil
}

func (s *SellerRepository) GetCurrentSale(song string, id_s int) (int, error) {
	var count int
	sqlStatment := fmt.Sprintf(`select sale from records_in_salon where id_rec = (select id_rec from records where title = '%s') and id_salon = '%d';`, song, id_s)
	if err := s.store.db.QueryRow(sqlStatment).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return 0, store.ErrRecordNotFound
		}
		return 0, err
	}
	return count, nil
}

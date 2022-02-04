package sqlstore

import (
	"database/sql"
	"fmt"
	"log"

	tools "github.com/AnnDutova/static/internal/app/model"
	"github.com/AnnDutova/static/internal/app/model/seller"
	model "github.com/AnnDutova/static/internal/app/model/user"
	"github.com/AnnDutova/static/internal/app/store"
)

type ToolsRepository struct {
	store *Store
}

func (t *ToolsRepository) GetAllCompositions() (*model.MusicInner, error) {
	mas := []model.MusicCard{}
	mInner := &model.MusicInner{}
	sqlRequest := fmt.Sprintf(`select records.title, mus_name, salon.title from records_in_salon 
	left join salon using(id_salon)
	left join records using(id_rec)
	left join musician using(id_mus)
	group by records.title;`)
	rows, err := t.store.db.Query(sqlRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		musCard := &model.MusicCard{}
		if err := rows.Scan(&musCard.Song, &musCard.Author, &musCard.Salon); err != nil {
			return nil, err
		}
		mas = append(mas, *musCard)
	}
	mInner.Music = mas
	return mInner, nil
}

func (t *ToolsRepository) GetAllCollections() (*seller.Collection, error) {
	col := &seller.Collection{}
	cInnerMas := []seller.CollectionInner{}

	sqlRequest := fmt.Sprintf(`select id_collection, price, count_, sale, collections.title, salon.title from collections_in_salon
	left join collections using (id_collection)
	left join salon using(id_salon)
	group by id_collection;`)
	rows, err := t.store.db.Query(sqlRequest)
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
		if err := rows.Scan(&collection_id, &colInner.Price, &colInner.Count, &sale, &colInner.Title, &colInner.Owner); err != nil {
			return nil, err
		}
		colInner.Price = colInner.Price - colInner.Price*sale/100
		colInnerMas := []seller.CollectionCard{}
		sqlRequest := fmt.Sprintf(`select title, mus_name from collection_details
			left join records using(id_rec)
			left join musician using(id_mus) 
			where id_collection = '%d';`, collection_id)
		rows, err := t.store.db.Query(sqlRequest)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			colCard := &seller.CollectionCard{}
			if err := rows.Scan(&colCard.Title, &colCard.Artist); err != nil {
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
func (t *ToolsRepository) GetAllValuesCollections() (*tools.CollectionValueInner, error) {
	mas := []tools.CollectionValueCard{}
	cvInner := &tools.CollectionValueInner{}
	sqlRequest := fmt.Sprintf(`select title, count_max from collections_value;`)
	rows, err := t.store.db.Query(sqlRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		cvCard := &tools.CollectionValueCard{}
		if err := rows.Scan(&cvCard.ValueTitle, &cvCard.Count); err != nil {
			return nil, err
		}
		mas = append(mas, *cvCard)
	}
	cvInner.ValueCards = mas
	return cvInner, nil
}

func (t *ToolsRepository) GetAllGenres() ([]string, error) {
	var mas []string
	rows, err := t.store.db.Query(`select name from genre`)
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
func (t *ToolsRepository) GetAllArtists() ([]string, error) {
	var mas []string
	rows, err := t.store.db.Query(`select mus_name from musician`)
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

func (t *ToolsRepository) GetAllMaxAndMinPrice() (int, int, error) {
	var max, min int
	sqlRequest := fmt.Sprintf(`call find_max_min_price()`)
	if err := t.store.db.QueryRow(sqlRequest).Scan(&max, &min); err != nil {
		if err == sql.ErrNoRows {
			return 0, 0, store.ErrRecordNotFound
		}
		return 0, 0, err
	}
	return max, min, nil
}

func (t *ToolsRepository) GetCountOfListenings(author string, song string) (int, error) {
	var count int
	sqlRequest := fmt.Sprintf(`Select bought from records, musician where records.id_mus = musician.id_mus and mus_name = '%s' and title = '%s';`, author, song)
	if err := t.store.db.QueryRow(sqlRequest).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return 0, store.ErrRecordNotFound
		}
		return 0, err
	}
	return count, nil
}

func (t *ToolsRepository) GetRecomendation(preferences []string) (*seller.Collection, error) {
	var collections_id, genre_id []int
	for i := 0; i < len(preferences); i++ {
		sqlRequest := fmt.Sprintf(`Select id_genre from genre where name = '%s';`, preferences[i])
		var genre int
		if err := t.store.db.QueryRow(sqlRequest).Scan(&genre); err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}
		genre_id = append(genre_id, genre)
	}
	log.Print("revert prefe{strings} to pref{id_genre}")
	rows, err := t.store.db.Query(`call return_main_collection_genre()`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var col_id, genre int
		if err := rows.Scan(&col_id, &genre); err != nil {
			return nil, err
		}
		if exist(genre_id, genre) {
			collections_id = append(collections_id, col_id)
		}
	}
	log.Print("Create collections id mas")
	collection := &seller.Collection{}
	if len(collections_id) > 0 {
		for i := 0; i < len(collections_id); i++ {
			sqlRequest := fmt.Sprintf(`select salon.id_salon, salon.title, collections.title, count_, sale, price from collections_in_salon
			left join salon using(id_salon), collections where 
			collections.id_collection = collections_in_salon.id_collection and collections_in_salon.id_collection="%d";`, collections_id[i])
			rows, err := t.store.db.Query(sqlRequest)
			if err != nil {
				if err == sql.ErrNoRows {
					return nil, store.ErrRecordNotFound
				}
				return nil, err
			}
			defer rows.Close()
			log.Print("Find id-salon, salon.title, count_, sale, price of Collection ")
			for rows.Next() {
				mas := []seller.CollectionCard{}
				cInner := &seller.CollectionInner{}
				var sale, salon int
				if err := rows.Scan(&salon, &cInner.Owner, &cInner.Title, &cInner.Count, &sale, &cInner.Price); err != nil {
					return nil, err
				}
				cInner.Price = cInner.Price - cInner.Price*sale/100

				log.Print("Set new price ", cInner.Price, "Take collection ", collections_id[i], " Owner ", cInner.Owner, " Count ", cInner.Count)

				sqlRequest := fmt.Sprintf(`select records.id_rec, records.title, mus_name from collection_details, records, musician
				where  collection_details.id_rec = records.id_rec and records.id_mus = musician.id_mus and id_collection="%d";`, collections_id[i])
				rows, err = t.store.db.Query(sqlRequest)
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
					if err := rows.Scan(&record, &colCard.Title, &colCard.Artist); err != nil {
						return nil, err
					}
					log.Print("Take record ", record, " Title ", colCard.Title, " Artist ", colCard.Artist)
					sqlRequest = fmt.Sprintf(`select count_, sale, price from records_in_salon where id_rec = "%d"`, record)
					if err := t.store.db.QueryRow(sqlRequest).Scan(&colCard.Count, &colCard.Sale, &colCard.Price); err != nil {
						if err == sql.ErrNoRows {
							return nil, store.ErrRecordNotFound
						}
						return nil, err
					}
					mas = append(mas, *colCard)
				}
				cInner.Collection = mas
				collection.Collection = append(collection.Collection, *cInner)
			}
		}
	}
	return collection, nil
}

func (t *ToolsRepository) GetComposePrice(author, song, salon string) (int, error) {
	var id_rec int
	sqlRequest := fmt.Sprintf(`Select id_rec from records
	left join musician using(id_mus) where title = '%s' and mus_name = '%s';`, song, author)
	if err := t.store.db.QueryRow(sqlRequest).Scan(&id_rec); err != nil {
		if err == sql.ErrNoRows {
			return 0, store.ErrRecordNotFound
		}
		return 0, err
	}
	var price, sale int
	sqlRequest = fmt.Sprintf(`Select price, sale from records_in_salon
	left join salon using(id_salon)
	where title = '%s' and id_rec = '%d';`, salon, id_rec)
	if err := t.store.db.QueryRow(sqlRequest).Scan(&price, &sale); err != nil {
		if err == sql.ErrNoRows {
			return 0, store.ErrRecordNotFound
		}
		return 0, err
	}
	price = price - price*sale/100
	return price, nil
}
func (t *ToolsRepository) GetCurrentSale(author, song, salon string) (int, error) {
	var sale int
	sqlRequest := fmt.Sprintf(`Select sale from records_in_salon where 
		id_rec = (select id_rec from records
		left join musician using(id_mus) where title = '%s' and mus_name='%s') 
		and id_salon = (select id_salon from salon where title = '%s');`, song, author, salon)
	if err := t.store.db.QueryRow(sqlRequest).Scan(&sale); err != nil {
		if err == sql.ErrNoRows {
			return 0, store.ErrRecordNotFound
		}
		return 0, err
	}
	return sale, nil
}

func exist(mas []int, number int) bool {
	for i := 0; i < len(mas); i++ {
		if mas[i] == number {
			return true
		}
	}
	return false
}

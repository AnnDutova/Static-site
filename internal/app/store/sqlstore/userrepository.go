package sqlstore

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	model "github.com/AnnDutova/static/internal/app/model/user"
	"github.com/AnnDutova/static/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	sqlStatment := `INSERT INTO users(username, email, encrypted_password) VALUE (?, ?, ?);`
	res, err := r.store.db.Exec(sqlStatment, u.Username, u.Email, u.EncryptedPassword)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = int(id)
	return nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	u := &model.User{}
	sqlRequest := fmt.Sprintf(`SELECT id_user, username, email, encrypted_password FROM users WHERE username="%s"`, username)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&u.ID, &u.Username, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) AddToCustomer(u *model.User) error {
	sqlRequest := fmt.Sprintf(`insert into customer (id_user, wallet) Value(%d, %d)`, u.ID, 0)
	if _, err := r.store.db.Exec(sqlRequest); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	return nil
}

func (r *UserRepository) GetCurrentAccount(u *model.User) error {
	sqlRequest := fmt.Sprintf(`select wallet from customer, users where customer.id_user = users.id_user and users.username = '%s';`, u.Username)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&u.Money); err != nil {
		if err == sql.ErrNoRows {
			u.Money = 0
			return nil
		}
		log.Print(err)
		return err
	}
	return nil
}

func (r *UserRepository) GetBuyTransaction(u *model.User, sum int) error {
	sqlRequest := fmt.Sprintf(`call buy_transaction(%d, %d)`, sum, u.ID)
	if _, err := r.store.db.Exec(sqlRequest); err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) GetTransaction(u *model.User, sum int) error {
	sqlRequest := fmt.Sprintf(`call transaction(%d, %d)`, sum, u.ID)
	if _, err := r.store.db.Exec(sqlRequest); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) AvailableProfile(u *model.User) (bool, error) {
	var customer int
	var res int
	sqlRequest := fmt.Sprintf(`select id_customer from customer where id_user = '%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&customer); err != nil {
		if err == sql.ErrNoRows {
			return false, store.ErrRecordNotFound
		}
		return false, err
	}
	sqlRequest = fmt.Sprintf(`select id_status from status_of_customer where id_customer='%d';`, customer)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&res); err != nil {
		if err == sql.ErrNoRows {
			return false, store.ErrRecordNotFound
		}
		return false, err
	}
	if res == 1 {
		return true, nil
	}
	return false, nil
}

func (r *UserRepository) CreateCustomerStatus(u *model.User) error {
	var customer int
	sqlRequest := fmt.Sprintf(`select id_customer from customer where id_user = '%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&customer); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	sqlRequest = fmt.Sprintf(`Insert into status_of_customer(id_customer, start_time, end_time, id_status) Value(%d, current_date(), DATE_ADD(current_date(),INTERVAL 365 DAY), 1 );`, customer)
	if _, err := r.store.db.Exec(sqlRequest); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) WhoAmI(u *model.User) (*model.IsAuthorized, error) {
	isAuth := &model.IsAuthorized{}
	person := &model.WhoIs{}
	var id int
	sqlRequest := fmt.Sprintf(`select id_customer from customer where id_user='%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			var seller int
			sqlRequest := fmt.Sprintf(`select id_seller from seller where id_user='%d';`, u.ID)
			if err := r.store.db.QueryRow(sqlRequest).Scan(&seller); err != nil {
				if err == sql.ErrNoRows {
					var administrator int
					sqlRequest := fmt.Sprintf(`select id_administrator from administrator where id_user='%d';`, u.ID)
					if err := r.store.db.QueryRow(sqlRequest).Scan(&administrator); err != nil {
						if err == sql.ErrNoRows {
							return nil, store.ErrRecordNotFound
						}
						return nil, err
					}
					person.ID = administrator
					person.IsAuth = true
					isAuth.IsAdministrator = person
					isAuth.IsCustomer = nil
					isAuth.IsSeller = nil

				}
				return nil, err
			}
			person.ID = seller
			person.IsAuth = true
			isAuth.IsSeller = person
			isAuth.IsAdministrator = nil
			isAuth.IsCustomer = nil
			return isAuth, nil
		}
		return nil, err
	}
	person.ID = id
	person.IsAuth = true
	isAuth.IsCustomer = person
	isAuth.IsAdministrator = nil
	isAuth.IsSeller = nil
	return isAuth, nil
}

func (r *UserRepository) FindBucketCondition(u *model.User) (*model.Bucket, int, error) {
	var total int
	bct := &model.Bucket{}
	mas := []model.BucketElem{}
	var customer int
	sqlRequest := fmt.Sprintf(`select id_customer from customer where id_user = '%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&customer); err != nil {
		if err == sql.ErrNoRows {
			return nil, 0, store.ErrRecordNotFound
		}
		return nil, 0, err
	}

	log.Print("Find id_customer ", customer)
	sqlRequest = fmt.Sprintf(`call bucket_music_user_page(%d)`, customer)
	rows, err := r.store.db.Query(sqlRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	defer rows.Close()
	for rows.Next() {
		el := &model.BucketElem{}
		if err := rows.Scan(&el.Salon, &el.Author, &el.Song, &el.Count, &el.Price); err != nil {
			return nil, 0, err
		}
		total += el.Price
		mas = append(mas, *el)
	}

	sqlRequest = fmt.Sprintf(`select id_salon, collections.id_collection, collections.title, collections_value.title, Sum(count_) as count from order_detail_collections 
	left join order_ using(id_order) left join order_info using(id_order), collections, collections_value
	where order_detail_collections.id_collection = collections.id_collection and collections_value.id_value = collections.id_value and
	id_status = '1' and id_customer = '%d'
	group by collections.title, id_salon;`, customer)
	rows, err = r.store.db.Query(sqlRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	defer rows.Close()
	for rows.Next() {
		log.Print("Result of loop")
		el := &model.BucketElem{}
		var id_salon, id_collection int
		if err := rows.Scan(&id_salon, &id_collection, &el.Song, &el.Author, &el.Count); err != nil {
			return nil, 0, err
		}
		log.Print(id_collection, " Salon ", id_salon, " Song ", el.Song, " Author ", el.Author, " Count ", el.Count)
		sqlRequest = fmt.Sprintf(`select title from salon where id_salon = '%d';`, id_salon)
		if err := r.store.db.QueryRow(sqlRequest).Scan(&el.Salon); err != nil {
			if err == sql.ErrNoRows {
				return nil, 0, store.ErrRecordNotFound
			}
			return nil, 0, err
		}
		log.Print("Salon title ", el.Salon)
		var sale int
		sqlRequest = fmt.Sprintf(`select price, sale from collections_in_salon where id_salon = '%d' and id_collection='%d'`, id_salon, id_collection)
		if err := r.store.db.QueryRow(sqlRequest).Scan(&el.Price, &sale); err != nil {
			if err == sql.ErrNoRows {
				return nil, 0, store.ErrRecordNotFound
			}
			return nil, 0, err
		}
		el.Price = el.Price - el.Price*sale/100
		log.Print("Final price for customer ", el.Price)
		total += el.Price
		mas = append(mas, *el)
	}
	bct.Bucket = mas
	log.Print("Loop end")
	return bct, total, nil
}

func (r *UserRepository) FindMusicContainer(u *model.User) (*model.MusicInner, error) {
	var customer int
	sqlRequest := fmt.Sprintf(`select id_customer from customer where id_user = '%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&customer); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	mas := []model.MusicCard{}
	mInner := &model.MusicInner{}
	sqlRequest = fmt.Sprintf(`call music_inner_user_page(%d)`, customer)
	rows, err := r.store.db.Query(sqlRequest)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		musCard := &model.MusicCard{}
		if err := rows.Scan(&musCard.Salon, &musCard.Author, &musCard.Song); err != nil {
			return nil, err
		}
		mas = append(mas, *musCard)
	}
	mInner.Music = mas

	if len(mInner.Music) == 0 {
		return nil, nil
	} else {
		return mInner, nil
	}
}

func (r *UserRepository) FindCollectionContainer(u *model.User) (*model.CollectionInner, error) {
	var customer int
	sqlRequest := fmt.Sprintf(`select id_customer from customer where id_user = '%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&customer); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	mas := []model.CollectionCard{}
	cInner := &model.CollectionInner{}
	sqlRequest = fmt.Sprintf(`call collections_inner_user_page(%d)`, customer)
	rows, err := r.store.db.Query(sqlRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		colCard := &model.CollectionCard{}
		if err := rows.Scan(&colCard.Salon, &colCard.Author, &colCard.Song); err != nil {
			return nil, err
		}
		mas = append(mas, *colCard)
	}
	cInner.Collection = mas

	if len(cInner.Collection) == 0 {
		return nil, nil
	} else {
		return cInner, nil
	}
}

func (r *UserRepository) AddToBucket(u *model.User, card *model.MusicCard) error {
	var customer int
	sqlRequest := fmt.Sprintf(`select id_customer from customer where id_user = '%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&customer); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	order := -1
	sqlRequest = fmt.Sprintf(`Select id_order from order_ left join order_info using(id_order) where id_status=1 and id_customer='%d';`, customer)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&order); err != nil {
		if err == sql.ErrNoRows {
			sqlStatment := fmt.Sprintf(`insert into order_(id_customer) Value('%d')`, customer)
			res, err := r.store.db.Exec(sqlStatment)
			if err != nil {
				return err
			}
			id, err := res.LastInsertId()
			if err != nil {
				return err
			}
			order = int(id)
			sqlStatment = fmt.Sprintf(`insert into order_info(id_order, id_status, start_day) Value('%d',1,current_date())`, order)
			res, err = r.store.db.Exec(sqlStatment)
			log.Print(err)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	log.Print(order)
	var record int
	sqlRequest = fmt.Sprintf(`Select id_rec from records, musician where records.id_mus = musician.id_mus and mus_name = '%s' and title = '%s';`, card.Author, card.Song)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&record); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	var salon int
	sqlRequest = fmt.Sprintf(`Select id_salon from salon where title="%s";`, card.Salon)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&salon); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	sqlRequest = fmt.Sprintf(`Insert into order_detail_records(id_order, id_rec, id_salon, count_) Value(%d,%d,%d,%d)`, order, record, salon, 1)
	_, err := r.store.db.Exec(sqlRequest)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) AddCollectionToBucket(u *model.User, card *model.CollectionCard) error {
	var cus_id int
	sqlRequest := fmt.Sprintf(`select id_customer from customer where id_user = '%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&cus_id); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	order := -1
	sqlRequest = fmt.Sprintf(`Select id_order from order_ left join order_info using(id_order) where id_status=1 and id_customer='%d';`, cus_id)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&order); err != nil {
		if err == sql.ErrNoRows {
			log.Print("Meet with no Rows")
			sqlStatment := fmt.Sprintf(`insert into order_(id_customer) Value('%d')`, cus_id)
			res, err := r.store.db.Exec(sqlStatment)
			if err != nil {
				return err
			}
			id, err := res.LastInsertId()
			if err != nil {
				return err
			}
			order = int(id)
			log.Print("Order ", order)
			sqlStatment = fmt.Sprintf(`insert into order_info(id_order, id_status, start_day) Value('%d', 1, current_date())`, order)
			res, err = r.store.db.Exec(sqlStatment)
			if err != nil {
				return err
			}
		} else {
			log.Print(err)
			return err
		}
	}
	log.Print(order)
	var collection, salon int
	sqlRequest = fmt.Sprintf(`Select id_collection, id_salon from collections_in_salon
	left join salon using(id_salon)
	left join collections using(id_collection)
	where salon.title='%s' and collections.title="%s"`, card.Salon, card.Song)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&collection, &salon); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	log.Print(collection, salon)
	sqlRequest = fmt.Sprintf(`Insert into order_detail_collections(id_order, id_collection, id_salon, count_) Value(%d,%d,%d,%d)`, order, collection, salon, 1)
	_, err := r.store.db.Exec(sqlRequest)
	if err != nil {
		return err
	}
	log.Print("After insert")
	return nil
}

func (r *UserRepository) DeliteFromBucket(u *model.User, card *model.MusicCard) error {
	var customer int
	sqlRequest := fmt.Sprintf(`Select id_customer from customer where id_user='%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&customer); err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}
	log.Print(customer)
	order := -1
	sqlRequest = fmt.Sprintf(`Select id_order from order_ left join order_info using(id_order) where id_status=1 and id_customer='%d';`, customer)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&order); err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}
	log.Print(order)
	var record, id_salon int
	sqlRequest = fmt.Sprintf(`Select id_rec, salon.id_salon from records_in_salon
	left join records using(id_rec), salon 
	where salon.id_salon = records_in_salon.id_salon and salon.title='%s' and records.title="%s";`, card.Salon, card.Song)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&record, &id_salon); err != nil {
		if err == sql.ErrNoRows {
			var collection int
			sqlRequest = fmt.Sprintf(`Select id_collection, salon.id_salon from collections_in_salon
			left join collections using(id_collection), salon 
			where salon.id_salon = collections_in_salon.id_salon and salon.title='%s' and collections.title="%s";`, card.Salon, card.Song)
			if err := r.store.db.QueryRow(sqlRequest).Scan(&collection, &id_salon); err != nil {
				return err
			}
			log.Print("Coll ", collection)
			sqlRequest = fmt.Sprintf(`delete from order_detail_collections where id_collection='%d' and id_order='%d' and id_salon ='%d';`, collection, order, id_salon)
			_, err := r.store.db.Exec(sqlRequest)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	log.Print(record)
	sqlRequest = fmt.Sprintf(`delete from order_detail_records where id_rec='%d' and id_order='%d' and id_salon ='%d';`, record, order, id_salon)
	_, err := r.store.db.Exec(sqlRequest)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) BlockUser(u *model.User) error {
	sqlRequest := fmt.Sprintf(`update status_of_customer set id_status=3 where id_customer = (
	select id_customer from customer where id_user='%d') limit 1;`, u.ID)
	_, err := r.store.db.Exec(sqlRequest)
	log.Print(err)
	if err != nil {
		return err
	}
	sqlRequest = fmt.Sprintf(`update status_of_customer set start_time = current_date() where id_customer = (
	select id_customer from customer where id_user='%d') limit 1;`, u.ID)
	_, err = r.store.db.Exec(sqlRequest)
	log.Print(err)
	if err != nil {
		return err
	}
	sqlRequest = fmt.Sprintf(`update status_of_customer set end_time = DATE_ADD(current_date(),INTERVAL 31 DAY) where id_customer = (
	select id_customer from customer where id_user='%d') limit 1;`, u.ID)
	_, err = r.store.db.Exec(sqlRequest)
	log.Print(err)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) BuyAllBucket(u *model.User) error {
	sqlRequest := fmt.Sprintf(`Select id_order from order_
		left join customer using(id_customer)
		left join order_info using(id_order)
		where id_user = '%d'and id_status ='1';`, u.ID)
	rows, err := r.store.db.Query(sqlRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var order_id int
		if err := rows.Scan(&order_id); err != nil {
			return err
		}
		sqlRequest = fmt.Sprintf(`Update order_info set id_status = '3' where id_order="%d"`, order_id)
		_, err = r.store.db.Exec(sqlRequest)
		log.Print(err)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *UserRepository) GetPreferences(u *model.User) ([]string, error) {
	mas := []string{}
	var customer int
	sqlRequest := fmt.Sprintf(`Select id_customer from customer where id_user='%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&customer); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	log.Print(customer)
	sqlRequest = fmt.Sprintf(`select name from preferences, genre where preferences.id_genre = genre.id_genre and id_customer='%d';`, customer)
	rows, err := r.store.db.Query(sqlRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		mas = append(mas, name)
	}
	return mas, nil
}

func (r *UserRepository) AddPreferences(u *model.User, pref string) error {
	var customer int
	sqlRequest := fmt.Sprintf(`Select id_customer from customer where id_user='%d';`, u.ID)
	if err := r.store.db.QueryRow(sqlRequest).Scan(&customer); err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}
	mas := strings.Split(pref, ",")
	for i := 0; i < len(mas)-1; i++ {
		log.Print(mas[i])
		sqlRequest = fmt.Sprintf(`Insert into preferences(id_genre, id_customer) Value(%s,%d);`, mas[i], customer)
		if err := r.store.db.QueryRow(sqlRequest).Scan(&customer); err != nil {
			if err == sql.ErrNoRows {
				log.Print(err)
				return err
			}
			log.Print(err)
			return err
		}
	}
	return nil
}

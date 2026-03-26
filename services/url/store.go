package url

import (
	"database/sql"
	"fmt"

	"github.com/AnirudhV16/UShort/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

/*func (s *Store) Checkurl(shorturl string) (bool, error) {
	//check if url exists in the db
	_, err := s.db.Exec("select * from urls where shorturl == ?", shorturl)
	if err != nil {
		//no shorturl present in db
		return false, nil
	}
	return true, fmt.Errorf("url already exists...")
}*/

func (s *Store) GetUrlByShorturl(shorturl string) (string, error) {
	//details associated to url if url exists in the db
	row := s.db.QueryRow(
		"SELECT id, short_url, original_url, created_at FROM urls WHERE short_url = ?",
		shorturl,
	)
	u := new(types.URL)
	err := row.Scan(&u.ID, &u.Short_url, &u.Original_url, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}
	/*if u.ID == 0 {
		return "", fmt.Errorf("user not found")
	}*/
	return u.Original_url, nil
}

func (s *Store) AddUrl(url types.URL) error {
	//here to add a new url we need the user id, we get that from context, if user not logged in? then??
	//not compulsory of having a short url associated to user?
	_, err := s.db.Exec("INSERT INTO urls (Short_url, Original_url) VALUES (?, ?)", url.Short_url, url.Original_url)
	if err != nil {
		return nil
	}
	return fmt.Errorf("duplicate shorturl found....")
}

/*func scanRowIntoUrl(rows *sql.Rows) (*types.URL, error) {
	url := new(types.URL)

	err := rows.Scan(
		&url.ID,
		&url.Short_url,
		&url.Original_url,
		&url.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return url, nil
}*/

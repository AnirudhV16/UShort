package types

import "time"

//our stores type
type Mystore map[string]string

type RegisterUserPayload struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        int       `json:"ID"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserStore interface {
	GetUserByGmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type URLPayload struct {
	Original_url string `json:"Original_url"`
}

type URL struct {
	ID           int       `json:"ID"`
	Short_url    string    `json:"Short_url"`
	Original_url string    `json:"Original_url"`
	CreatedAt    time.Time `json:"CreatedAt"`
}

type UrlStore interface {
	GetUrlById(int) (string, error)
	AddUrl(int, string) error
}

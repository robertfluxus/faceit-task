package user

type User struct {
	ID        string
	FirstName string
	LastName  string
	Nickname  string
	Password  string
	Email     string
	Country   string
}

const (
	FieldUnknown  = "unknown"
	FieldCountry  = "country"
	FieldNickname = "nickname"
)

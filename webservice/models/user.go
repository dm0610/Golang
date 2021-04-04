package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID int32 = 1
)

func GetUsers() []*User {
	return users
}

func AddUsers(u User) (User, error) {
	u.ID = int(nextID)
	nextID++
	users = append(users, &u)
	return u, nil
}

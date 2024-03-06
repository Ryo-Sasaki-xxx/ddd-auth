package domains

type User struct {
	email UserEmail
	Id    UserId
	name  UserName
}

func NewUser(email UserEmail, id UserId, name UserName) (*User, error) {
	user := &User{}

	user.ChangeEmail(email)
	user.Id = id
	user.ChangeName(name)

	return user, nil
}

func (u *User) ChangeEmail(email UserEmail) {
	u.email = email
}

func (u *User) ChangeName(name UserName) {
	u.name = name
}

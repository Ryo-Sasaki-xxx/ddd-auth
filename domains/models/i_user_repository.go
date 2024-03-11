package domains

type IUserRepository interface {
	Find(UserId) bool
	Save(User) error
}

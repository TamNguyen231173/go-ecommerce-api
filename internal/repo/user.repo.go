package repo

//type UserRepo struct{}
//
//func NewUserRepo() *UserRepo {
//	return &UserRepo{}
//}
//
//func (ur *UserRepo) GetUserById(uid string) string {
//	return "this is user repo"
//}

type IUserRepo interface {
	GetUserByEmail(email string) string
}

type UserRepo struct{}

func (u UserRepo) GetUserByEmail(email string) string {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository() IUserRepo {
	return &UserRepo{}
}

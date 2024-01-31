package domain

const CollectionUser = "users"

type User struct{}

type UserUpdate struct{}

type UserRepository interface {
	Create(*User) (*User, error)
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)

	Update(*UserUpdate) error
}

type UserUsecase interface {
	Create(*User) (*User, error)
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)

	Update(*UserUpdate) error
}

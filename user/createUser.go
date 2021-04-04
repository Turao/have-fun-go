package user

import "log"

type CreateUser interface {
	Execute(name string) (User, error)
}

type CreateUserUseCase struct {
	Repository Repository
}

func (uc CreateUserUseCase) Execute(name string) (User, error) {
	log.Println("[create user use-case]", "creating user...")
	user := New(name)

	created, err := uc.Repository.CreateUser(user)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}
	return created, nil
}

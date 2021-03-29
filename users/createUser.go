package users

import "log"

type CreateUser interface {
	Execute(name string) (*user, error)
}

type CreateUserUseCase struct {
	Repository Repository
}

func (uc *CreateUserUseCase) Execute(name string) (*user, error) {
	log.Println("[createUser] Creating User...")
	user := New(name)

	created, err := uc.Repository.CreateUser(user)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}
	return created, nil
}

package user

import "log"

type ListUsers interface {
	Execute() ([]*user, error)
}

type ListUsersUseCase struct {
	Repository Repository
}

func (uc ListUsersUseCase) Execute() ([]*user, error) {
	log.Println("[listUser]", "Listing User...")

	users, err := uc.Repository.GetUsers()
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}
	return users, nil
}

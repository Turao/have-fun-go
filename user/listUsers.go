package user

import "log"

type ListUsers interface {
	Execute() ([]User, error)
}

type ListUsersUseCase struct {
	Repository Repository
}

func (uc ListUsersUseCase) Execute() ([]User, error) {
	log.Println("[list user use-case]", "listing users...")

	users, err := uc.Repository.GetUsers()
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}
	return users, nil
}

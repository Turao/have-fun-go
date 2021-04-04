package user

import "testing"

func givenUserExists(t *testing.T, repo Repository) User {
	user, err := repo.CreateUser(New("dummy"))
	if err != nil {
		t.Fatal("failed on helper function")
	}
	return user
}

package store

import "github.com/ramizkalabayov/goproject/internal/app/model"

type UserRepository struct {
	store *Store
}
// Create ...
func (r *UserRepository) Create(u. *mode.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id"
		u.Email,
		u.EncryptedPassword
		).Scan($u.ID); err != nil {
		return nil, err
	}

	return nil, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string)(*model.User, error){
	return nil, nil
}
package repository

import (
	"database/sql"
	"user-simple-crud/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {

	return &userRepository{db: db}
}

func (r *userRepository) GetByID(id int) (*domain.User, error) {

	var user domain.User
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]domain.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Create(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	return err
}

func (r *userRepository) Update(user *domain.User) error {
	_, err := r.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
	return err
}

func (r *userRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

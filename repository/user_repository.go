package repository

import (
	"database/sql"

	user "github.com/adosalkanovicc/go_crud/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) GetUsers(limit, offset int) ([]*user.User, error) {
	rows, err := ur.DB.Query("SELECT id, name, email, age FROM users LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		var id, age int32
		var name, email string
		if err := rows.Scan(&id, &name, &email, &age); err != nil {
			return nil, err
		}

		users = append(users, &user.User{
			Id:    id,
			Name:  name,
			Email: email,
			Age:   age,
		})
	}

	return users, nil
}

func (ur *UserRepository) CreateUser(u *user.User) error {
	_, err := ur.DB.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)", u.Name, u.Email, u.Age)
	return err
}

func (ur *UserRepository) UpdateUser(u *user.User) error {
	_, err := ur.DB.Exec("UPDATE users SET name=?, email=?, age=? WHERE id=?", u.Name, u.Email, u.Age, u.Id)
	return err
}

func (ur *UserRepository) DeleteUser(id int) error {
	_, err := ur.DB.Exec("DELETE FROM users WHERE id=?", id)
	return err
}

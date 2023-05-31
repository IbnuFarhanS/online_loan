package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/online_loan/model"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Save(newUser model.User) (model.User, error) {
	query := "INSERT INTO users (username, password, nik, name, address, phone_number, user_limit, role_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	result, err := r.DB.Exec(query, newUser.Username, newUser.Password, newUser.NIK, newUser.Name, newUser.Address, newUser.PhoneNumber, newUser.Limit, newUser.RoleID, newUser.CreatedAt)
	if err != nil {
		return model.User{}, err
	}
	id, _ := result.LastInsertId()
	newUser.ID = id

	return newUser, nil
}

func (r *userRepository) Update(updatedUser model.User) (model.User, error) {
	query := "UPDATE users SET username = $1, password = $2, nik = $3, name = $4, address = $5, phone_number = $6, user_limit = $7, role_id = $8, created_at = $9 WHERE id = $10"
	_, err := r.DB.Exec(query, updatedUser.Username, updatedUser.Password, updatedUser.NIK, updatedUser.Name, updatedUser.Address, updatedUser.PhoneNumber, updatedUser.Limit, updatedUser.RoleID, updatedUser.CreatedAt, updatedUser.ID)
	if err != nil {
		return model.User{}, err
	}

	return updatedUser, nil
}

func (r *userRepository) Delete(id int64) (model.User, error) {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return model.User{}, err
	}

	return model.User{ID: id}, nil
}

func (r *userRepository) FindById(id int64) (model.User, error) {
	query := "SELECT id, username, password, nik, name, address, phone_number, user_limit, role_id, created_at FROM users WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.NIK, &user.Name, &user.Address, &user.PhoneNumber, &user.Limit, &user.RoleID, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindAll() ([]model.User, error) {
	query := "SELECT id, username, password, nik, name, address, phone_number, user_limit, role_id, created_at FROM users"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.NIK, &user.Name, &user.Address, &user.PhoneNumber, &user.Limit, &user.RoleID, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindByUsername(username string) (model.User, error) {
	query := "SELECT id, username, password, nik, name, address, phone_number, user_limit, role_id, created_at FROM users WHERE username = $1"
	row := r.DB.QueryRow(query, username)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.NIK, &user.Name, &user.Address, &user.PhoneNumber, &user.Limit, &user.RoleID, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

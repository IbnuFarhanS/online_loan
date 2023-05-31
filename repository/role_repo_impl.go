package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/online_loan/model"
)

type roleRepository struct {
	DB *sql.DB
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepository{DB: db}
}

func (r *roleRepository) Save(newRole model.Role) (model.Role, error) {
	query := "INSERT INTO roles (name) VALUES ($1)"
	result, err := r.DB.Exec(query, newRole.Name)
	if err != nil {
		return model.Role{}, err
	}
	id, _ := result.LastInsertId()
	newRole.ID = id

	return newRole, nil
}

func (r *roleRepository) Update(updatedRole model.Role) (model.Role, error) {
	query := "UPDATE roles SET name = $1 WHERE id = $2"
	_, err := r.DB.Exec(query, updatedRole.Name, updatedRole.ID)
	if err != nil {
		return model.Role{}, err
	}

	return updatedRole, nil
}

func (r *roleRepository) Delete(id int64) (model.Role, error) {
	query := "DELETE FROM roles WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return model.Role{}, err
	}

	return model.Role{ID: id}, nil
}

func (r *roleRepository) FindById(id int64) (model.Role, error) {
	query := "SELECT id, name FROM roles WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	role := model.Role{}
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return model.Role{}, err
	}

	return role, nil
}

func (r *roleRepository) FindAll() ([]model.Role, error) {
	query := "SELECT id, name FROM roles"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []model.Role{}
	for rows.Next() {
		role := model.Role{}
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) FindByName(name string) (model.Role, error) {
	query := "SELECT id, name FROM roles WHERE name = $1"
	row := r.DB.QueryRow(query, name)
	role := model.Role{}
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return model.Role{}, err
	}

	return role, nil
}

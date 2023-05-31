package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/online_loan/model"
)

type acceptStatusRepository struct {
	DB *sql.DB
}

func NewAcceptStatusRepository(db *sql.DB) AcceptStatusRepository {
	return &acceptStatusRepository{DB: db}
}

func (r *acceptStatusRepository) Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	query := "INSERT INTO accept_status (id_transaction, status, created_at) VALUES ($1, $2, $3)"
	result, err := r.DB.Exec(query, newAcceptStatus.TransactionID, newAcceptStatus.Status, newAcceptStatus.CreatedAt)
	if err != nil {
		return model.AcceptStatus{}, err
	}
	id, _ := result.LastInsertId()
	newAcceptStatus.ID = id

	return newAcceptStatus, nil
}

func (r *acceptStatusRepository) Update(updateAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	query := "UPDATE accept_status SET id_transaction = $1, status = $2, created_at = $3 WHERE id = $4"
	_, err := r.DB.Exec(query, updateAcceptStatus.TransactionID, updateAcceptStatus.Status, updateAcceptStatus.CreatedAt, updateAcceptStatus.ID)
	if err != nil {
		return model.AcceptStatus{}, err
	}

	return updateAcceptStatus, nil
}

func (r *acceptStatusRepository) Delete(id int64) (model.AcceptStatus, error) {
	query := "DELETE FROM accept_status WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return model.AcceptStatus{}, err
	}

	return model.AcceptStatus{ID: id}, nil
}

func (r *acceptStatusRepository) FindById(id int64) (model.AcceptStatus, error) {
	query := "SELECT id, id_transaction, status, created_at FROM accept_status WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	acceptStatus := model.AcceptStatus{}
	err := row.Scan(&acceptStatus.ID, &acceptStatus.TransactionID, &acceptStatus.Status, &acceptStatus.CreatedAt)
	if err != nil {
		return model.AcceptStatus{}, err
	}

	return acceptStatus, nil
}

func (r *acceptStatusRepository) FindAll() ([]model.AcceptStatus, error) {
	query := "SELECT id, id_transaction, status, created_at FROM accept_status"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	acceptStatuses := []model.AcceptStatus{}
	for rows.Next() {
		acceptStatus := model.AcceptStatus{}
		err := rows.Scan(&acceptStatus.ID, &acceptStatus.TransactionID, &acceptStatus.Status, &acceptStatus.CreatedAt)
		if err != nil {
			return nil, err
		}
		acceptStatuses = append(acceptStatuses, acceptStatus)
	}

	return acceptStatuses, nil
}

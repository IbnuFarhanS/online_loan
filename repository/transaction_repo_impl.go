package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/online_loan/model"
)

type transactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{DB: db}
}

func (r *transactionRepository) Save(newTransaction model.Transaction) (model.Transaction, error) {
	query := "INSERT INTO transactions (id_user, id_product, amount, status, created_at, due_date) VALUES ($1, $2, $3, $4, $5, $6)"
	result, err := r.DB.Exec(query, newTransaction.UserID, newTransaction.ProductID, newTransaction.Amount, newTransaction.Status, newTransaction.CreatedAt, newTransaction.DueDate)
	if err != nil {
		return model.Transaction{}, err
	}
	id, _ := result.LastInsertId()
	newTransaction.ID = id

	return newTransaction, nil
}

func (r *transactionRepository) Update(updateTransaction model.Transaction) (model.Transaction, error) {
	query := "UPDATE transactions SET id_user = $1, id_product = $2, amount = $3, status = $4, created_at = $5, due_date = $6 WHERE id = $7"
	_, err := r.DB.Exec(query, updateTransaction.UserID, updateTransaction.ProductID, updateTransaction.Amount, updateTransaction.Status, updateTransaction.CreatedAt, updateTransaction.DueDate, updateTransaction.ID)
	if err != nil {
		return model.Transaction{}, err
	}

	return updateTransaction, nil
}

func (r *transactionRepository) Delete(id int64) (model.Transaction, error) {
	query := "DELETE FROM transactions WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return model.Transaction{}, err
	}

	return model.Transaction{ID: id}, nil
}

func (r *transactionRepository) FindById(id int64) (model.Transaction, error) {
	query := "SELECT id, id_user, id_product, amount, status, created_at, due_date FROM transactions WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	transaction := model.Transaction{}
	err := row.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Amount, &transaction.Status, &transaction.CreatedAt, &transaction.DueDate)
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

func (r *transactionRepository) FindAll() ([]model.Transaction, error) {
	query := "SELECT id, id_user, id_product, amount, status, created_at, due_date FROM transactions"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []model.Transaction{}
	for rows.Next() {
		transaction := model.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Amount, &transaction.Status, &transaction.CreatedAt, &transaction.DueDate)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *transactionRepository) UpdateStatus(transactionID uint, status bool) error {
	query := "UPDATE transactions SET status = $1 WHERE id = $2"
	_, err := r.DB.Exec(query, status, transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) FindByUserID(userID int64) ([]model.Transaction, error) {
	query := "SELECT id, id_user, id_product, amount, status, created_at, due_date FROM transactions WHERE id_user = $1"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []model.Transaction{}
	for rows.Next() {
		transaction := model.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Amount, &transaction.Status, &transaction.CreatedAt, &transaction.DueDate)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

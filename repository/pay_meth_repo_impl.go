package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/online_loan/model"
)

type paymentMethodRepository struct {
	DB *sql.DB
}

func NewPaymentMethodRepository(db *sql.DB) PaymentMethodRepository {
	return &paymentMethodRepository{DB: db}
}

func (r *paymentMethodRepository) Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	query := "INSERT INTO payment_methods (name, created_at) VALUES ($1, $2)"
	result, err := r.DB.Exec(query, newPaymentMethod.Name, newPaymentMethod.CreatedAt)
	if err != nil {
		return model.PaymentMethod{}, err
	}
	id, _ := result.LastInsertId()
	newPaymentMethod.ID = id

	return newPaymentMethod, nil
}

func (r *paymentMethodRepository) Update(updatePaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	query := "UPDATE payment_methods SET name = $1, created_at = $2 WHERE id = $3"
	_, err := r.DB.Exec(query, updatePaymentMethod.Name, updatePaymentMethod.CreatedAt, updatePaymentMethod.ID)
	if err != nil {
		return model.PaymentMethod{}, err
	}

	return updatePaymentMethod, nil
}

func (r *paymentMethodRepository) Delete(id int64) (model.PaymentMethod, error) {
	query := "DELETE FROM payment_methods WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return model.PaymentMethod{}, err
	}

	return model.PaymentMethod{ID: id}, nil
}

func (r *paymentMethodRepository) FindById(id int64) (model.PaymentMethod, error) {
	query := "SELECT id, name, created_at FROM payment_methods WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	paymentMethod := model.PaymentMethod{}
	err := row.Scan(&paymentMethod.ID, &paymentMethod.Name, &paymentMethod.CreatedAt)
	if err != nil {
		return model.PaymentMethod{}, err
	}

	return paymentMethod, nil
}

func (r *paymentMethodRepository) FindAll() ([]model.PaymentMethod, error) {
	query := "SELECT id, name, created_at FROM payment_methods"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	paymentMethods := []model.PaymentMethod{}
	for rows.Next() {
		paymentMethod := model.PaymentMethod{}
		err := rows.Scan(&paymentMethod.ID, &paymentMethod.Name, &paymentMethod.CreatedAt)
		if err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, paymentMethod)
	}

	return paymentMethods, nil
}

func (r *paymentMethodRepository) FindByName(name string) (model.PaymentMethod, error) {
	query := "SELECT id, name, created_at FROM payment_methods WHERE name = $1"
	row := r.DB.QueryRow(query, name)
	paymentMethod := model.PaymentMethod{}
	err := row.Scan(&paymentMethod.ID, &paymentMethod.Name, &paymentMethod.CreatedAt)
	if err != nil {
		return model.PaymentMethod{}, err
	}

	return paymentMethod, nil
}

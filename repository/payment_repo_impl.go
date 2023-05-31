package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/online_loan/model"
)

type paymentRepository struct {
	DB *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{DB: db}
}

func (r *paymentRepository) Save(newPayment model.Payment) (model.Payment, error) {
	query := "INSERT INTO payments (id_transaction, payment_amount, id_payment_method, payment_date) VALUES ($1, $2, $3, $4)"
	result, err := r.DB.Exec(query, newPayment.TransactionID, newPayment.PaymentAmount, newPayment.PaymentMethodID, newPayment.PaymentDate)
	if err != nil {
		return model.Payment{}, err
	}
	id, _ := result.LastInsertId()
	newPayment.ID = id

	return newPayment, nil
}

func (r *paymentRepository) Update(updatePayment model.Payment) (model.Payment, error) {
	query := "UPDATE payments SET id_transaction = $1, payment_amount = $2, id_payment_method = $3, payment_date = $4 WHERE id = $5"
	_, err := r.DB.Exec(query, updatePayment.TransactionID, updatePayment.PaymentAmount, updatePayment.PaymentMethodID, updatePayment.PaymentDate, updatePayment.ID)
	if err != nil {
		return model.Payment{}, err
	}

	return updatePayment, nil
}

func (r *paymentRepository) Delete(id int64) (model.Payment, error) {
	query := "DELETE FROM payments WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return model.Payment{}, err
	}

	return model.Payment{ID: id}, nil
}

func (r *paymentRepository) FindById(id int64) (model.Payment, error) {
	query := "SELECT id, id_transaction, payment_amount, id_payment_method, payment_date FROM payments WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	payment := model.Payment{}
	err := row.Scan(&payment.ID, &payment.TransactionID, &payment.PaymentAmount, &payment.PaymentMethodID, &payment.PaymentDate)
	if err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}

func (r *paymentRepository) FindAll() ([]model.Payment, error) {
	query := "SELECT id, id_transaction, payment_amount, id_payment_method, payment_date FROM payments"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payments := []model.Payment{}
	for rows.Next() {
		payment := model.Payment{}
		err := rows.Scan(&payment.ID, &payment.TransactionID, &payment.PaymentAmount, &payment.PaymentMethodID, &payment.PaymentDate)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

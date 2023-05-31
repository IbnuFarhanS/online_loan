package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/online_loan/model"
)

type productRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{DB: db}
}

func (r *productRepository) Save(newProduct model.Product) (model.Product, error) {
	query := "INSERT INTO products (name, installment, interest, created_at) VALUES ($1, $2, $3, $4)"
	result, err := r.DB.Exec(query, newProduct.Name, newProduct.Installment, newProduct.Interest, newProduct.CreatedAt)
	if err != nil {
		return model.Product{}, err
	}
	id, _ := result.LastInsertId()
	newProduct.ID = id

	return newProduct, nil
}

func (r *productRepository) Update(updatedProduct model.Product) (model.Product, error) {
	query := "UPDATE products SET name = $1, installment = $2, interest = $3, created_at = $4 WHERE id = $5"
	_, err := r.DB.Exec(query, updatedProduct.Name, updatedProduct.Installment, updatedProduct.Interest, updatedProduct.CreatedAt, updatedProduct.ID)
	if err != nil {
		return model.Product{}, err
	}

	return updatedProduct, nil
}

func (r *productRepository) Delete(id int64) (model.Product, error) {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return model.Product{}, err
	}

	return model.Product{ID: id}, nil
}

func (r *productRepository) FindById(id int64) (model.Product, error) {
	query := "SELECT id, name, installment, interest, created_at FROM products WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	product := model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Installment, &product.Interest, &product.CreatedAt)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (r *productRepository) FindAll() ([]model.Product, error) {
	query := "SELECT id, name, installment, interest, created_at FROM products"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []model.Product{}
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Installment, &product.Interest, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) FindByName(name string) (model.Product, error) {
	query := "SELECT id, name, installment, interest, created_at FROM products WHERE name = $1"
	row := r.DB.QueryRow(query, name)
	product := model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Installment, &product.Interest, &product.CreatedAt)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

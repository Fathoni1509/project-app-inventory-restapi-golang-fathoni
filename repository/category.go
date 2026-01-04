package repository

import (
	"context"
	"errors"
	"project-app-inventory-restapi-golang-fathoni/database"
	"project-app-inventory-restapi-golang-fathoni/model"
)

type CategoryRepository interface {
	GetListCategories() ([]model.Category, error)
	GetListCategoryById(category_id int) (model.Category, error)
	AddCategory(category *model.Category) error
	UpdateCategory(category_id int, category *model.Category) error
	DeleteCategory(category_id int) error
}

type categoryRepository struct {
	db database.PgxIface
}

func NewCategoryRepository(db database.PgxIface) CategoryRepository {
	return &categoryRepository{db: db}
}

// get list categories
func (repo *categoryRepository) GetListCategories() ([]model.Category, error) {
	query := `SELECT category_id, name, description, created_at, updated_at, deleted_at FROM categories WHERE deleted_at IS NULL ORDER BY category_id`

	rows, err := repo.db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listCategories []model.Category
	var list model.Category
	for rows.Next() {
		err := rows.Scan(&list.CategoryId, &list.Name, &list.Description, &list.CreatedAt, &list.UpdatedAt, &list.DeletedAt)
		if err != nil {
			return nil, err
		}
		listCategories = append(listCategories, list)
	}

	return listCategories, nil
}

// get list category by ID
func (repo *categoryRepository) GetListCategoryById(category_id int) (model.Category, error) {
	query := `SELECT category_id, name, description, created_at, updated_at, deleted_at 
		FROM categories
		WHERE deleted_at IS NULL AND category_id=$1`

	var category model.Category

	err := repo.db.QueryRow(context.Background(), query, category_id).Scan(&category.CategoryId, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt, &category.DeletedAt)

	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

// add category
func (repo *categoryRepository) AddCategory(category *model.Category) error {
	query := `INSERT INTO categories (name, description, created_at, updated_at) VALUES
	($1, $2, NOW(), NOW()) RETURNING category_id`

	_, err := repo.db.Exec(context.Background(), query,
		category.Name,
		category.Description,
	)

	if err != nil {
		return err
	}

	return nil
}

// update category by ID
func (repo *categoryRepository) UpdateCategory(category_id int, category *model.Category) error {
	query := `UPDATE categories
		SET name=$1, description=$2, updated_at=NOW()
		WHERE deleted_at IS NULL AND category_id=$3`

	commandTag, err := repo.db.Exec(context.Background(), query,
		category.Name,
		category.Description,
		category_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("category not found")
	}

	return nil
}

// delete category by ID
func (repo *categoryRepository) DeleteCategory(category_id int) error {
	query := `UPDATE categories
		SET deleted_at=NOW()
		WHERE category_id=$1 AND deleted_at IS NULL`

	commandTag, err := repo.db.Exec(context.Background(), query,
		category_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
        return errors.New("category not found")
    }

	return nil
}

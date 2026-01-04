package repository

import (
	"context"
	"errors"
	"project-app-inventory-restapi-golang-fathoni/database"
	"project-app-inventory-restapi-golang-fathoni/model"
)

type UserRepository interface {
	GetListUsers() ([]model.User, error)
	GetListUserById(user_id int) (model.User, error)
	AddUser(user *model.User) error
	UpdateUser(user_id int, user *model.User) error
	DeleteUser(user_id int) error
}

type userRepository struct {
	db database.PgxIface
}

func NewUserRepository(db database.PgxIface) UserRepository {
	return &userRepository{db: db}
}

// get list users
func (repo *userRepository) GetListUsers() ([]model.User, error) {
	query := `SELECT 
		user_id, 
		username, 
		email,
		role,
		created_at, 
		updated_at, 
		deleted_at
	FROM users 
	WHERE deleted_at IS NULL 
	ORDER BY user_id`

	rows, err := repo.db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listUsers []model.User
	var list model.User
	for rows.Next() {
		err := rows.Scan(&list.UserId, &list.Username, &list.Email, &list.Role, &list.CreatedAt, &list.UpdatedAt, &list.DeletedAt)
		if err != nil {
			return nil, err
		}
		listUsers = append(listUsers, list)
	}

	return listUsers, nil
}

// get list user by ID
func (repo *userRepository) GetListUserById(user_id int) (model.User, error) {
	query := `SELECT 
		user_id, 
		username, 
		email,
		role,
		created_at, 
		updated_at, 
		deleted_at
	FROM users 
	WHERE deleted_at IS NULL AND user_id=$1
	ORDER BY user_id`

	var user model.User

	err := repo.db.QueryRow(context.Background(), query, user_id).Scan(&user.UserId, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// add user
func (repo *userRepository) AddUser(user *model.User) error {
	query := `INSERT INTO users (username, email, password, role, created_at, updated_at) VALUES
	($1, $2, $3, $4, NOW(), NOW()) RETURNING user_id`

	_, err := repo.db.Exec(context.Background(), query,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
	)

	if err != nil {
		return err
	}

	return nil
}

// update user by ID
func (repo *userRepository) UpdateUser(user_id int, user *model.User) error {
	query := `UPDATE users
		SET username=$1, email=$2, password=$3, role=$4, updated_at=NOW()
		WHERE deleted_at IS NULL AND user_id=$5`

	commandTag, err := repo.db.Exec(context.Background(), query,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
		user_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}

// delete user by ID
func (repo *userRepository) DeleteUser(user_id int) error {
	query := `UPDATE users
		SET deleted_at=NOW()
		WHERE user_id=$1 AND deleted_at IS NULL`

	commandTag, err := repo.db.Exec(context.Background(), query,
		user_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
        return errors.New("user not found")
    }

	return nil
}

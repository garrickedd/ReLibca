package repository

import (
	"errors"
	"fmt"

	"github.com/garrickedd/ReLibca/src/domain/model"
	"github.com/garrickedd/ReLibca/src/shared/config"
	"github.com/jmoiron/sqlx"
)

type RepoUserIF interface {
	CreateUser(data *model.User) (*config.Result, error)
	UpdateUser(data *model.User) (*config.Result, error)
	DeleteUser(data *model.User) (*config.Result, error)
	Get_Users(data *model.User, search string) ([]model.User, error)
	GetAuthUser(user string) (*model.User, error)
}

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

// CRUD: Create, Read, Update, Delete
// POST /v1/users (create new user)
// GET /v1/users (list user) /v1/users?page=1
// GET /v1/users/:id (get user details by id)
// (PUT || PATCH) /v1/users/:id (update user by id)
// DELETE /v1/users/:id (delete user by id)

// POST
func (r RepoUser) CreateUser(data *model.User) (*config.Result, error) {
	queryUser := `INSERT INTO users(
		first_name,
		last_name,
		username,
		email,
		pass,
		phone,
		role
		)
		VALUE(
			:first_name,
			:last_name,
			:username,
			:email,
			:pass,
			:phone,
			:role
		)`

	_, err := r.NamedExec(queryUser, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 data user Created"}, nil
}

// PUT
func (r RepoUser) UpdateUser(data *model.User) (*config.Result, error) {
	queryUser := `UPDATE public.users SET
	email = COALESCE(NULLIF(:email, ''), email),
	pass = COALESCE(NULLIF(:pass, ''), pass),
	phone = COALESCE(NULLIF(:phone, ''), phone),
	role = COALESCE(NULLIF(:role, ''), role),
	`

	_, err := r.NamedExec(queryUser, data)
	fmt.Println(r.NamedExec(queryUser, data))
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 data user Updated"}, nil
}

// DELETE
func (r RepoUser) DeleteUser(data *model.User) (*config.Result, error) {
	queryUser := `DELETE FROM public.users WHERE id_user=:id_user`

	_, err := r.NamedExec(queryUser, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 data user Deleted"}, nil
}

// READ
func (r RepoUser) Get_Users(data *model.User, search string) ([]model.User, error) {
	users_data := []model.User{}
	if search == "" {
		r.Select(&users_data, `select u.username, u.role, u.created_at FROM public.users u`)
	} else {
		r.Select(&users_data, `select u.username, u.role, u.created_at FROM public.users u`, search)
	}
	if len(users_data) == 0 {
		return nil, errors.New("data not found")
	}
	return users_data, nil
}

func (r RepoUser) GetAuthUser(user string) (*model.User, error) {
	var result model.User
	q := `SELECT id_user, username, role, pass FROM public.users WHERE username = ?`

	if err := r.Get(&result, r.Rebind(q), user); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username not found")
		}

		return nil, err
	}

	return &result, nil
}

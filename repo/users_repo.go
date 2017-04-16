package repo

import (
	"github.com/jackc/pgx"
	"reviewer/entity"
)

type UsersRepo struct {
	pool *pgx.ConnPool
}

const (
	InsertUser    = "insert into users(username, password, email) values($1, $2, $3)"
	GetByUsername = "select * from users where username=$1"
)

func NewUsersRepo(pool *pgx.ConnPool) *UsersRepo {
	return &UsersRepo{
		pool: pool,
	}
}

func (repo *UsersRepo) Create(user *entity.User) error {
	_, err := repo.pool.Exec(InsertUser, user.Username, user.Password, user.Email)
	return err
}

func (repo *UsersRepo) GetByUsername(username string) (*entity.User, error) {
	rows, err := repo.pool.Query(GetByUsername, username)
	if err != nil {
		return nil, err
	}
	var user entity.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil

}

package repo

import (
	"github.com/jackc/pgx"
	"log"
	"reviewer/entity"
)

type UsersRepo struct {
	pool *pgx.ConnPool
}

const (
	InsertUser    = "insert into users(username, password, email) values($1, $2, $3)"
	GetByUsername = "select * from users where username=$1"
)

func NewUsersRepo() *UsersRepo {
	poolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			User:     "musatov",
			Password: "",
			Database: "reviewer",
		},
		MaxConnections: 5,
	}
	pool, err := pgx.NewConnPool(poolConfig)
	if err != nil {
		log.Fatal("Unable to config connection pool")
	}
	return &UsersRepo{
		pool: pool,
	}
}

func (repo *UsersRepo) Create(user *entity.User) error {
	_, err := repo.pool.Exec(InsertUser, user.Username, user.Password, user.Email)
	return err
}

func (repo *UsersRepo) GetByUsername(username string) (user *entity.User, err error) {
	rows, err := repo.pool.Query(GetByUsername, username)
	if err != nil {
		return nil, err
	}
	var id int
	var password string
	var email string
	for rows.Next() {
		if err := rows.Scan(&id, &username, &password, &email); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	result := &entity.User{
		Id: id,
		Username: username,
		Password: password,
		Email: email,
	}
	return result, nil

}

package repo

import (
	"github.com/jackc/pgx"
	"reviewer/entity"
)

const (
	InsertReview       = "insert into reviews(title, body, user_id, is_public) values($1, $2, $3, $4) returning id"
	QueryReviews       = "select * from reviews where user_id = $1 limit 50"
	QueryPublicReviews = "select * from reviews where is_public = true limit 50"
)

type ReviewsRepo struct {
	pool *pgx.ConnPool
}

func NewReviewsRepo(pool *pgx.ConnPool) *ReviewsRepo {
	return &ReviewsRepo{
		pool: pool,
	}
}

func (repo *ReviewsRepo) CreateReview(review *entity.Review) (id int, err error) {
	row := repo.pool.QueryRow(InsertReview, review.Title, review.Body, review.User, review.Public)
	err = row.Scan(&id)
	return
}

func (repo *ReviewsRepo) DeleteReview(id int) error {
	_, err := repo.pool.Exec("delete from reviews where id = $1", id)
	return err
}

func (repo *ReviewsRepo) GetReviews(user int) (reviews []*entity.Review, err error) {
	rows, err := repo.pool.Query(QueryReviews, user)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var review entity.Review
		if err := rows.Scan(&review.Id, &review.Title, &review.Body, &review.User, &review.Public); err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}
	return
}

func (repo *ReviewsRepo) GetAllReviews() (reviews []*entity.Review, err error) {
	rows, err := repo.pool.Query(QueryPublicReviews)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var review entity.Review
		if err := rows.Scan(&review.Id, &review.Title, &review.Body, &review.User, &review.Public); err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}
	return
}

package repository

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/o-ga09/golang-svelte-app/counter"
)

type Repository struct {
	repo dynamo.Table
}

func New() *Repository {
	return &Repository{
		repo: connect(),
	}
}

func tableName() string {
	return os.Getenv("SST_Table_tableName_counter")
}

func connect() dynamo.Table {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
	return db.Table(tableName())
}

func (r *Repository) Get() (counter.Clicks, error) {
	res := &counter.Clicks{}
	err := r.repo.Get("counter", "clicks").One(res)
	return *res, err
}

func (r *Repository) Put(counter counter.Clicks) error {
	err := r.repo.Put(counter).Run()
	return err
}

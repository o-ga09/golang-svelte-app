package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/o-ga09/golang-svelte-app/notes/repository"
)

type Repository struct {
	repo dynamo.Table
}

func New(conn dynamo.Table) *Repository {
	return &Repository{
		repo: conn,
	}
}

func Connect(tableName string) dynamo.Table {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
	return db.Table(tableName)
}

func (r *Repository) Get() (interface{}, error) {
	var res interface{}
	err := r.repo.Get("counter", "clicks").One(res)
	return res, err
}

func (r *Repository) Put(param interface{}) error {
	err := r.repo.Put(param).Run()
	return err
}

func(r *Repository) GetAllNotes(userid string) ([]repository.Note, error) {
	res := []repository.Note{}
	err := r.repo.Get("userid",userid).All(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func(r *Repository) GetNote(userid string, noteid string) (repository.Note, error) {
	res := repository.Note{}
	err := r.repo.Get("userid",userid).Range("noteid",dynamo.Equal,noteid).One(res)
	if err != nil {
		return repository.Note{}, err
	}

	return res, nil
}
func(r *Repository) CreateNote(note repository.Note) error {
	err := r.repo.Put(note).Run()
	if err != nil {
		return err
	}

	return nil
}
func(r *Repository) UpdateNote(note repository.Note) error {
	err := r.repo.Put(note).Run()
	if err != nil {
		return err
	}

	return nil
}
func(r *Repository) DeleteNote(userid string, noteid string) error {
	err := r.repo.Delete("userid",userid).Range("noteid", noteid).Run()
	if err != nil {
		return err
	}

	return nil
}

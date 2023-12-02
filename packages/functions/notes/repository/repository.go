package repository

import "time"

type NoteRepository interface {
	GetAllNotes(userid string) ([]Note, error)
	GetNote(userid string, noteid string) (Note, error)
	CreateNote(Note) error
	UpdateNote(Note) error
	DeleteNote(userid string, noteid string) error
}

type Note struct {
	UserId     string `dynamo:"userId,omitempty"`
	NoteId     string `dynamo:"noteId,omitempty"`
	Contents   string `dynamo:"contents,omitempty"`
	Image_url  string `dynamo:"image_url,omitempty"`
	Created_at time.Time `dynamo:"created_at,omitempty"`
	Updated_at time.Time `dynamo:"updated_at,omitempty"`
	Deleted_at time.Time `dynamo:"deleted_at,omitempty"`
}
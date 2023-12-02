package note

import domain "github.com/o-ga09/golang-svelte-app/domain/note"

//go:generate moq -out notes_mock.go . INotes
type INotes interface {
	GetAll(id domain.UserId) (domain.Notes, error)
	GetNoteById(userId domain.UserId,noteId domain.NoteId) (domain.Note, error)
	Create(note domain.Note) error
	Update(note domain.Note) error
	Delete(userid domain.UserId, noteid domain.NoteId) error
}
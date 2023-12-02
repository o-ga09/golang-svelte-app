package notes

import (
	domain "github.com/o-ga09/golang-svelte-app/domain/note"
	"github.com/o-ga09/golang-svelte-app/notes/repository"
)

type NotesApp struct {
	Repo repository.NoteRepository
}

func New(r repository.NoteRepository) *NotesApp {
	return &NotesApp{Repo: r}
}

func(n *NotesApp) GetAll(userid domain.UserId) (domain.Notes, error) {
	res, err := n.Repo.GetAllNotes(userid.Value)
	if err != nil {
		return nil, err
	}
	notes := domain.Notes{}
	for _, note := range res {
		n := domain.Note{
			UserId: domain.UserId{Value: note.UserId},
			NoteId: domain.NoteId{Value: note.NoteId},
			Contents: domain.Contents{Value: note.Contents},
			Image_url: domain.Image_url{Value: note.Image_url},
			Created_at: domain.Created_at{Value: note.Created_at},
			Updated_at: domain.Updated_at{Value: note.Updated_at},
			Deleted_at: domain.Deleted_at{Value: note.Deleted_at},
		}
		notes = append(notes, n)
	}
	return notes, nil
}
func(n *NotesApp) GetNoteById(userid domain.UserId, noteid domain.NoteId) (domain.Note, error) {
	res, err := n.Repo.GetNote(userid.Value,noteid.Value)
	if err != nil {
		return domain.Note{}, err
	}

	note := domain.Note{
		UserId: domain.UserId{Value: res.UserId},
			NoteId: domain.NoteId{Value: res.NoteId},
			Contents: domain.Contents{Value: res.Contents},
			Image_url: domain.Image_url{Value: res.Image_url},
			Created_at: domain.Created_at{Value: res.Created_at},
			Updated_at: domain.Updated_at{Value: res.Updated_at},
			Deleted_at: domain.Deleted_at{Value: res.Deleted_at},
	}

	return note, nil
}
func(n *NotesApp) Create(note domain.Note) error {
	newNote := repository.Note{
		UserId: note.UserId.Value,
		NoteId: note.NoteId.Value,
		Contents: note.Contents.Value,
		Image_url: note.Image_url.Value,
		Created_at: note.Created_at.Value,
		Updated_at: note.Updated_at.Value,
		Deleted_at: note.Deleted_at.Value,
	}

	err := n.Repo.CreateNote(newNote)
	if err != nil {
		return err
	}

	return nil
}
func(n *NotesApp) Update(note domain.Note) error {
	newNote := repository.Note{
		UserId: note.UserId.Value,
		NoteId: note.NoteId.Value,
		Contents: note.Contents.Value,
		Image_url: note.Image_url.Value,
		Created_at: note.Created_at.Value,
		Updated_at: note.Updated_at.Value,
		Deleted_at: note.Deleted_at.Value,
	}

	err := n.Repo.UpdateNote(newNote)
	if err != nil {
		return err
	}

	return nil
}
func(n *NotesApp) Delete(userid domain.UserId, noteid domain.NoteId) error {
	err := n.Repo.DeleteNote(userid.Value, noteid.Value)
	if err != nil {
		return err
	}

	return nil
}
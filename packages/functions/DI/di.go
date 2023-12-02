package di

import (
	"os"

	"github.com/o-ga09/golang-svelte-app/counter"
	counterHandler "github.com/o-ga09/golang-svelte-app/handlers/counter"
	noteHandler "github.com/o-ga09/golang-svelte-app/handlers/notes"
	"github.com/o-ga09/golang-svelte-app/notes"
	"github.com/o-ga09/golang-svelte-app/repository"
)

func NewCounterHandler() *counterHandler.Handler {
	tableName := os.Getenv("SST_Table_tableName_counter")
	conn := repository.Connect(tableName)
	repo := repository.New(conn)
	count := counter.New(repo)
	return counterHandler.New(count)
}

func NewNoteHandler() *noteHandler.NoteHandler{
	tableName := os.Getenv("SST_Table_tableName_notes")
	conn := repository.Connect(tableName)
	repo := repository.New(conn)
	notes := notes.New(repo)
	return noteHandler.New(notes)
}

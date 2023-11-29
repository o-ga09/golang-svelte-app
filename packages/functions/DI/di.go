package di

import (
	"github.com/o-ga09/golang-svelte-app/counter"
	"github.com/o-ga09/golang-svelte-app/handlers"
	"github.com/o-ga09/golang-svelte-app/repository"
)

func New() *handlers.Handler {
	repo := repository.New()
	count := counter.New(repo)
	return handlers.New(count)
}

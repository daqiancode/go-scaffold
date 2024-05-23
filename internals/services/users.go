package services

import (
	"go-scaffold/internals/caches"
	"go-scaffold/internals/tables"
)

type Users struct{}

func (Users) Get(id uint64) (tables.User, error) {
	return caches.UserCache.Get(id)
}

func (Users) Delete(id uint64) error {
	_, err := caches.UserCache.Delete(id)
	return err
}

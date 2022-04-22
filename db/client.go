package db

import (
	"github.com/go-swagger/go-swagger/examples/accountmgr/domain"
)

type DataStore interface {
	AddAccount(account *domain.Account) (string, error)
}
type DatastoreFactory func() (DataStore, error)

var factories map[string]DatastoreFactory

func RegisterDataStore(key string, value DatastoreFactory) {
	if factories == nil {
		factories = make(map[string]DatastoreFactory)
	}
	factories[key] = value
}
func NewDataStore(dbType string) (DataStore, error) {
	return factories[dbType]()
}

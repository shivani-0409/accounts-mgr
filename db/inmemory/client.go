package inmemory

import (
	"github.com/go-swagger/go-swagger/examples/accountmgr/db"
	"github.com/go-swagger/go-swagger/examples/accountmgr/domain"
)

func init() {
	db.RegisterDataStore("inmemory", NewClient)
}

func NewClient() (db.DataStore, error) {
	return &client{ds: make(map[string]*domain.Account)}, nil
}

type client struct {
	ds map[string]*domain.Account
}

func (c *client) AddAccount(account *domain.Account) (string, error) {
	c.ds[account.ID] = account
	return account.ID, nil
}

package service

import (
	"fmt"

	"github.com/go-swagger/go-swagger/examples/accountmgr/db"
	"github.com/go-swagger/go-swagger/examples/accountmgr/domain"
)

type Manager interface {
	CreateAccount(account *domain.Account) (string, error)
}

func NewManager(dbType string) Manager {

	ds, err := db.NewDataStore(dbType)
	if err != nil {
		fmt.Println("The err is", err)
		return nil
	}
	return &manager{ds: ds}
}

package service

import (
	"time"

	"github.com/go-swagger/go-swagger/examples/accountmgr/db"
	"github.com/go-swagger/go-swagger/examples/accountmgr/domain"
	"github.com/segmentio/ksuid"
)

type manager struct {
	ds db.DataStore
}

func (m *manager) CreateAccount(input *domain.Account) (string, error) {
	input.CreatedAt = time.Now().UTC()
	input.UpdatedAt = time.Now().UTC()
	input.ID = ksuid.New().String()
	return m.ds.AddAccount(input)
}

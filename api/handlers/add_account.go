package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-swagger/go-swagger/examples/accountmgr"
	"github.com/go-swagger/go-swagger/examples/accountmgr/domain"
	"github.com/go-swagger/go-swagger/examples/accountmgr/gen/restapi/operations/account"
)

func NewAddAccount(rt *accountmgr.Runtime) account.AddAccountHandler {
	return &addAccount{rt: rt}
}

type addAccount struct {
	rt *accountmgr.Runtime
}

func (au *addAccount) Handle(aup account.AddAccountParams) middleware.Responder {
	acc := &domain.Account{Name: *aup.Body.Name, UserName: "", UserID: *aup.Body.UserID, Source: aup.Body.Source}
	if _, err := au.rt.GetManager().CreateAccount(acc); err != nil {
		return account.NewAddAccountBadRequest().WithPayload(asErrorResponse(err.(*domain.Error)))
	}

	return account.NewAddAccountCreated()
}

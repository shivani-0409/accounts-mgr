package handlers

import (
	"github.com/go-openapi/swag"
	"github.com/go-swagger/go-swagger/examples/accountmgr/domain"
	"github.com/go-swagger/go-swagger/examples/accountmgr/gen/models"
)

func asErrorResponse(e *domain.Error) *models.Error {
	er := &models.Error{
		Code:    int64(e.Code),
		Message: swag.String(e.Message),
	}
	return er
}

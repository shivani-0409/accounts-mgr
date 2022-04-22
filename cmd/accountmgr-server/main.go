package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-swagger/go-swagger/examples/accountmgr"

	"github.com/go-swagger/go-swagger/examples/accountmgr/api/handlers"
	_ "github.com/go-swagger/go-swagger/examples/accountmgr/db/inmemory"
	_ "github.com/go-swagger/go-swagger/examples/accountmgr/db/mongo"
	"github.com/go-swagger/go-swagger/examples/accountmgr/gen/restapi"
	"github.com/go-swagger/go-swagger/examples/accountmgr/gen/restapi/operations"
)

func main() {
	var portFlag = flag.Int("port", 4000, "Port to run this service on")
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	// create new service API
	api := operations.NewAccountmgrAPI(swaggerSpec)
	v := accountmgr.NewRunTime("Create Account API")
	api.AccountAddAccountHandler = handlers.NewAddAccount(v)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	flag.Parse()
	server.Port = *portFlag
	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

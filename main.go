package main

import (
	"github.com/pentabyte/tvet/api/api/endpoints"
	"github.com/pentabyte/tvet/api/api/model"
)

func main() {
	model.Migration()
	//database.StartRedisServer()
	endpoints.CreateEndpoint()
}
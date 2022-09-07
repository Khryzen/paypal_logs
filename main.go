package main

import (
	"net/http"

	"github.com/mbdeguzman/paypal_logs/api"
	"github.com/mbdeguzman/paypal_logs/models"
	"github.com/uadmin/uadmin"
)

func main() {
	RegisterModels()
	RegisterHandlers()
	DatabaseConfig()
	Server()
}

func RegisterModels() {
	uadmin.Register(
		models.PaypalLogs{},
	)
}

func RegisterHandlers() {
	http.HandleFunc("/api/", uadmin.Handler(api.APIhandler))
}

func DatabaseConfig() {
	uadmin.Database = &uadmin.DBSettings{
		Type:     "mysql",
		Host:     "localhost",
		Name:     "paypal_logs",
		User:     "root",
		Password: "Allen is Great 200%",
		Port:     3306,
	}
}

func Server() {
	uadmin.RootURL = "/admin/"
	uadmin.Port = 8080
	uadmin.StartServer()
}

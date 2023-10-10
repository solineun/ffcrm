package srvconfig

import "net/http"

func GetConfiguredSrv() *http.Server {
	return &http.Server{
		Addr: "localhost:8080",
	}
}
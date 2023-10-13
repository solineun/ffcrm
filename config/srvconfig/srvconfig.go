package srvconfig

import (
	"net/http"
	"time"
)

func GetConfiguredSrv() *http.Server {
	return &http.Server{
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,	
	}
}


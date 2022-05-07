package service

import (
	"net/http"
	"time"
)

func HttpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

package main

import (
	"net/http"
	"time"
)

func NewClient(flags Flags) *http.Client {
	tr := http.DefaultTransport.(*http.Transport)
	tr.MaxIdleConnsPerHost = flags.MaxIdleConnsPerHost
	tr.MaxIdleConns = flags.MaxIdleConns

	return &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}
}

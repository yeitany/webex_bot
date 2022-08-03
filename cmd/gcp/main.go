package main

import (
	"net/http"

	"github.com/yeitany/webex_bot/utils"

	gcp_hendler "github.com/yeitany/webex_bot/handlers/gcp"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func newGCPHandler(config *utils.Config) *gcp_hendler.GCPHandler {
	return &gcp_hendler.GCPHandler{
		Config: config,
	}
}

func main() {

	config := utils.NewConfig()

	gcpHandler := newGCPHandler(&config)

	http.HandleFunc("/webhooks/gcp", gcpHandler.ServeHttp)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		//log.Printf("method: %v\npath:/healthz", req.Method)
		w.WriteHeader(http.StatusOK)
	})

	go http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":9001", promhttp.Handler())

}

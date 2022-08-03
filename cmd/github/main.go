package main

import (
	"net/http"

	"github.com/yeitany/webex_bot/utils"

	gibhub_handler "github.com/yeitany/webex_bot/handlers/github"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func newGithubHandler(config *utils.Config) *gibhub_handler.GitHubWebhookHandler {
	return &gibhub_handler.GitHubWebhookHandler{
		Config: config,
		PRhandler: gibhub_handler.GithubPRHandler{
			Actions: []string{"opened"},
		},
		Pinghandler: gibhub_handler.PingHandler{},
	}
}

func main() {

	config := utils.NewConfig()
	githubHandler := newGithubHandler(&config)

	http.HandleFunc("/webhooks/github", githubHandler.ServeHttp)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		//log.Printf("method: %v\npath:/healthz", req.Method)
		w.WriteHeader(http.StatusOK)
	})
	go http.ListenAndServe(":9001", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

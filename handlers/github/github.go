package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yeitany/webex_bot/pkg/webex"

	"github.com/yeitany/webex_bot/utils"

	"github.com/prometheus/client_golang/prometheus"
)

type GitHubWebhookHandler struct {
	Config      *utils.Config
	PRhandler   GithubPRHandler
	Pinghandler PingHandler
}

func (g *GitHubWebhookHandler) sendMessageToWebex(text string) error {
	webexMessage := webex.WebexMessage{
		Text:    text,
		RoomId:  g.Config.RoomID,
		Token:   g.Config.Token,
		BaseUrl: g.Config.BaseUrl,
	}
	err := webexMessage.Send()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (g *GitHubWebhookHandler) ServeHttp(w http.ResponseWriter, req *http.Request) {
	githubEvent := req.Header.Get("X-GitHub-Event")
	var eventHandler GithubEvents
	switch githubEvent {
	case "pull_request":
		eventHandler = g.PRhandler
	case "ping":
		eventHandler = g.Pinghandler
	default:
		log.Printf("unexpected event %v", githubEvent)
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	bodyAsBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("error parsing request %v", err)
		utils.ReadBodyError.With(prometheus.Labels{"handler": "github"}).Inc()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	httpStatusCode := http.StatusOK
	text, err := eventHandler.handle(bodyAsBytes)
	if err != nil {
		httpStatusCode = http.StatusBadRequest
		log.Printf("error : %v", err)
	}
	if len(text) > 0 {
		err = g.sendMessageToWebex(text)
		if err != nil {
			httpStatusCode = http.StatusInternalServerError
			log.Printf("error: %v", err)
		}
	}

	w.WriteHeader(httpStatusCode)

}

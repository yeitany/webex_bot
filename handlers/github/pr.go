package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/yeitany/webex_bot/pkg/github"
	"github.com/yeitany/webex_bot/utils"

	"github.com/prometheus/client_golang/prometheus"
)

type GithubPRHandler struct {
	Actions []string
}

func (g GithubPRHandler) validateAction(eventAction string) bool {
	if len(g.Actions) == 0 {
		return true
	}

	for _, action := range g.Actions {
		if action == eventAction {
			return true
		}
	}
	return false
}

func (g GithubPRHandler) handle(bodyAsBytes []byte) (string, error) {
	var gpr github.GitHubPR
	err := json.Unmarshal(bodyAsBytes, &gpr)
	if err != nil {
		log.Printf("failed to Unmarshal body %v", err)
		utils.UnmarshalBodyError.With(prometheus.Labels{"handler": "github"}).Inc()
		return "", err
	}

	if !g.validateAction(gpr.Action) {
		return "", errors.New("unsupported action")
	}
	return fmt.Sprintf("%v is request review for PR:\n%v", gpr.PullRequest.User.Login, gpr.PullRequest.HTMLURL), nil
}

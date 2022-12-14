package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yeitany/webex_bot/pkg/webex"

	"github.com/yeitany/webex_bot/utils"

	"github.com/yeitany/webex_bot/pkg/gcp"

	"github.com/prometheus/client_golang/prometheus"
)

type GCPHandler struct {
	Config *utils.Config
}

func (g *GCPHandler) ServeHttp(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("err %v", err)
		utils.ReadBodyError.With(prometheus.Labels{"handler": "gpc"}).Inc()
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var x gcp.AutoGenerated
	err = json.Unmarshal(body, &x)
	if err != nil {
		log.Printf("err %v", err)
		utils.UnmarshalBodyError.With(prometheus.Labels{"handler": "gpc"}).Inc()
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Print(x.Incident.Documentation.Content)
	webexMessage := webex.WebexMessage{
		Text:    x.Incident.Documentation.Content,
		RoomId:  g.Config.RoomID,
		Token:   g.Config.Token,
		BaseUrl: g.Config.BaseUrl,
	}
	err = webexMessage.Send()
	if err != nil {
		log.Printf("err %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

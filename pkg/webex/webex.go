package webex

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type WebexMessage struct {
	Text    string `json:"text"`
	Token   string `json:"-"`
	RoomId  string `json:"roomId"`
	BaseUrl string `json:"-"`
}

func (w *WebexMessage) Send() error {
	webexMessageBytes, err := json.Marshal(w)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	log.Printf("%v", string(webexMessageBytes))

	req, err := http.NewRequest("POST", w.BaseUrl, bytes.NewReader(webexMessageBytes))
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", w.Token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		webexSendMessageError.With(prometheus.Labels{}).Inc()
		log.Printf("err %v", err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		webexUnmarshalResponse.With(prometheus.Labels{}).Inc()
		log.Printf("err %v", err)
		return err
	}
	log.Printf("%+v", string(body))
	if resp.StatusCode >= 400 {
		webexSendMessageError.With(prometheus.Labels{}).Inc()
		log.Printf("failed sending message to webex status_code: %v", resp.StatusCode)
		return errors.New(string(body))
	}

	return nil
}

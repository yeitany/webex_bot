package webex

import "github.com/prometheus/client_golang/prometheus"

var webexSendMessageError = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "webex_send_message_error",
		Help: "counts the webex sending message error",
	},
	[]string{},
)

var webexUnmarshalResponse = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "webex_unmarshl_response_error",
		Help: "counts the errors of unmarshal responses",
	},
	[]string{},
)

func init() {
	prometheus.MustRegister(webexUnmarshalResponse)
	prometheus.MustRegister(webexSendMessageError)
}

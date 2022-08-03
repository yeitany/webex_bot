package utils

import "github.com/prometheus/client_golang/prometheus"

var ReadBodyError = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "gcp_incident_read_body_errors",
		Help: "counts the gcp incident read body errors",
	},
	[]string{"handler"},
)

var UnmarshalBodyError = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "gcp_unmarashl_body_error",
		Help: "counts the gcp incidents body unmarashl error",
	},
	[]string{"handler"},
)

func init() {
	prometheus.MustRegister(ReadBodyError)
	prometheus.MustRegister(UnmarshalBodyError)
}

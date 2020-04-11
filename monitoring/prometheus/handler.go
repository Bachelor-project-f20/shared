package prometheus

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ServePrometheus() func() {
	return func() {
		go func() {
			fmt.Println("Serving metrics API")

			h := http.NewServeMux()
			h.Handle("/metrics", promhttp.Handler())

			http.ListenAndServe(":9100", h)
		}()
	}
}

package common

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal("xxxx")
		}
		log.Info("xxxxx：" + strconv.Itoa(port))
	}()
}

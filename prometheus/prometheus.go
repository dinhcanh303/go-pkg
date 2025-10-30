package prometheus

import (
	"fmt"
	"go-pkg/thread"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	once    sync.Once
	enabled atomic.Bool
)

// Enabled returns whether the Prometheus agent is enabled.
func Enabled() bool {
	return enabled.Load()
}

// Enable activates the Prometheus agent.
func Enable() {
	enabled.Store(true)
}

// StartAgent starts the Prometheus metrics HTTP server based on the provided configuration.
func StartAgent(conf PrometheusConf) {
	if len(conf.Host) == 0 {
		return
	}
	once.Do(func() {
		Enable()
		thread.GoSafe(func() {
			http.Handle(conf.Path, promhttp.Handler())
			addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
			if err := http.ListenAndServe(addr, nil); err != nil {
				// log.Error(err)
			}
		})
	})
}

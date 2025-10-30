package prometheus

type PrometheusConf struct {
	Host string `json:"host",optional"`
	Path string `json:"path",default=/metrics"`
	Port int    `json:"port",default=9101"`
}

package conf

type Config struct {
	GRPCAddress string
	HTTPAddress string
	PrometheusAddress string

	TLSCert string
	TLSKey string
	TLSHostName string
}

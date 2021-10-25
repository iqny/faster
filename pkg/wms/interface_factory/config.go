package interface_factory

type Config struct {
	AppKey     string
	AppSecret  string
	Session    string
	CustomerId string
	GatewayUrl string
	Secret     string //用于hmac加密
}
type ConfigFunc func(c *Config)

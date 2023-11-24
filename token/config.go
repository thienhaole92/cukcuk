package token

type Config struct {
	url    string
	appID  string
	domain string
	secret string
}

func NewConfig(url string, appID string, domain string, secret string) *Config {
	return &Config{
		url:    url,
		appID:  appID,
		domain: domain,
		secret: secret,
	}
}

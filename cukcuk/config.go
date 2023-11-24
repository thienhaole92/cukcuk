package cukcuk

type Config struct {
	url string
	companyCode string
}

func NewConfig(url string, companyCode string) *Config {
	return &Config{
		url: url,
		companyCode: companyCode,
	}
}

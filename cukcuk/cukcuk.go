package cukcuk

func New(config *Config, auth TokenClient) *Api {
	return &Api{
		config: config,
		auth: auth,
	}
}
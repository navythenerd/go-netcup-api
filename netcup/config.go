package netcup

type Config struct {
	ApiKey         string `json:"apiKey"`
	ApiPassword    string `json:"apiPassword"`
	CustomerNumber string `json:"customerNumber"`
}

func ReadConfigFromFile(cfg *Config) error {

	return nil
}

func ReadConfigFromEnv(cfg *Config) error {

	return nil
}

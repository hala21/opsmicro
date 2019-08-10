package config

type JwtConfig interface {
	GetSecretKey() string
}

type defaultJwtConfig struct {
	SecretKey string `json:"secretKey"`
}

func (j defaultJwtConfig) GetSecretKey() string {
	return j.SecretKey
}

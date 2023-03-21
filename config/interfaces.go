package config

type CORSConfigInterface interface {
	GetOrigins() []string
}

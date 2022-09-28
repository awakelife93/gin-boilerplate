package config

type Config struct {
	requestIdFieldName string
}

var config Config

func RequestIdFieldName() string {
	return config.requestIdFieldName
}

func Initialize() {
	config = Config{
		requestIdFieldName: "requestId",
	}
}

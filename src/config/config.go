package config

type Config struct {
	requestIdFieldName  string
	generateSampleCount int
}

var config Config

func RequestIdFieldName() string {
	return config.requestIdFieldName
}

func GenerateSampleCount() int {
	return config.generateSampleCount
}

func Initialize() {
	config = Config{
		requestIdFieldName:  "requestId",
		generateSampleCount: 100,
	}
}

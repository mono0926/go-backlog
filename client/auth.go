package client

type Auth struct {
	TeamKey string
	ApiKey  string
}

func NewAuth(teamKey string, apiKey string) *Auth {
	return &Auth{TeamKey: teamKey, ApiKey: apiKey}
}

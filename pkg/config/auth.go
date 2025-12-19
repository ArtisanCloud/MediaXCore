package config

type AuthType string

const (
	AuthFlowAuthCode   AuthType = "auth_code"
	AuthFlowClientCred AuthType = "client_credential"
)

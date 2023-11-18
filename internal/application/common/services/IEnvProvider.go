package services

type IEnvProvider interface {
	Get(string) string
}

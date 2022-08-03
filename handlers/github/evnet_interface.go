package handlers

type GithubEvents interface {
	handle(bodyAsBytes []byte) (string, error)
}

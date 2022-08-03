package handlers

type PingHandler struct{}

func (p PingHandler) handle(bodyAsBytes []byte) (string, error) {
	return "", nil
}

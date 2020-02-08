package apiserver

// Apiserver ...
type APIServer struct {
}

// New ...
func new() *APIServer {
	return &APIServer{}
}

// APIServer ...
func (s *APIServer) Start() error {
	return nil
}

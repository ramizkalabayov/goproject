package apiserver

//Apiserver

type APIServer struct {
}

func new() *APIServer {
	return &APIServer{}
}

func (s *APIServer) Start() error {
	return nil
}

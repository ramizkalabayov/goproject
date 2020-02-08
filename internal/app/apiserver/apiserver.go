package apiserver

import (
	"github.com/sirupsen/logrus"
)

// Apiserver ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// APIServer ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return nil
	}
	s.logger.Info("starting api server")
	return nil
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

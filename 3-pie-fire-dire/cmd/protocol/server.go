package protocol

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type server struct {
	grpc *GRPCServer
	rest *RESTServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) WithGRPC() *server {
	s.grpc = NewGRPCServer()
	return s
}

func (s *server) WithREST() *server {
	s.rest = NewRESTServer()
	return s
}

func (s *server) Start() {
	stop := make(chan os.Signal, 1)

	if s.grpc != nil {
		go func() {
			if err := s.grpc.Start(); err != nil {
				stop <- os.Interrupt
			}
		}()
	}

	if s.rest != nil {
		go func() {
			if err := s.rest.Start(); err != nil {
				stop <- os.Interrupt
			}
		}()
	}

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	s.Stop()
}

func (s *server) Stop() {
	var wg sync.WaitGroup

	if s.grpc != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.rest.Stop(); err != nil {
				log.Printf("[*] Error shutting down gRPC server %v\n", err)
			}
			log.Println("[!] gRPC server shutdown complete.")
		}()
	}

	if s.rest != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.rest.Stop(); err != nil {
				log.Printf("[*] Error shutting down REST server: %v\n", err)
			}
			log.Println("[!] REST server shutdown complete.")
		}()
	}

	wg.Wait()
}

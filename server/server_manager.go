package server

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

type startStopper interface {
	Start() error
	Stop(context.Context) error
}

// ServerManager manages the starting and stopping of servers
type ServerManager struct {
	servers         []startStopper
	ShutdownTimeout time.Duration
}

func NewServerManager(servers ...startStopper) *ServerManager {
	return &ServerManager{
		servers: servers,
	}
}

func (s *ServerManager) Start() error {
	done := make(chan struct{})
	errChan := make(chan error)
	var wg sync.WaitGroup

	for _, server := range s.servers {
		wg.Add(1)
		go func(server startStopper) {
			defer wg.Done()
			err := server.Start()
			if err != nil {
				errChan <- err
			}
		}(server)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	for {
		select {
		case <-done:
			return nil
		case err := <-errChan:
			if err != nil {
				stopErr := s.Stop()
				if stopErr != nil {
					return fmt.Errorf("Failed to cleanup after a failed server start: %s", stopErr)
				}
				return err
			}
		}
	}
}

func (s *ServerManager) Stop() error {
	errs := make([]string, 0, len(s.servers))
	for _, server := range s.servers {
		// shut down gracefully, but wait no longer than 5 seconds before exiting
		ctx, cncl := context.WithTimeout(context.Background(), s.ShutdownTimeout)
		defer cncl()

		err := server.Stop(ctx)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to stop all servers: %s", strings.Join(errs, ", "))
	}
	return nil
}

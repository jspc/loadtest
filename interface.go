package golo

import (
	"fmt"
	"net"
	"net/rpc"
)

const (
	// RPCAddr is the default host on which a schedule listens
	// and an agent connects to
	RPCAddr = "127.0.0.1:9999"
)

// NullArg is a set of args that don't do anything
// but that can be put into rpc calls to aid readability
type NullArg struct{}

// Runner is the interface to implement in scheduler
// code; it provides a single function `Run()` which
// takes no arguments, and returns nothing
type Runner interface {
	Run()
}

// Server will expose scheduler code over RPC for agents
// to run and work with.
type Server struct {
	runner Runner
}

// NewServer takes scheduler code which implements the Runner
// interface and returns a Server
func NewServer(r Runner) Server {
	return Server{r}
}

// Run is the RPC interface into scheduler code
func (s Server) Run(_ *NullArg, _ *NullArg) error {
	s.runner.Run()

	return nil
}

// StartListener will start an RPC server on loadtest.RPCAddr
// and register Server ahead of Agents scheduling jobs
func StartListener(server Server) (err error) {
	s, l, err := setupListener(server)
	if err != nil {
		return
	}

	s.Accept(l)

	return fmt.Errorf("Server has gone away")
}

func setupListener(server Server) (s *rpc.Server, l net.Listener, err error) {
	s = rpc.NewServer()
	s.Register(&server)

	l, err = net.Listen("tcp", RPCAddr)

	return
}

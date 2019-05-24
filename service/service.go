package service

import (
	"crypto/tls"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/dialogs/dialog-go-lib/logger"
	pkgerr "github.com/pkg/errors"
	"google.golang.org/grpc"
)

// service base object
type service struct {
	interrupt chan os.Signal
	addr      string
	mu        sync.RWMutex
}

// create base service object
func newService() *service {
	return &service{
		interrupt: make(chan os.Signal),
	}
}

// Close service (io.Closer implementation)
func (s *service) Close() (err error) {
	close(s.interrupt)
	return nil
}

// SetAddr set listener address
func (s *service) SetAddr(val string) {

	s.mu.Lock()
	s.addr = val
	s.mu.Unlock()
}

// GetAddr returns listener address
func (s *service) GetAddr() (addr string) {

	s.mu.RLock()
	addr = s.addr
	s.mu.RUnlock()

	return
}

func (s *service) serve(name, addr string, run func(retval chan<- error), stop func(*logger.Logger)) error {

	l, err := logger.New(&logger.Config{}, map[string]interface{}{
		name: addr,
	})
	if err != nil {
		return pkgerr.Wrap(err, "logger: "+name)
	}

	signal.Notify(s.interrupt, os.Interrupt, syscall.SIGTERM)

	retval := make(chan error)
	go run(retval)

	l.Info("The service is ready to listen and serve")

	killSignal := <-s.interrupt
	switch killSignal {
	case os.Interrupt:
		l.Info("Got SIGINT...")
	case syscall.SIGTERM:
		l.Info("Got SIGTERM...")
	}

	l.Info("The service is shutting down...")
	stop(l)
	l.Info("The service is done")

	return <-retval
}

// PingConn ping tcp connection
func PingConn(addr string, tries int, timeout time.Duration, tlsConf *tls.Config) (err error) {

	for i := 0; i < tries; i++ {
		var conn io.Closer
		if tlsConf != nil {
			dialer := &net.Dialer{KeepAlive: timeout}
			conn, err = tls.DialWithDialer(dialer, "tcp", addr, tlsConf)
		} else {
			conn, err = net.DialTimeout("tcp", addr, timeout)
		}
		if err == nil {
			defer conn.Close()
			return
		}

		if isLast := i == tries-1; !isLast {
			time.Sleep(time.Second)
		}
	}

	return
}

// PingGRPC ping grpc server
func PingGRPC(addr string, tries int, opts ...grpc.DialOption) (err error) {

	for i := 0; i < tries; i++ {
		var conn *grpc.ClientConn
		conn, err = grpc.Dial(addr, opts...)
		if err == nil {
			defer conn.Close()
			return
		}

		if isLast := i == tries-1; !isLast {
			time.Sleep(time.Second)
		}
	}

	return
}

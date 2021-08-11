package neo

import (
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

// Neo is an gRPC micro server implementation
type Neo struct {
	server                   *grpc.Server
	address                  string
	listener                 net.Listener
	unaryServerInterceptors  []grpc.UnaryServerInterceptor
	streamServerInterceptors []grpc.StreamServerInterceptor
	reflection               bool
	tls                      bool
	mutualTLS                bool
	serverCertFile           string
	serverKeyFile            string
	clientCACertFile         string
}

// New creates a new instance of Neo
func New(opts ...Option) (*Neo, error) {
	neo := &Neo{
		address: ":9090",
	}
	for _, f := range opts {
		f(neo)
	}

	if neo.listener == nil {
		lis, err := defaultListener(neo.address)
		if err != nil {
			return nil, err
		}
		neo.listener = lis
	}

	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(neo.unaryServerInterceptors...),
		grpc.ChainStreamInterceptor(neo.streamServerInterceptors...),
	}

	if neo.tls {
		tls, err := loadTLSCredentials(neo.mutualTLS, neo.serverCertFile, neo.serverKeyFile, neo.clientCACertFile)
		if err != nil {
			return nil, err
		}
		serverOptions = append(serverOptions, grpc.Creds(tls))
	}

	neo.server = grpc.NewServer(serverOptions...)

	return neo, nil
}

// Register function signature
type Register func(server *grpc.Server)

// Register allows you to register your Proto methods
func (n *Neo) Register(register Register) *Neo {
	register(n.server)
	return n
}

// Serve stars the gRPC server
func (n *Neo) Serve() error {
	if n.reflection {
		reflection.Register(n.server)
	}
	if err := n.server.Serve(n.listener); err != nil {
		return err
	}
	return nil
}

func defaultListener(address string) (net.Listener, error) {
	return net.Listen("tcp", address)
}

// GracefulStop stops the gRPC server gracefully. It stops the server from accepting new connections and
// RPCs and blocks until all the pending RPCs are finished.
func (n *Neo) GracefulStop() {
	n.server.GracefulStop()
}

// Stop stops the gRPC server. It immediately closes all open connections and listeners. It cancels all active RPCs on
// the server side and the corresponding pending RPCs on the client side will get notified by connection errors.
func (n *Neo) Stop() {
	n.server.Stop()
}

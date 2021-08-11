package neo

import (
	"net"

	"google.golang.org/grpc"
)

// Option functions for Neo
type Option func(*Neo)

// WithListener sets the Listener for Neo
func WithListener(listener net.Listener) Option {
	return func(neo *Neo) {
		neo.listener = listener
	}
}

// WithAddress sets the address for neo
func WithAddress(address string) Option {
	return func(neo *Neo) {
		neo.address = address
	}
}

// WithUnaryServerInterceptors adds UnaryServerInterceptors
func WithUnaryServerInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(neo *Neo) {
		neo.unaryServerInterceptors = interceptors
	}
}

// WithStreamServerInterceptors adds StreamServerInterceptors
func WithStreamServerInterceptors(interceptors ...grpc.StreamServerInterceptor) Option {
	return func(neo *Neo) {
		neo.streamServerInterceptors = interceptors
	}
}

// WithReflection enabled grpc reflection for proto definitions
func WithReflection() Option {
	return func(neo *Neo) {
		neo.reflection = true
	}
}

// TODO: Implement With TLS functions
// https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph

// WithServerSideTLS ...
func WithServerSideTLS(serverCertFile, serverKeyFile, clientCACertFile string) Option {
	return func(neo *Neo) {
		neo.tls = true
		neo.mutualTLS = false
		neo.serverCertFile = serverCertFile
		neo.serverKeyFile = serverKeyFile
		neo.clientCACertFile = clientCACertFile
	}
}

// WithMutualTLS ...
func WithMutualTLS(serverCertFile, serverKeyFile, clientCACertFile string) Option {
	return func(neo *Neo) {
		neo.tls = true
		neo.mutualTLS = true
		neo.serverCertFile = serverCertFile
		neo.serverKeyFile = serverKeyFile
		neo.clientCACertFile = clientCACertFile
	}
}

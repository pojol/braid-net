package server

import (
	"github.com/pojol/braid-go/depend/tracer"
	"google.golang.org/grpc"
)

// Parm Service 配置
type Parm struct {
	ListenAddr string

	Tracer tracer.ITracer

	Interceptors []grpc.UnaryServerInterceptor
}

// Option config wraps
type Option func(*Parm)

// WithListen 服务器侦听地址配置
func WithListen(address string) Option {
	return func(c *Parm) {
		c.ListenAddr = address
	}
}

func AppendInterceptors(interceptor grpc.UnaryServerInterceptor) Option {
	return func(c *Parm) {
		c.Interceptors = append(c.Interceptors, interceptor)
	}
}

func WithTracer(t tracer.ITracer) Option {
	return func(c *Parm) {
		c.Tracer = t
	}
}
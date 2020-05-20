package dispatcher

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/pojol/braid/log"

	"github.com/opentracing/opentracing-go"

	"github.com/pojol/braid/cache/pool"
	"github.com/pojol/braid/service/balancer"
	"github.com/pojol/braid/service/discover"
	"github.com/pojol/braid/tracer"
	"google.golang.org/grpc"
)

type (

	// IClient client的抽象接口
	IClient interface {
		Run()
		GetConn(target string) (*pool.ClientConn, error)
		Close()
	}

	// Client 调用器
	Client struct {
		cfg config

		discov discover.IDiscover

		refushTick *time.Ticker

		poolMgr sync.Map
		sync.Mutex
	}
)

var (
	c *Client

	// ErrServiceNotAvailable 服务不可用，通常是因为没有查询到中心节点(coordinate)
	ErrServiceNotAvailable = errors.New("caller service not available")

	// ErrConfigConvert 配置转换失败
	ErrConfigConvert = errors.New("Convert linker config")

	// ErrCantFindNode 在注册中心找不到对应的服务节点
	ErrCantFindNode = errors.New("Can't find service node in center")
)

// New 构建指针
func New(name string, consulAddress string, opts ...Option) IClient {
	const (
		defaultPoolInitNum  = 8
		defaultPoolCapacity = 32
		defaultPoolIdle     = 120
		defaultTracing      = false
	)

	c = &Client{
		cfg: config{
			ConsulAddress: consulAddress,
			PoolInitNum:   defaultPoolInitNum,
			PoolCapacity:  defaultPoolCapacity,
			PoolIdle:      defaultPoolIdle,
			Tracing:       defaultTracing,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	// 这里后面需要做成可选项
	balancer.New()
	c.discov = discover.New(name, consulAddress)

	return c
}

// GetConn 获取rpc client连接
func (c *Client) GetConn(target string) (*pool.ClientConn, error) {
	var caConn *pool.ClientConn
	var caPool *pool.GRPCPool

	address, err := c.findNode(target, "")
	if err != nil {
		return nil, err
	}

	caPool, err = c.pool(address)
	if err != nil {
		return nil, err
	}

	connCtx, connCancel := context.WithTimeout(context.Background(), time.Second)
	defer connCancel()
	caConn, err = caPool.Get(connCtx)
	if err != nil {
		return nil, err
	}

	return caConn, nil
}

// Find 通过查找器获取目标
func (c *Client) findNode(target string, key string) (string, error) {
	var address string
	var err error
	var nod *balancer.Node

	wb, err := balancer.GetGroup(target)
	if err != nil {
		goto EXT
	}

	nod, err = wb.Next()
	if err != nil {
		goto EXT
	}

	address = nod.Address

EXT:
	if err != nil {
		// log
		log.SysError("rpcClient", "findNode", err.Error())
	}

	return address, err
}

// Pool 获取grpc连接池
func (c *Client) pool(address string) (p *pool.GRPCPool, err error) {

	factory := func() (*grpc.ClientConn, error) {
		var conn *grpc.ClientConn
		var err error

		if c.cfg.Tracing {
			interceptor := tracer.ClientInterceptor(opentracing.GlobalTracer())
			conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor))
		} else {
			conn, err = grpc.Dial(address, grpc.WithInsecure())
		}

		if err != nil {
			return nil, err
		}

		return conn, nil
	}

	pi, ok := c.poolMgr.Load(address)
	if !ok {
		p, err = pool.NewGRPCPool(factory, c.cfg.PoolInitNum, c.cfg.PoolCapacity, c.cfg.PoolIdle)
		if err != nil {
			goto EXT
		}

		c.poolMgr.Store(address, p)
		pi = p
	}

	p = pi.(*pool.GRPCPool)

EXT:
	if err != nil {
		log.SysError("rpcClient", "pool", err.Error())
	}

	return p, err
}

// Run 执行服务发现逻辑
func (c *Client) Run() {
	c.discov.Run()
}

// Close 关闭服务发现逻辑
func (c *Client) Close() {
	c.discov.Close()
}
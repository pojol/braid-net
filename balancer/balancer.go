package balancer

import (
	"errors"
	"sync"
)

// SelectorCfg 负载均衡选择器配置
type SelectorCfg struct {
}

// Selector 负载均衡选择器
type Selector struct {
	m sync.Map
}

var (
	// ErrBalanceEmpty 没有权重节点
	ErrBalanceEmpty = errors.New("weighted node list is empty")
)

// Node 权重节点
type Node struct {
	ID      string
	Address string

	// 被选中的次数
	Tick int
	// 节点的权重值
	Weight float32
}

// Nodes 权重节点队列
type Nodes []Node

func (s Nodes) Len() int { return len(s) }

func (s Nodes) Less(i, j int) bool {
	return (float32(s[i].Tick) * s[i].Weight) < (float32(s[j].Tick) * s[j].Weight)
}

func (s Nodes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Balancer 均衡器
type Balancer interface {
	// Add 添加一个新节点
	Add(Node)
	// Rmv 移除一个旧节点
	Rmv(string)
	// Next 获取一个节点
	Next() (*Node, error)
}

var (
	selector *Selector

	defaultSelectorCfg = SelectorCfg{}
)

// New 初始化负载均衡选择器
func New() *Selector {
	selector = &Selector{}
	return selector
}

// Init 初始化均衡器
func (s *Selector) Init(cfg interface{}) error {
	return nil
}

func (s *Selector) Run() {

}

func (s *Selector) Close() {

}

// GetSelector 获取负载均衡选择器
func GetSelector(boxName string) Balancer {
	return selector.group(boxName)
}

// Group 获取组
func (s *Selector) group(boxName string) Balancer {

	b, ok := s.m.Load(boxName)
	if !ok {
		b = &LeastConnBalancer{
			BoxName: boxName,
		}

		s.m.Store(boxName, b)
	}

	return b.(Balancer)
}
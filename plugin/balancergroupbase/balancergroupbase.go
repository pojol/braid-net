package balancergroupbase

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/pojol/braid/module"
	"github.com/pojol/braid/module/balancer"
	"github.com/pojol/braid/module/discover"
	"github.com/pojol/braid/module/logger"
	"github.com/pojol/braid/module/mailbox"
)

const (
	// Name 基础的负载均衡容器实现
	Name = "BalancerGroupBase"
)

type baseBalancerGroupBuilder struct {
	opts []interface{}
}

func (b *baseBalancerGroupBuilder) Name() string {
	return Name
}

func (b *baseBalancerGroupBuilder) Type() string {
	return module.TyBalancerGroup
}

func newBaseBalancerGroup() module.Builder {
	return &baseBalancerGroupBuilder{}
}

func (b *baseBalancerGroupBuilder) AddOption(opt interface{}) {
	b.opts = append(b.opts, opt)
}

func (b *baseBalancerGroupBuilder) Build(serviceName string, mb mailbox.IMailbox, logger logger.ILogger) (module.IModule, error) {

	p := Parm{}
	for _, opt := range b.opts {
		opt.(Option)(&p)
	}

	bbg := &baseBalancerGroup{
		serviceName: serviceName,
		parm:        p,
		mb:          mb,
		logger:      logger,
		group:       make(map[string]*targetBalancerMap),
	}
	for _, strategy := range p.strategies {
		bbg.group[strategy] = &targetBalancerMap{
			targets: make(map[string]balancer.IBalancer),
		}
	}

	return bbg, nil
}

type targetBalancerMap struct {
	targets map[string]balancer.IBalancer
}

func (tbm *targetBalancerMap) get(target string) (balancer.IBalancer, error) {
	if _, ok := tbm.targets[target]; ok {
		return tbm.targets[target], nil
	}

	return nil, errors.New("can't find balancer, with target")
}

func (tbm *targetBalancerMap) exist(serviceName string) bool {
	if _, ok := tbm.targets[serviceName]; !ok {
		return false
	}

	return true
}

type baseBalancerGroup struct {
	mb          mailbox.IMailbox
	serviceName string

	parm Parm

	addConsumer mailbox.IConsumer
	rmvConsumer mailbox.IConsumer
	upConsumer  mailbox.IConsumer

	logger logger.ILogger

	// Strategy, Target
	group map[string]*targetBalancerMap

	lock sync.RWMutex
}

func (bbg *baseBalancerGroup) Init() {
	var err error

	bbg.addConsumer, err = bbg.mb.ProcSub(discover.AddService).AddShared()
	if err != nil {
		panic(err)
	}

	bbg.rmvConsumer, err = bbg.mb.ProcSub(discover.RmvService).AddShared()
	if err != nil {
		panic(err)
	}

	bbg.upConsumer, err = bbg.mb.ProcSub(discover.UpdateService).AddShared()
	if err != nil {
		panic(err)
	}

}

func (bbg *baseBalancerGroup) Run() {

	bbg.addConsumer.OnArrived(func(msg *mailbox.Message) error {

		nod := discover.Node{}
		json.Unmarshal(msg.Body, &nod)

		bbg.lock.Lock()
		defer bbg.lock.Unlock()

		for strategy := range bbg.group {

			if !bbg.group[strategy].exist(nod.Name) {
				b := balancer.GetBuilder(strategy)
				ib, _ := b.Build(bbg.logger)
				bbg.group[strategy].targets[nod.Name] = ib
			}

			bbg.group[strategy].targets[nod.Name].Add(nod)
		}

		return nil
	})

	bbg.rmvConsumer.OnArrived(func(msg *mailbox.Message) error {

		nod := discover.Node{}
		json.Unmarshal(msg.Body, &nod)

		bbg.lock.Lock()
		defer bbg.lock.Unlock()

		for k := range bbg.group {
			if _, ok := bbg.group[k]; ok {
				b, err := bbg.group[k].get(nod.Name)
				if err != nil {
					bbg.logger.Errorf("remove service err %s", err.Error())
					break
				}

				b.Rmv(nod)
			}
		}

		return nil
	})

	bbg.upConsumer.OnArrived(func(msg *mailbox.Message) error {

		nod := discover.Node{}
		json.Unmarshal(msg.Body, &nod)

		bbg.lock.Lock()
		defer bbg.lock.Unlock()

		for k := range bbg.group {
			if _, ok := bbg.group[k]; ok {
				b, err := bbg.group[k].get(nod.Name)
				if err != nil {
					bbg.logger.Errorf("update service err %s", err.Error())
					break
				}

				b.Update(nod)
			}
		}

		return nil
	})

}

func (bbg *baseBalancerGroup) Pick(ty string, target string) (discover.Node, error) {

	bbg.lock.RLock()
	defer bbg.lock.RUnlock()

	var nod discover.Node

	if _, ok := bbg.group[ty]; ok {

		b, err := bbg.group[ty].get(target)
		if err != nil {
			return nod, err
		}

		return b.Pick()
	}

	return nod, errors.New("can't find balancer, with strategy")
}

func (bbg *baseBalancerGroup) Close() {
	bbg.addConsumer.Exit()
	bbg.rmvConsumer.Exit()
	bbg.upConsumer.Exit()
}

func init() {
	module.Register(newBaseBalancerGroup())
}
package pubsubproc

import (
	"sync"

	"github.com/pojol/braid/internal/braidsync"
	"github.com/pojol/braid/module/pubsub"
)

const (
	// PubsubName 进程内的消息通知
	PubsubName = "ProcPubsub"
)

type procPubsubBuilder struct {
}

func newProcPubsub() pubsub.Builder {
	return &procPubsubBuilder{}
}

func (*procPubsubBuilder) Build() pubsub.IPubsub {

	ps := &procPubsub{
		subscriber: make(map[string]pubsub.ISubscriber),
	}

	return ps
}

func (*procPubsubBuilder) Name() string {
	return PubsubName
}

// Consumer 消费者
type procConsumer struct {
	buff   *braidsync.Unbounded
	exitCh *braidsync.Switch
}

func (c *procConsumer) OnArrived(handler pubsub.HandlerFunc) {

	go func() {
		for {
			select {
			case msg := <-c.buff.Get():
				handler(msg.(*pubsub.Message))
				c.buff.Load()
			case <-c.exitCh.Done():
			}

			if c.exitCh.HasOpend() {
				return
			}
		}
	}()

}

func (c *procConsumer) Exit() {
	c.exitCh.Open()
}

func (c *procConsumer) IsExited() bool {
	return c.exitCh.HasOpend()
}

func (c *procConsumer) PutMsg(msg *pubsub.Message) {
	c.buff.Put(msg)
}

type procSubscriber struct {
	group []pubsub.IConsumer
	sync.Mutex
}

func (ps *procSubscriber) AddConsumer() pubsub.IConsumer {

	ps.Lock()
	defer ps.Unlock()

	c := &procConsumer{
		buff:   braidsync.NewUnbounded(),
		exitCh: braidsync.NewSwitch(),
	}

	ps.group = append(ps.group, c)

	return c
}

func (ps *procSubscriber) AppendConsumer() pubsub.IConsumer {
	return nil
}

func (ps *procSubscriber) PutMsg(msg *pubsub.Message) {
	for i := 0; i < len(ps.group); i++ {
		if ps.group[i].IsExited() {
			ps.group = append(ps.group[:i], ps.group[i+1:]...)
			i--
		} else {
			ps.group[i].PutMsg(msg)
		}
	}
}

// procPubsub 进程内使用的pub-sub消息分发队列
type procPubsub struct {
	sync.RWMutex
	subscriber map[string]pubsub.ISubscriber
}

func (kps *procPubsub) Sub(topic string) pubsub.ISubscriber {
	kps.Lock()
	defer kps.Unlock()

	if _, ok := kps.subscriber[topic]; !ok {
		kps.subscriber[topic] = &procSubscriber{}
	}

	return kps.subscriber[topic]
}

func (kps *procPubsub) Pub(topic string, msg *pubsub.Message) {
	kps.RLock()
	defer kps.RUnlock()

	if _, ok := kps.subscriber[topic]; ok {
		kps.subscriber[topic].PutMsg(msg)
	}

}

func init() {
	pubsub.Register(newProcPubsub())
}

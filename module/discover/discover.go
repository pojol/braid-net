// discover 模块，主要用于服务发现
//
// 这个模块主要通过 topic : ServiceUpdate 消息，发布集群中相关服务的变更信息
package discover

import (
	"encoding/json"

	"github.com/pojol/braid-go/module"
	"github.com/pojol/braid-go/module/mailbox"
)

const (
	ServiceUpdate = "discover.serviceUpdate"

	// EventAddService 有一个新的服务加入到集群
	EventAddService = "event_add_service"

	// EventRemoveService 有一个旧的服务从集群中退出
	EventRemoveService = "event_remove_service"

	// EventUpdateService 有一个旧的服务产生了信息的变更（通常是指权重
	EventUpdateService = "event_update_service"
)

// Node 发现节点结构
type Node struct {
	ID string
	// 负载均衡节点的名称，这个名称主要用于均衡节点分组。
	Name    string
	Address string

	// 节点的权重值
	Weight int
}

type UpdateMsg struct {
	Nod   Node
	Event string
}

func EncodeUpdateMsg(event string, nod Node) *mailbox.Message {
	byt, _ := json.Marshal(&UpdateMsg{
		Event: event,
		Nod:   nod,
	})

	return &mailbox.Message{
		Body: byt,
	}
}

func DecodeUpdateMsg(msg *mailbox.Message) UpdateMsg {
	dmsg := UpdateMsg{}
	json.Unmarshal(msg.Body, &dmsg)
	return dmsg
}

// IDiscover discover interface
type IDiscover interface {
	module.IModule
}

## Braid
**Braid** 提供统一的交互模型，通过注册`模块`（支持自定义），构建属于自己的微服务。

---

[![Go Report Card](https://goreportcard.com/badge/github.com/pojol/braid)](https://goreportcard.com/report/github.com/pojol/braid)
[![drone](http://123.207.198.57:8001/api/badges/pojol/braid/status.svg?branch=develop)](dev)
[![codecov](https://codecov.io/gh/pojol/braid/branch/master/graph/badge.svg)](https://codecov.io/gh/pojol/braid)

<img src="https://i.postimg.cc/wB1K7w0Z/braid.png" width="700">


### 交互模型
> braid.Mailbox 统一的交互模型

| 共享（多个消息副本 | 竞争（只被消费一次 | 进程内 | 集群内 | 发布 | 订阅 |
| ---- | ---- | ---- | ---- | ---- | ---- |
|Shared | Competition | Proc | Cluster | Pub | Sub |

> `范例` 在集群内`订阅`一个共享型的消息

```go
// 订阅一个信道`topic` 这个信道在进程（Proc 内广播（Shared
consumer := braid.Mailbox().Sub(mailbox.Proc, topic).Shared()
// 注册消息到达函数（线程安全
consumer.OnArrived(func (msg *mailbox.Message) error {
  return nil
})
```



### 微服务
> braid.Module 默认提供的微服务组件

|**Discover**|**Balancer**|**Elector**|**RPC**|**Tracer**|**LinkCache**|
|-|-|-|-|-|-|
|服务发现|负载均衡|选举|RPC|分布式追踪|链路缓存|
|discoverconsul|balancerrandom|electorconsul|grpc-client|jaegertracer|linkerredis|
||[balancerswrr](https://github.com/pojol/braid/wiki/balancerswrr)|electork8s|grpc-server|||



### 构建
> 通过注册模块(braid.Module)，构建braid的运行环境。

```go
b, _ := braid.New(ServiceName)

// 注册插件
b.RegistModule(
  braid.Discover(         // Discover 模块
    discoverconsul.Name,  // 模块名（基于consul实现的discover模块，通过模块名可以获取到模块的构建器
    discoverconsul.WithConsulAddr(consulAddr)), // 模块的可选项
  braid.GRPCClient(grpcclient.Name),
  braid.Elector(
    electorconsul.Name,
    electorconsul.WithConsulAddr(consulAddr),
  ),
  braid.LinkCache(linkerredis.Name),
  braid.JaegerTracing(tracer.WithHTTP(jaegerAddr), tracer.WithProbabilistic(0.01)))

b.Init()  // 初始化注册在braid中的模块
b.Run()   // 运行
defer b.Close() // 释放
```



#### Wiki
> Wiki 中提供了一个较为详细的Guide，用于帮助用户更加全面的理解braid的设计理念

https://github.com/pojol/braid/wiki

#### Sample
https://github.com/pojol/braid-sample



#### Web
* 流向图
> 用于监控链路上的连接数以及分布情况

```shell
$ docker pull braidgo/sankey:latest
$ docker run -d -p 8888:8888/tcp braidgo/sankey:latest \
    -consul http://172.17.0.1:8500 \
    -redis redis://172.17.0.1:6379/0
```
<img src="https://i.postimg.cc/sX0xHZmF/image.png" width="600">


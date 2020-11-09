## Braid
**Braid** 提供统一的交互模型，通过注册`模块`（支持自定义），构建属于自己的微服务。

---

[![Go Report Card](https://goreportcard.com/badge/github.com/pojol/braid)](https://goreportcard.com/report/github.com/pojol/braid)
[![drone](http://123.207.198.57:8001/api/badges/pojol/braid/status.svg?branch=develop)](dev)
[![codecov](https://codecov.io/gh/pojol/braid/branch/master/graph/badge.svg)](https://codecov.io/gh/pojol/braid)

<img src="https://i.postimg.cc/cLKDkf03/image.png" width="800">

### 微服务
> braid.Module 默认提供的微服务组件

|  服务  | 简介  |
|  ----  | ----  | 
| **Discover**  | 发现插件，主要提供 `Add` `Rmv` `Update` 等接口 |
| **Elector** | 选举插件，主要提供 `Wait` `Slave` `Master` 等接口 |
| **Rpc** | RPC插件，主要用于发起RPC请求 `Invoke` 和开启RPC Server |
| **Tracer** | 分布式追踪插件，监控分布式系统中的调用细节，目前有`grpc` `echo` `redis` 等追踪 |
| **LinkCache** | 服务访问链路缓存插件，主要用于缓存token（用户唯一凭证）的链路信息 |



### 交互模型
> braid.Mailbox 统一的交互模型

* **同步**
> 在braid中目前只提供了这`唯一的一个`同步语义的接口,用于发起rpc调用

```go
// ctx 上下文
// targetService 目标服务 例（`login`
// methon 请求目标函数
// token 用户唯一凭证（可为空
braid.Invoke(ctx, targetService, methon, token, args, reply)
```

* **异步**
> braid中异步语义的简介，在内部或是将来功能的扩充，都应该优先使用异步语义

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


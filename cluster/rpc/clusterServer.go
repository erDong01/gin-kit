package rpc

import (
	"context"
	"github.com/erDong01/micro-kit/actor"
	"github.com/erDong01/micro-kit/cluster/common"
	"github.com/erDong01/micro-kit/network"
	"github.com/erDong01/micro-kit/rpc"
	"github.com/erDong01/micro-kit/tools"
	"sync"
)

type (
	HashClusterMap       map[uint32]*common.ClusterInfo
	HashClusterSocketMap map[uint32]*common.ClusterInfo
	//ClusterServer 集群服务器
	ClusterServer struct {
		actor.Actor
		*Service         //集群注册
		clusterMap       HashClusterMap
		clusterSocketMap HashClusterSocketMap
		clusterLocker    *sync.RWMutex
		service          *network.ServerSocket //socket管理
		hashRing         *tools.HashRing       //hash一致性
	}

	IClusterServer interface {
		InitService(info *common.ClusterInfo, Endpoints []string)
		RegisterClusterCall() //注册集群通用回调
		AddCluster(info *common.ClusterInfo)
		DelCluster(info *common.ClusterInfo)
		GetCluster(head rpc.RpcHead) *common.ClusterInfo
		GetClusterBySocket(uint32) *common.ClusterInfo

		BindServer(socket *network.ServerSocket)
		SendMsg(rpc.RpcHead, string, ...interface{}) //发送给集群特定服务器
		Send(rpc.RpcHead, []byte)                    //发送给集群特定服务器

		RandomCluster(head rpc.RpcHead) rpc.RpcHead //随机分配

		sendPoint(rpc.RpcHead, []byte)               //发送给集群特定服务器
		balanceSend(rpc.RpcHead, []byte)             //负载给集群特定服务器
		boardCastSend(head rpc.RpcHead, buff []byte) //给集群广播
	}
)

func (this *ClusterServer) InitService(info *common.ClusterInfo, Endpoints []string) {
	this.clusterLocker = &sync.RWMutex{}
	//注册服务器
	this.Service = NewService(info, Endpoints)
	this.clusterMap = make(HashClusterMap)
	this.clusterSocketMap = make(HashClusterSocketMap)
	this.hashRing = tools.NewHashRing()
}

func (this *ClusterServer) RegisterClusterCall() {
	this.RegisterCall("COMMON_RegisterRequest", func(ctx context.Context, info *common.ClusterInfo) {
		pServerInfo := info
		pServerInfo.SocketId = this.GetRpcHead(ctx).SocketId

		this.AddCluster(pServerInfo)
	})

	//链接断开
	this.RegisterCall("DISCONNECT", func(ctx context.Context, socketId uint32) {
		pCluster := this.GetClusterBySocket(socketId)
		if pCluster != nil {
			this.DelCluster(pCluster)
		}
	})
}

func (this *ClusterServer) AddCluster(info *common.ClusterInfo) {
	this.clusterLocker.Lock()
	this.clusterMap[info.Id()] = info
	this.clusterSocketMap[info.SocketId] = info
	this.clusterLocker.Unlock()
	this.hashRing.Add(info.IpString())
	this.service.SendMsg(rpc.RpcHead{SocketId: info.SocketId}, "COMMON_RegisterResponse")
	switch info.Type {
	case rpc.SERVICE_GATESERVER:
		tools.GLOG.Printf("ADD Gate SERVER: [%d]-[%s:%d]", info.SocketId, info.Ip, info.Port)
	}
}

func (this *ClusterServer) DelCluster(info *common.ClusterInfo) {
	this.clusterLocker.RLock()
	_, bEx := this.clusterMap[info.Id()]
	this.clusterLocker.RUnlock()
	if bEx {
		this.clusterLocker.Lock()
		delete(this.clusterMap, info.Id())
		delete(this.clusterSocketMap, info.SocketId)
		this.clusterLocker.Unlock()
	}

	this.hashRing.Remove(info.IpString())
	tools.GLOG.Printf("服务器断开连接socketid[%d]", info.SocketId)
	switch info.Type {
	case rpc.SERVICE_GATESERVER:
		tools.GLOG.Printf("与Gate服务器断开连接,id[%d]", info.SocketId)
	}
}

func (this *ClusterServer) GetCluster(head rpc.RpcHead) *common.ClusterInfo {
	this.clusterLocker.RLock()
	defer this.clusterLocker.RUnlock()
	pClient, bEx := this.clusterMap[head.ClusterId]
	if bEx {
		return pClient
	}
	return nil
}

func (this *ClusterServer) GetClusterBySocket(socketId uint32) *common.ClusterInfo {
	this.clusterLocker.RLock()
	defer this.clusterLocker.RUnlock()
	pClient, bEx := this.clusterSocketMap[socketId]
	if bEx {
		return pClient
	}
	return nil
}

func (this *ClusterServer) BindServer(pService *network.ServerSocket) {
	this.service = pService
}

func (this *ClusterServer) sendPoint(head rpc.RpcHead, buff []byte) {
	pCluster := this.GetCluster(head)
	if pCluster != nil {
		head.SocketId = pCluster.SocketId
		this.service.Send(head, buff)
	}
}

func (this *ClusterServer) balanceSend(head rpc.RpcHead, buff []byte) {
	_, head.ClusterId = this.hashRing.Get64(head.Id)
	pCluster := this.GetCluster(head)
	if pCluster != nil {
		head.SocketId = pCluster.SocketId
		this.service.Send(head, buff)
	}
}

func (this *ClusterServer) boardCastSend(head rpc.RpcHead, buff []byte) {
	clusterList := []*common.ClusterInfo{}
	this.clusterLocker.RLock()
	for _, v := range this.clusterMap {
		clusterList = append(clusterList, v)
	}
	this.clusterLocker.RUnlock()
	for _, v := range clusterList {
		head.SocketId = v.SocketId
		this.service.Send(head, buff)
	}
}

func (this *ClusterServer) SendMsg(head rpc.RpcHead, funcName string, params ...interface{}) {
	buff := rpc.Marshal(head, funcName, params...)
	this.Send(head, buff)
}

func (this *ClusterServer) Send(head rpc.RpcHead, buff []byte) {
	if head.DestServerType != rpc.SERVICE_GATESERVER {
		this.balanceSend(head, buff)
	} else {
		switch head.SendType {
		case rpc.SEND_BALANCE:
			this.balanceSend(head, buff)
		case rpc.SEND_POINT:
			this.sendPoint(head, buff)
		default:
			this.boardCastSend(head, buff)
		}
	}
}

func (this *ClusterServer) RandomCluster(head rpc.RpcHead) rpc.RpcHead {
	if head.Id == 0 {
		head.Id = int64(uint32(tools.RAND.RandI(1, 0xFFFFFFFF)))
	}
	_, head.ClusterId = this.hashRing.Get64(head.Id)
	pCluster := this.GetCluster(head)
	if pCluster != nil {
		head.SocketId = pCluster.SocketId
	}
	return head
}

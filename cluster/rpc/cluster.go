package rpc

import (
	"context"
	"github.com/erDong01/micro-kit/actor"
	"github.com/erDong01/micro-kit/cluster/common"
	"github.com/erDong01/micro-kit/cluster/etcdv3"
	"github.com/erDong01/micro-kit/network"
	"github.com/erDong01/micro-kit/rpc"
	"github.com/erDong01/micro-kit/tools"
	"github.com/erDong01/micro-kit/tools/vector"
	"reflect"
	"sync"
)

type (
	Service    etcdv3.Service
	Master     etcdv3.Master
	Snowflake  etcdv3.Snowflake
	PlayerRaft etcdv3.PlayerRaft
	//IClusterPacket 集群包管理
	IClusterPacket interface {
		actor.IActor
		SetClusterId(uint32)
	}

	ClusterNode struct {
		*network.ClientSocket
		*common.ClusterInfo
	}

	//Cluster 集群客户端
	Cluster struct {
		actor.Actor
		clusterMap     map[uint32]*ClusterNode
		clusterLocker  *sync.RWMutex
		packet         IClusterPacket
		master         *Master
		hashRing       *tools.HashRing //hash一致性
		clusterInfoMap map[uint32]*common.ClusterInfo
		packetFuncList *vector.Vector //call back
	}

	ICluster interface {
		//actor.IActor
		Init(info *common.ClusterInfo, Endpoints []string)
		AddCluster(info *common.ClusterInfo)
		DelCluster(info *common.ClusterInfo)
		GetCluster(head rpc.RpcHead) *ClusterNode

		BindPacket(IClusterPacket)
		BindPacketFunc(network.PacketFunc)
		SendMsg(rpc.RpcHead, string, ...interface{}) //发送给集群特定服务器
		Send(rpc.RpcHead, []byte)                    //发送给集群特定服务器

		RandomCluster(head rpc.RpcHead) rpc.RpcHead ///随机分配

		sendPoint(rpc.RpcHead, []byte)     //发送给集群特定服务器
		balanceSend(rpc.RpcHead, []byte)   //负载给集群特定服务器
		boardCastSend(rpc.RpcHead, []byte) //给集群广播
	}
)

//NewService 注册服务器
func NewService(info *common.ClusterInfo, Endpoints []string) *Service {
	service := &etcdv3.Service{}
	service.Init(info, Endpoints)
	return (*Service)(service)
}

//NewMaster 监控服务器
func NewMaster(info *common.ClusterInfo, Endpoints []string, pActor actor.IActor) *Master {
	master := &etcdv3.Master{}
	master.Init(info, Endpoints, pActor)
	return (*Master)(master)
}

//NewSnowflake uuid生成器
func NewSnowflake(Endpoints []string) *Snowflake {
	uuid := &etcdv3.Snowflake{}
	uuid.Init(Endpoints)
	return (*Snowflake)(uuid)
}

func (this *Cluster) Init(info *common.ClusterInfo, Endpoints []string) {
	this.Actor.Init()
	this.clusterLocker = &sync.RWMutex{}
	this.clusterMap = make(map[uint32]*ClusterNode)
	this.master = NewMaster(info, Endpoints, &this.Actor)
	this.hashRing = tools.NewHashRing()
	this.clusterInfoMap = make(map[uint32]*common.ClusterInfo)
	this.packetFuncList = vector.NewVector()
	//集群新加member
	this.RegisterCall("Cluster_Add", func(ctx context.Context, info *common.ClusterInfo) {
		_, bEx := this.clusterInfoMap[info.Id()]
		if !bEx {
			this.AddCluster(info)
		}
	})
	//集群删除member
	this.RegisterCall("Cluster_Del", func(ctx context.Context, info *common.ClusterInfo) {
		delete(this.clusterInfoMap, info.Id())
		this.DelCluster(info)
	})

	//链接断开
	this.RegisterCall("DISCONNECT", func(ctx context.Context, ClusterId uint32) {
		pInfo, bEx := this.clusterInfoMap[ClusterId]
		if bEx {
			this.DelCluster(pInfo)
		}
		delete(this.clusterInfoMap, ClusterId)
	})
	this.Actor.Start()
}
func (this *Cluster) AddCluster(info *common.ClusterInfo) {
	pClient := new(network.ClientSocket)
	pClient.Init(info.Ip, int(info.Port))
	packet := reflect.New(reflect.ValueOf(this.packet).Elem().Type()).Interface().(IClusterPacket)
	packet.Init()
	packet.SetClusterId(info.Id())
	pClient.BindPacketFunc(packet.PacketFunc)
	for _, v := range this.packetFuncList.Values() {
		pClient.BindPacketFunc(v.(network.PacketFunc))
	}
	this.clusterLocker.Lock()
	this.clusterMap[info.Id()] = &ClusterNode{ClientSocket: pClient, ClusterInfo: info}
	this.clusterLocker.Unlock()
	this.hashRing.Add(info.IpString())
	pClient.Start()
}

func (this *Cluster) DelCluster(info *common.ClusterInfo) {
	this.clusterLocker.RLock()
	pCluster, bEx := this.clusterMap[info.Id()]
	this.clusterLocker.RUnlock()
	if bEx {
		pCluster.CallMsg("STOP_ACTOR")
		pCluster.Stop()
	}

	this.clusterLocker.Lock()
	delete(this.clusterMap, info.Id())
	this.clusterLocker.Unlock()
	this.hashRing.Remove(info.IpString())
}

func (this *Cluster) GetCluster(head rpc.RpcHead) *ClusterNode {
	this.clusterLocker.RLock()
	pCluster, bEx := this.clusterMap[head.ClusterId]
	this.clusterLocker.RUnlock()
	if bEx {
		return pCluster
	}
	return nil
}

func (this *Cluster) BindPacketFunc(callfunc network.PacketFunc) {
	this.packetFuncList.PushBack(callfunc)
}

func (this *Cluster) BindPacket(packet IClusterPacket) {
	this.packet = packet
}

func (this *Cluster) sendPoint(head rpc.RpcHead, buff []byte) {
	pCluster := this.GetCluster(head)
	if pCluster != nil {
		pCluster.Send(head, buff)
	}
}

func (this *Cluster) balanceSend(head rpc.RpcHead, buff []byte) {
	_, head.ClusterId = this.hashRing.Get64(head.Id)
	pClient := this.GetCluster(head)
	if pClient != nil {
		pClient.Send(head, buff)
	}
}

func (this *Cluster) boardCastSend(head rpc.RpcHead, buff []byte) {
	clusterList := []*ClusterNode{}
	this.clusterLocker.RLock()
	for _, v := range this.clusterMap {
		clusterList = append(clusterList, v)
	}
	this.clusterLocker.RUnlock()
	for _, v := range clusterList {
		v.Send(head, buff)
	}
}

func (this *Cluster) SendMsg(head rpc.RpcHead, funcName string, params ...interface{}) {
	buff := rpc.Marshal(head, funcName, params...)
	this.Send(head, buff)
}

func (this *Cluster) Send(head rpc.RpcHead, buff []byte) {
	switch head.SendType {
	case rpc.SEND_BALANCE:
		this.balanceSend(head, buff)
	case rpc.SEND_POINT:
		this.sendPoint(head, buff)
	default:
		this.boardCastSend(head, buff)
	}
}

func (this *Cluster) RandomCluster(head rpc.RpcHead) rpc.RpcHead {
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

package cluster

import (
	"fmt"
	"log"
	"strings"

	"github.com/erDong01/micro-kit/actor"
	"github.com/erDong01/micro-kit/cluster/etcdv3"
	"github.com/erDong01/micro-kit/common"
	"github.com/erDong01/micro-kit/rpc"
	"github.com/nats-io/nats.go"
)

type (
	Service    etcdv3.Service
	Master     etcdv3.Master
	Snowflake  etcdv3.Snowflake
	PlayerRaft etcdv3.PlayerRaft
)

// NewService 注册服务器
func NewService(info *common.ClusterInfo, Endpoints []string) *Service {
	service := &etcdv3.Service{}
	service.Init(info, Endpoints)
	return (*Service)(service)
}

// NewMaster 监控服务器
func NewMaster(info common.IClusterInfo, Endpoints []string, pActor actor.IActor) *Master {
	master := &etcdv3.Master{}
	master.Init(info, Endpoints, pActor)
	return (*Master)(master)
}

// NewSnowflake uuid生成器
func NewSnowflake(Endpoints []string) *Snowflake {
	uuid := &etcdv3.Snowflake{}
	uuid.Init(Endpoints)
	return (*Snowflake)(uuid)
}

func NewPlayerRaft(Endpoints []string) *PlayerRaft {
	playerRaft := &etcdv3.PlayerRaft{}
	playerRaft.Init(Endpoints)
	return (*PlayerRaft)(playerRaft)
}
func (this *PlayerRaft) GetPlayer(Id int64) *rpc.PlayerClusterInfo {
	return (*etcdv3.PlayerRaft)(this).GetPlayer(Id)
}

func (this *PlayerRaft) Publish(info *rpc.PlayerClusterInfo) bool {
	return (*etcdv3.PlayerRaft)(this).Publish(info)
}

func (this *PlayerRaft) Lease(leaseId int64) error {
	return (*etcdv3.PlayerRaft)(this).Lease(leaseId)
}

func GetRpcChannel(head rpc.RpcHead) string {
	return fmt.Sprintf("%s/%s/%d", etcdv3.ETCD_DIR, strings.ToLower(head.DestServerType.String()), head.ClusterId)
}

func GetRpcTopicChannel(head rpc.RpcHead) string {
	return fmt.Sprintf("%s/%s", etcdv3.ETCD_DIR, strings.ToLower(head.DestServerType.String()))
}
func GetRpcCallChannel(head rpc.RpcHead) string {
	return fmt.Sprintf("%s/%s/call/%d", etcdv3.ETCD_DIR, strings.ToLower(head.DestServerType.String()), head.ClusterId)
}

func GetChannel(clusterInfo common.ClusterInfo) string {
	return fmt.Sprintf("%s/%s/%d", etcdv3.ETCD_DIR, clusterInfo.String(), clusterInfo.Id())
}

func GetTopicChannel(clusterInfo common.ClusterInfo) string {
	return fmt.Sprintf("%s/%s", etcdv3.ETCD_DIR, clusterInfo.String())
}

func GetCallChannel(clusterInfo common.ClusterInfo) string {
	return fmt.Sprintf("%s/%s/call/%d", etcdv3.ETCD_DIR, clusterInfo.String(), clusterInfo.Id())
}

func SetupNatsConn(connectString string, appDieChan chan bool, options ...nats.Option) (*nats.Conn, error) {
	natsOptions := append(
		options,
		nats.DisconnectHandler(func(_ *nats.Conn) {
			log.Println("disconnected from nats!")
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("reconnected to nats server %s with address %s in cluster %s!", nc.ConnectedServerId(), nc.ConnectedAddr(), nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			err := nc.LastError()
			if err == nil {
				log.Println("nats connection closed with no error.")
				return
			}

			log.Fatalf("nats connection closed. reason: %q", nc.LastError())
			if appDieChan != nil {
				appDieChan <- true
			}
		}),
	)

	nc, err := nats.Connect(connectString, natsOptions...)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

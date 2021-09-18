package common

import (
	"fmt"
	"github.com/erDong01/micro-kit/rpc"
	"github.com/erDong01/micro-kit/tools"
	"strings"
)

type (
	ClusterInfo  rpc.ClusterInfo
	IClusterInfo interface {
		Id() uint32
		String() string
		ServiceType() rpc.SERVICE
		IpString() string
		RaftIp() string
	}
)

func (this ClusterInfo) IpString() string {
	return fmt.Sprintf("%s:%d", this.Ip, this.Port)
}

func (this *ClusterInfo) RaftIp() string {
	return fmt.Sprintf("%s:%d", this.Ip, this.Port+10000)
}

func (this *ClusterInfo) String() string {
	return strings.ToLower(this.Type.String())
}

func (this *ClusterInfo) Id() uint32 {
	return tools.ToHash(this.IpString())
}

func (this *ClusterInfo) ServiceType() rpc.SERVICE {
	return this.Type
}

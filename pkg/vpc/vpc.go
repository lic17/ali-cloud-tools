package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/lic17/ali-cloud-tools/pkg/client"
)

type Vpc struct {
	client *vpc.Client
}

func NewVpc(client *client.Client) *Vpc {

	v := &Vpc{}
	v.client = client.Vpc

	return v
}

package main

import (
	"fmt"

	ali_vpc "github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/lic17/ali-cloud-tools/pkg/client"
	"github.com/lic17/ali-cloud-tools/pkg/ecs"
	"github.com/lic17/ali-cloud-tools/pkg/vpc"
)

func main() {
	client := client.NewClient("cn-beijing")

	setTags(client)
}

func setTags(client *client.Client) {
	e := ecs.NewEcs(client)
	v := vpc.NewVpc(client)

	eips := v.GetAllUseEIP()
	fmt.Println("len eips:", len(eips))
	for _, d := range eips {
		instanceId := d.InstanceId
		eipId := d.AllocationId
		fmt.Println("instance id:", instanceId)
		fmt.Println("eip id:", eipId)

		tags := e.GetEcsTags(instanceId)
		var addTags []ali_vpc.TagResourcesTag

		for _, t := range tags.Tag {

			var tag ali_vpc.TagResourcesTag

			tag.Value = t.TagValue
			tag.Key = t.TagKey

			addTags = append(addTags, tag)
		}

		if len(addTags) > 0 {
			v.SetEIPTags(eipId, addTags)
		}
	}
}

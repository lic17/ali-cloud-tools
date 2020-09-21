package main

import (
	"fmt"
	"sync"

	ali_vpc "github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/lic17/ali-cloud-tools/pkg/client"
	"github.com/lic17/ali-cloud-tools/pkg/ecs"
	"github.com/lic17/ali-cloud-tools/pkg/vpc"
)

func main() {
	client := client.NewClient("cn-beijing")

	setEipsTags(client)
}

func setEipsTags(client *client.Client) {

	var wg sync.WaitGroup

	e := ecs.NewEcs(client)
	v := vpc.NewVpc(client)

	eips := v.GetAllUseEIP()
	fmt.Println("len eips:", len(eips))
	for _, eip := range eips {

		wg.Add(1)

		go func(eip ali_vpc.EipAddress) {

			defer wg.Done()

			instanceId := eip.InstanceId
			eipId := eip.AllocationId

			tags := e.GetEcsTags(instanceId)
			var addTags []ali_vpc.TagResourcesTag

			for _, t := range tags.Tag {

				var tag ali_vpc.TagResourcesTag

				tag.Value = t.TagValue
				tag.Key = t.TagKey

				addTags = append(addTags, tag)
			}

			fmt.Println("instance id: ", instanceId, "eip id: ", eipId, "get tags: ", eip.Tags, "add tags: ", addTags)

			if len(addTags) > 0 {
				v.SetEIPTags(eipId, addTags)
			}

		}(eip)
	}

	wg.Wait()

}

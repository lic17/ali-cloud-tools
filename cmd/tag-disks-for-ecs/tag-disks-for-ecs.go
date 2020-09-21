package main

import (
	"fmt"

	ali_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/lic17/ali-cloud-tools/pkg/client"
	"github.com/lic17/ali-cloud-tools/pkg/ecs"
)

func main() {
	client := client.NewClient("cn-beijing")
	e := ecs.NewEcs(client)

	setTags(e)
}

func setTags(e *ecs.Ecs) {

	disks := e.GetAllUseDisk()
	for _, d := range disks {
		instanceId := d.InstanceId
		diskId := d.DiskId
		fmt.Println("instance id:", instanceId)
		fmt.Println("disk id:", diskId)

		tags := e.GetEcsTags(instanceId)
		var addTags []ali_ecs.AddTagsTag

		for _, t := range tags.Tag {

			var tag ali_ecs.AddTagsTag

			tag.Value = t.TagValue
			tag.Key = t.TagKey

			addTags = append(addTags, tag)
		}

		if len(addTags) > 0 {
			e.SetDiskTags(diskId, addTags)
		}
	}

}

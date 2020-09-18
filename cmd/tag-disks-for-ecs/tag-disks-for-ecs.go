package main

import (
	ali_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/lic17/ali-cloud-tools/pkg/client"
	"github.com/lic17/ali-cloud-tools/pkg/ecs"
)

func main() {
	client := client.NewClient("cn-beijing")
	e := ecs.NewEcs(client)

	disks := e.GetAllUseDisk()
	for _, d := range disks.Disk {
		instanceId := d.InstanceId
		diskId := d.DiskId

		tags := e.GetEcsTags(instanceId)
		var addTags []ali_ecs.AddTagsTag
		for _, t := range tags.Tag {

			var tag ali_ecs.AddTagsTag

			tag.Value = t.TagValue
			tag.Key = t.TagKey

			addTags = append(addTags, tag)
		}

		e.SetDiskTags(diskId, addTags)
	}
}

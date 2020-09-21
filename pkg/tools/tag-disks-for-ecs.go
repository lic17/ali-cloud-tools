package tools

import (
	"fmt"
	"sync"

	ali_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/lic17/ali-cloud-tools/pkg/client"
	"github.com/lic17/ali-cloud-tools/pkg/ecs"
)

func SetDisksTags(client *client.Client) {

	var wg sync.WaitGroup

	e := ecs.NewEcs(client)

	disks := e.GetAllUseDisk()
	for _, d := range disks {
		wg.Add(1)

		go func(d ali_ecs.Disk) {

			defer wg.Done()

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

			fmt.Println("SET DISKS TAGS: instance id: ", instanceId, "disk id: ", diskId, "get tags: ", d.Tags, "add tags: ", addTags)
			if len(addTags) > 0 {
				e.SetDiskTags(diskId, addTags)
			}
		}(d)
	}

	wg.Wait()

}

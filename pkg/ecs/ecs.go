package ecs

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/lic17/ali-cloud-tools/pkg/client"
)

type Ecs struct {
	client *ecs.Client
}

func NewEcs(client *client.Client) *Ecs {

	e := &Ecs{}
	e.client = client.Ecs

	return e
}

func (e *Ecs) GetAllUseDisk() *ecs.DisksInDescribeDisks {
	request := ecs.CreateDescribeDisksRequest()
	request.Scheme = "https"

	request.Status = "In_use"

	disks, err := e.client.DescribeDisks(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	return &disks.Disks
}

func (e *Ecs) GetEcsTags(id string) *ecs.TagsInDescribeTags {
	request := ecs.CreateDescribeTagsRequest()
	request.Scheme = "https"

	request.ResourceType = "instance"
	request.ResourceId = id

	tags, err := e.client.DescribeTags(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	return &tags.Tags
}

func (e *Ecs) SetDiskTags(id string, tags []ecs.AddTagsTag) {
	request := ecs.CreateAddTagsRequest()
	request.Scheme = "https"

	request.ResourceType = "disk"
	request.ResourceId = id
	request.Tag = &tags

	response, err := e.client.AddTags(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

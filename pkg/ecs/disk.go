package ecs

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func (e *Ecs) GetOnePageUseDisk(nextToken string) (*ecs.DescribeDisksResponse, error) {
	request := ecs.CreateDescribeDisksRequest()
	request.Scheme = "https"

	request.Status = "In_use"
	request.NextToken = nextToken
	request.MaxResults = requests.NewInteger(100)

	disks, err := e.client.DescribeDisks(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	return disks, err
}

func (e *Ecs) GetAllUseDisk() []ecs.Disk {
	var disks []ecs.Disk
	nextToken := ""
	for {
		diskRep, err := e.GetOnePageUseDisk(nextToken)
		if err != nil {
			fmt.Print(err.Error())
			break
		}
		disks = append(disks, diskRep.Disks.Disk...)
		if diskRep.NextToken != "" {
			nextToken = diskRep.NextToken
			continue
		}
		break
	}

	return disks
}

func (e *Ecs) SetDiskTags(id string, tags []ecs.AddTagsTag) error {
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
	return err
}

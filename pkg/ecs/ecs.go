package ecs

import (
	"errors"
	"fmt"
	"strings"

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

func (e *Ecs) GetEcsTags(id string) *ecs.TagsInDescribeTags {
	var err error
	var ecsTags *ecs.TagsInDescribeTags

	eId := id
	idSplit := strings.Split(id, "-")

	if idSplit[0] == "eni" {
		//eId = id
		eId, err = e.GetEcsIdByEni(id)
		if err != nil {
			fmt.Print(err.Error())
			return ecsTags
		}
	}

	request := ecs.CreateDescribeTagsRequest()
	request.Scheme = "https"

	request.ResourceType = "instance"
	request.ResourceId = eId

	tags, err := e.client.DescribeTags(request)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		ecsTags = &tags.Tags
	}
	return ecsTags
}

func (e *Ecs) GetEcsIdByEni(eniId string) (string, error) {

	id := ""

	request := ecs.CreateDescribeNetworkInterfacesRequest()
	request.Scheme = "https"

	request.NetworkInterfaceId = &[]string{eniId}

	eni, err := e.client.DescribeNetworkInterfaces(request)
	if err != nil {
		fmt.Print(err.Error())
		return id, err
	}
	if len(eni.NetworkInterfaceSets.NetworkInterfaceSet) <= 0 {
		err = errors.New("can't find instance id by network instance!")
	} else {
		id = eni.NetworkInterfaceSets.NetworkInterfaceSet[0].InstanceId
	}

	return id, err
}

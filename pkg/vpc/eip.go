package vpc

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

var pageSize = 100

func (v *Vpc) GetOnePageUseEIP(pageNum int) (*vpc.DescribeEipAddressesResponse, error) {
	request := vpc.CreateDescribeEipAddressesRequest()
	request.Scheme = "https"

	request.Status = "InUse"
	request.PageNumber = requests.NewInteger(pageNum)
	request.PageSize = requests.NewInteger(pageSize)

	eips, err := v.client.DescribeEipAddresses(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	return eips, err
}

func (v *Vpc) GetAllUseEIP() []vpc.EipAddress {
	var eips []vpc.EipAddress
	nextPage := 1
	for {
		eipRep, err := v.GetOnePageUseEIP(nextPage)
		if err != nil {
			fmt.Print(err.Error())
			break
		}
		eipsOnePage := eipRep.EipAddresses.EipAddress
		eips = append(eips, eipsOnePage...)
		if len(eipsOnePage) >= pageSize {
			nextPage += 1
			continue
		}
		break
	}

	return eips
}

func (v *Vpc) SetEIPTags(id string, tags []vpc.TagResourcesTag) error {
	request := vpc.CreateTagResourcesRequest()

	request.Scheme = "https"

	var ids []string
	request.ResourceType = "EIP"
	ids = append(ids, id)
	request.ResourceId = &ids
	request.Tag = &tags

	response, err := v.client.TagResources(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return err
}

package vpc

import (
	"encoding/json"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/lic17/ali-cloud-tools/pkg/client"
	"github.com/lic17/ali-cloud-tools/pkg/ecs"
)

var clit = client.NewClient("cn-beijing")
var v = NewVpc(clit)
var e = ecs.NewEcs(clit)
var addTags []vpc.TagResourcesTag

func TestGetAllUseEIP(t *testing.T) {
	eips := v.GetAllUseEIP()
	eipsJson, err := json.Marshal(eips)
	if err != nil {
		t.Log("JSON ERR:", err)
	}
	t.Log(string(eipsJson))
}

func TestGetEcsTags(t *testing.T) {
	tags := e.GetEcsTags("i-2ze8bpfv23ibbnalpw95")

	for _, t := range tags.Tag {

		var tag vpc.TagResourcesTag

		tag.Value = t.TagValue
		tag.Key = t.TagKey

		addTags = append(addTags, tag)
	}
	tagsJson, err := json.Marshal(tags)
	if err != nil {
		t.Log("JSON ERR:", err)
	}
	t.Log(string(tagsJson))
}

func TestSetEIPTags(t *testing.T) {
	eipId := "eip-2zeyjlv5e8ry689ysrvhw"
	err := v.SetEIPTags(eipId, addTags)
	if err != nil {
		t.Log("set tags err:", err)
	}
}

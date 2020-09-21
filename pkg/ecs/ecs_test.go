package ecs

import (
	"encoding/json"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/lic17/ali-cloud-tools/pkg/client"
)

var clit = client.NewClient("cn-beijing")
var e = NewEcs(clit)
var addTags []ecs.AddTagsTag

func TestGetAllUseDisk(t *testing.T) {
	disks := e.GetAllUseDisk()
	disksJson, err := json.Marshal(disks)
	if err != nil {
		t.Log("JSON ERR:", err)
	}
	t.Log(string(disksJson))
}

func TestGetEcsTags(t *testing.T) {
	tags := e.GetEcsTags("i-2ze8tabph0ks9mrixtwy")

	for _, t := range tags.Tag {

		var tag ecs.AddTagsTag

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

func TestSetDiskTags(t *testing.T) {
	diskId := "d-2zeiacwfy9ed91lmkbpu"
	err := e.SetDiskTags(diskId, addTags)
	if err != nil {
		t.Log("set tags err:", err)
	}
}

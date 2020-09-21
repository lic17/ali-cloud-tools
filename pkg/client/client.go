package client

import (
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

type Client struct {
	Ecs *ecs.Client
	Vpc *vpc.Client
}

func NewClient(regionId string) *Client {

	keyId := os.Getenv("ACCESS_KEY_ID")
	keySecret := os.Getenv("ACCESS_KEY_SECRET")

	client := &Client{}
	var err error

	client.Ecs, err = ecs.NewClientWithAccessKey(regionId, keyId, keySecret)
	if err != nil {
		fmt.Print(err.Error())
	}

	client.Vpc, err = vpc.NewClientWithAccessKey(regionId, keyId, keySecret)
	if err != nil {
		fmt.Print(err.Error())
	}

	return client
}

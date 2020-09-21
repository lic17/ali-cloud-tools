package main

import (
	"github.com/lic17/ali-cloud-tools/pkg/client"
	"github.com/lic17/ali-cloud-tools/pkg/tools"
)

func main() {
	client := client.NewClient("cn-beijing")

	tools.SetDisksTags(client)
	tools.SetEipsTags(client)
}

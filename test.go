package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func cTest() api.KVPairs {
	config := api.DefaultConfig()
	config.Token = "testToken"
	client, _ := api.NewClient(config)
	kvps, _, _ := client.KV().List("/", nil)
	return kvps
}

func optTest() api.KVPairs {
	client, _ := api.NewClient(api.DefaultConfig())
	opts := &api.QueryOptions{
		Token: "testToken",
	}
	kvps, _, _ := client.KV().List("/", opts)
	return kvps
}

func main() {
	fmt.Printf("Client Config Test: %+v\n", cTest())
	fmt.Printf("Client QueryOptions Test: %+v\n", optTest())
}

package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/consul/api"
)

type KVPair struct {
	Key         string
	CreateIndex uint64
	ModifyIndex uint64
	LockIndex   uint64
	Flags       uint64
	Value       string
	Session     string
}

func main() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf("Error creating Consul client: %v", err)
	}

	prefix := "myapp/config/"
	var lastIndex uint64
	lastModifyIndex := make(map[string]uint64)

	for {
		kvPairs, meta, err := client.KV().List(prefix, &api.QueryOptions{
			WaitIndex: lastIndex,
			WaitTime:  10 * time.Second,
		})
		if err != nil {
			log.Printf("Error fetching keys: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		if meta.LastIndex != lastIndex {
			lastIndex = meta.LastIndex

			updatedKeys := []api.KVPair{}
			for _, kv := range kvPairs {
				if lastIndex, ok := lastModifyIndex[kv.Key]; !ok || kv.ModifyIndex > lastIndex {
					lastModifyIndex[kv.Key] = kv.ModifyIndex
					updatedKeys = append(updatedKeys, *kv)
				}
			}

			if len(updatedKeys) > 0 {
				logUpdates(updatedKeys)
			}
		}
	}
}

func logUpdates(updatedKeys []api.KVPair) {
	logFile, err := os.OpenFile("watch.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	for _, kv := range updatedKeys {
		decodedValue, _ := base64.StdEncoding.DecodeString(string(kv.Value))
		logEntry := fmt.Sprintf("Key: %s, Value: %s, ModifyIndex: %d\n", kv.Key, string(decodedValue), kv.ModifyIndex)
		if _, err := logFile.WriteString(logEntry); err != nil {
			log.Printf("Error writing to log file: %v", err)
		}
	}
}

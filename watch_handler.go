package main

import (
	"encoding/json"
	"log"
	"os"
)

// KVPair represents a Consul KV pair
type KVPair struct {
	Key         string
	CreateIndex uint64
	ModifyIndex uint64
	LockIndex   uint64
	Flags       uint64
	Value       []byte
	Session     string
}

func main() {
	// Set up logging to a file
	logFile, err := os.OpenFile("/consul/data/consul-watch.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)

	// Decode JSON from STDIN
	var kvPairs []KVPair
	decoder := json.NewDecoder(os.Stdin)
	if err := decoder.Decode(&kvPairs); err != nil {
		logger.Printf("Error decoding JSON: %v\n", err)
		return
	}

	// Log each changed key
	for _, kv := range kvPairs {
		logger.Printf("Key changed: %s, ModifyIndex: %d, Value: %s\n",
			kv.Key,
			kv.ModifyIndex,
			string(kv.Value))
	}
}

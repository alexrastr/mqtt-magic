package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("MQTT MAGIC")

	var env = loadEnv()
	mqttSubscribe(env, 12)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

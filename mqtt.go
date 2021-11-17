package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func mqttSubscribe(env *Env, timeout time.Duration) {
	//Callback
	msgRcvd := func(client mqtt.Client, message mqtt.Message) {
		fmt.Println("mqtt got message")
		wakeMac(env.MacAddr)
	}

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%v:%v", env.MqttHost, env.MqttPort))
	opts.SetUsername(env.MqttUser)
	opts.SetPassword(env.MqttPassword)
	opts.SetPingTimeout(timeout * time.Second)
	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(env.WakeTopic, 0, msgRcvd); token.Wait() && token.Error() != nil {
			panic(token.Error().Error())
		}
	}
	opts.SetKeepAlive(timeout * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(timeout * time.Second)
	opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		fmt.Printf("!!!!!! mqtt connection lost error: %s\n" + err.Error())
	})
	opts.SetReconnectingHandler(func(c mqtt.Client, options *mqtt.ClientOptions) {
		fmt.Println("...... mqtt reconnecting ......")
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error().Error())
	}
}

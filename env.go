package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
)

type Env struct {
	MacAddr      string
	MqttUser     string
	MqttPassword string
	MqttHost     string
	MqttPort     string
	WakeTopic    string
}

func loadEnv() *Env {
	e := &Env{
		MacAddr:      os.Getenv("MAC_ADDR"),
		MqttUser:     os.Getenv("MQTT_USER"),
		MqttPassword: os.Getenv("MQTT_PASSWORD"),
		MqttHost:     os.Getenv("MQTT_HOST"),
		MqttPort:     os.Getenv("MQTT_PORT"),
		WakeTopic:    os.Getenv("WAKE_TOPIC"),
	}

	fmt.Println("=== LOAD ENVIRONMENT VARIABLES ===")
	e.print()
	fmt.Println("==================================")

	return e
}

//print all variables from the structure
func (e *Env) print() {
	v := reflect.ValueOf(e).Elem()
	for i := 0; i < v.NumField(); i++ {
		varName := v.Type().Field(i).Name
		varValue := v.Field(i).Interface()
		//hide password
		matched, _ := regexp.Match("(Password|Token)$", []byte(varName))
		if matched && varValue != "" {
			varValue = "************"
		}
		fmt.Printf("%v %v %v\n", varName, "=", varValue)
	}
}

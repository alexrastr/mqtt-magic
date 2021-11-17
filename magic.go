package main

import (
	"fmt"
	"net"

	"github.com/sabhiram/go-wol/wol"
)

func wakeMac(mac string) {
	var err error

	bcastAddr := fmt.Sprintf("%s:%s", "255.255.255.255", "9")
	udpAddr, err := net.ResolveUDPAddr("udp", bcastAddr)

	var localAddr *net.UDPAddr
	if err != nil {
		panic(err.Error())
	}

	mp, err := wol.New(mac)
	if err != nil {
		panic(err.Error())
	}

	bs, err := mp.Marshal()
	if err != nil {
		panic(err.Error())
	}

	conn, err := net.DialUDP("udp", localAddr, udpAddr)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	n, err := conn.Write(bs)
	if err == nil && n != 102 {
		err = fmt.Errorf("magic packet sent was %d bytes (expected 102 bytes sent)", n)
		panic(err.Error())
	}
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Magic packet sent successfully to %s\n", mac)
}

package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	iface    = "wlp3s0"
	snaplen  = int32(1600)
	promisc  = false
	timeout  = pcap.BlockForever
	filter   = "tcp and port 80"
	devFound = false
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}
	for _, device := range devices {
		if device.Name == iface {
			devFound = true
		}
		if !devFound {
			log.Panicln("DEVICE NOT EXIST")
		}
		handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
		if err != nil {
			log.Panicln(err)
		}
		defer handle.Close()
		if err := handle.SetBPFFilter(filter); err != nil {
			log.Panicln(err)
		}
		source := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range source.Packets() {
			fmt.Println(packet)
		}
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {

	args := os.Args[1:]
	var device = ""

	devices, err := pcap.FindAllDevs()

	if err != nil {
		panic("Error during devices scan " + err.Error())
	}

	println("Searching for devices...")
	for _, dev := range devices {
		if dev.Name == args[0] {
			println("Default device has been found! " + dev.Name)
			device = dev.Name
		}
	}

	if device == "" {
		panic("Device not found")
	}

	println("Starting scan...")

	live, err := pcap.OpenLive(device, 1600, false, pcap.BlockForever)

	if err != nil {
		panic("Unable to start live scan" + err.Error())
	}

	defer live.Close()

	if err := live.SetBPFFilter("udp and port 53"); err != nil {
		panic("Error during BPF filter")
	}

	pkgSource := gopacket.NewPacketSource(live, live.LinkType())

	for pkg := range pkgSource.Packets() {
		fmt.Println(pkg)
	}

}

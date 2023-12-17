package main

import (
	"fmt"
	"os"

	"github.com/0x000b/gosnifp/internal/logger"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {

	args := os.Args[1:]
	var device = ""

	logger := logger.Logger{
		Path: ".",
	}

	devices, err := pcap.FindAllDevs()

	if err != nil {
		panic("Error during devices scan " + err.Error())
	}

	logger.InfoLog("Searching for devices...")

	for _, dev := range devices {
		if dev.Name == args[0] {
			logger.InfoLog("Default device has been found! " + dev.Name)
			device = dev.Name
		}
	}

	if device == "" {
		panic("Device not found")
	}

	logger.InfoLog("Starting scan...")

	live, err := pcap.OpenLive(device, 1600, false, pcap.BlockForever)

	if err != nil {
		panic("Unable to start live scan" + err.Error())
	}

	defer live.Close()

	if err := live.SetBPFFilter("udp and port 53"); err != nil {
		panic("Error during BPF filter" + err.Error())
	}

	pkgSource := gopacket.NewPacketSource(live, live.LinkType())

	for pkg := range pkgSource.Packets() {
		fmt.Println(pkg)
	}

}

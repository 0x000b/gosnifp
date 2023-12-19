package main

import (
	"os"
	"strconv"
	"time"

	"github.com/0x000b/gosnifp/internal/logger"
	"github.com/0x000b/gosnifp/internal/models"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	dns     layers.DNS
	eth     layers.Ethernet
	ipv4    layers.IPv4
	ipv6    layers.IPv6
	tcp     layers.TCP
	udp     layers.UDP
	payload gopacket.Payload
	srcAddr string
	dstAddr string
)

func main() {

	args := os.Args[1:]
	var device = ""

	logger := logger.Logger{
		Path: ".",
	}

	devices, err := pcap.FindAllDevs()

	if err != nil {
		logger.ErrorLog("Error during devices scan")
	}

	logger.InfoLog("Searching for devices")

	for _, dev := range devices {
		if dev.Name == args[0] {
			logger.InfoLog("Default device has been found! " + dev.Name)
			device = dev.Name
		}
	}

	if device == "" {
		logger.ErrorLog("Device not found")
	}

	live, err := pcap.OpenLive(device, 1600, false, pcap.BlockForever)

	if err != nil {
		logger.ErrorLog(err.Error())
	}

	defer live.Close()

	if err := live.SetBPFFilter("udp and port 53"); err != nil {
		panic("Error during BPF filter: " + err.Error())
	}

	logger.InfoLog("BPF filter set [udp and port 53]")

	decodeParser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ipv4, &ipv6, &tcp, &udp, &dns, &payload)

	decodedLayers := make([]gopacket.LayerType, 0, 10)

	for {
		data, _, _ := live.ReadPacketData()

		decodeParser.DecodeLayers(data, &decodedLayers)

		for _, typ := range decodedLayers {

			if typ == layers.LayerTypeDNS {

				timer := time.Now()
				timestamp := timer.Format(time.DateTime)

				dnsAnswers := int(dns.ANCount)

				for _, question := range dns.Questions {

					question := models.NewQuestion(string(question.Name), question.Type.String())

					dnsMessage := models.NewDns(timestamp, srcAddr, dstAddr, question, dns.ResponseCode.String())

					logger.DNSLog(dnsMessage.ToString())

					if dnsAnswers > 0 {
						for _, answer := range dns.Answers {
							response := models.NewAnswer(string(answer.Name), answer.IP.String(), strconv.FormatUint(uint64(answer.TTL), 10))
							dnsMessage.Answers = append(dnsMessage.Answers, response)

							logger.DNSLog(srcAddr + " -> " + dstAddr + " " + response.Name + " " + string(response.IpAddr) + " " + string(response.TTL))
						}
					}

				}

			} else if typ == layers.LayerTypeIPv4 {
				srcAddr = ipv4.SrcIP.String()
				dstAddr = ipv4.DstIP.String()
			} else if typ == layers.LayerTypeIPv6 {
				srcAddr = ipv6.SrcIP.String()
				dstAddr = ipv6.DstIP.String()
			}

		}
	}

}

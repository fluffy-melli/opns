package devices

import (
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket/pcap"
	stnet "github.com/shirou/gopsutil/net"
)

type Device struct {
	Pcap    pcap.Interface
	Net     net.Interface
	Io      stnet.IOCountersStat
	Name    string
	Address net.IP
}

func DevicesInfo() {
	fmt.Println("---------------")
	for i, device := range Devices() {
		fmt.Printf("Device: %d/%d\n", i, len(Devices())-1)
		fmt.Printf("Name: %s\n", device.Name)
		fmt.Printf("Address: %s\n", device.Address)
		fmt.Printf(" - Rx: %v\n", device.Io.BytesRecv)
		fmt.Printf(" - Tx: %v\n", device.Io.BytesSent)
		fmt.Println("---------------")
	}
}

func Devices() []Device {
	respond := []Device{}
	for _, device := range PAC() {
		netIOCounters, err := stnet.IOCounters(true)
		if err != nil {
			log.Fatal("Error getting network stats:", err)
		}
		interfaces, err := net.Interfaces()
		if err != nil {
			log.Fatal("Error getting network stats:", err)
		}
		for _, iface := range interfaces {
			for _, ioc := range netIOCounters {
				if iface.Name == ioc.Name {
					addrs, err := iface.Addrs()
					if err != nil {
						continue
					}
					for _, addr := range addrs {
						for _, address := range device.Addresses {
							if address.IP.To4() == nil {
								continue
							}
							if ioc.BytesRecv == 0 && ioc.BytesSent == 0 {
								continue
							}
							if ipAddr, _, err := net.ParseCIDR(addr.String()); err == nil && ipAddr.String() == address.IP.String() {
								respond = append(respond, Device{
									Pcap:    device,
									Net:     iface,
									Io:      ioc,
									Name:    iface.Name,
									Address: address.IP,
								})
							}
						}
					}
				}
			}
		}
	}

	if len(respond) == 0 {
		log.Fatal("No devices found.")
	}
	return respond
}

func PAC() []pcap.Interface {
	respond := []pcap.Interface{}
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal("Error finding devices:", err)
	}
	if len(devices) == 0 {
		log.Fatal("No devices found.")
	}
	for _, device := range devices {
		if device.Addresses != nil && len(device.Addresses) > 0 {
			for _, address := range device.Addresses {
				if address.IP.To4() != nil && !address.IP.IsLoopback() {
					respond = append(respond, device)
				}
			}
		}
	}

	if len(respond) == 0 {
		log.Fatal("No devices found.")
	}
	return respond
}

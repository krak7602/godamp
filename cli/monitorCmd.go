package cli

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
	"log"
)

var filter string

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor a interface activity",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		device := args[0]
		handle, err := pcap.OpenLive(device, 65535, true, pcap.BlockForever)
		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()

		err = handle.SetBPFFilter(filter)
		if err != nil {
			log.Fatal(err)
		}

		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			fmt.Println(packet)
		}
	},
}

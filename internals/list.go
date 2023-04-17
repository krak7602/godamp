package internals

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
	"log"
)

func List(cmd *cobra.Command, args []string) {
	ifaces, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}

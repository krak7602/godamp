package cli

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
	"log"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the available interfaces",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ifaces, err := pcap.FindAllDevs()
		if err != nil {
			log.Fatal(err)
		}
		for _, iface := range ifaces {
			fmt.Println(iface)
		}
	},
}

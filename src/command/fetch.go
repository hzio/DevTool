package command

import (
	"fmt"
	"global"
	"gopkg.in/urfave/cli.v2"
	"net"
)

// ===========================================================
//
//       Operation commands declarations
//
// ===========================================================
var (
	IpFetchCommand = &cli.Command{
		Name:        "ip",
		Usage:       "Fetch the local IP address",
		Description: "Fetch the local IP address",
		Category:    global.CategoryFetch,
		Action:      execIpFetchCommand,
	}
)

// ===========================================================
//
//       Operation commands implementations
//
// ===========================================================
func execIpFetchCommand(ctx *cli.Context) error {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

	return nil
}

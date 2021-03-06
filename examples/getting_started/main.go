package main

import (
	"flag"
	"strings"

	"github.com/cocher/utils/log"
	"github.com/cocher/crypto/ed25519"
	"github.com/cocher/network"
	"github.com/cocher/network/backoff"
	"github.com/cocher/network/discovery"
	"github.com/cocher/network/nat"
)

func main() {
	// glog defaults to logging to a file, override this flag to log to console for testing
	flag.Set("logtostderr", "true")

	// process other flags
	portFlag := flag.Int("port", 3000, "port to listen to")
	hostFlag := flag.String("host", "localhost", "host to listen to")
	protocolFlag := flag.String("protocol", "tcp", "protocol to use (kcp/tcp)")
	peersFlag := flag.String("peers", "", "peers to connect to")
	natFlag := flag.Bool("nat", false, "enable nat traversal")
	reconnectFlag := flag.Bool("reconnect", false, "enable reconnections")
	flag.Parse()

	port := uint16(*portFlag)
	host := *hostFlag
	protocol := *protocolFlag
	natEnabled := *natFlag
	reconnectEnabled := *reconnectFlag
	peers := strings.Split(*peersFlag, ",")

	keys := ed25519.RandomKeyPair()

	log.Infof("Private Key: %s", keys.PrivateKeyHex())
	log.Infof("Public Key: %s", keys.PublicKeyHex())

	builder := network.NewBuilder()
	builder.SetKeys(keys)
	builder.SetAddress(network.FormatAddress(protocol, host, port))

	// Register NAT traversal Component.
	if natEnabled {
		nat.RegisterComponent(builder)
	}

	// Register the reconnection Component
	if reconnectEnabled {
		builder.AddComponent(new(backoff.Component))
	}

	// Register peer discovery Component.
	builder.AddComponent(new(discovery.Component))

	net, err := builder.Build()
	if err != nil {
		log.Fatal(err)
		return
	}

	go net.Listen()

	if len(peers) > 0 {
		net.Bootstrap(peers...)
	}

	select {}

}

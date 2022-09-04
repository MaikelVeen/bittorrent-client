package main

import (
	"fmt"
	"log"

	"github.com/MaikelVeen/bittorrent-client/torrent"
)

func main() {
	tor, err := torrent.ReadFile("goliath.torrent")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tor.Info.Name)
}

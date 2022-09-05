package main

import (
	"fmt"
	"log"

	"github.com/MaikelVeen/bittorrent-client/torrent"
)

func main() {
	tor, err := torrent.Open("night_of_the_living_dead_archive.torrent")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tor.TrackerURL(torrent.RandomPeerID(), torrent.Port))
}

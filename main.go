package main

import (
	"log"

	"github.com/MaikelVeen/bittorrent-client/torrent"
)

func main() {
	tor, err := torrent.Open("night_of_the_living_dead_archive.torrent")
	if err != nil {
		log.Fatal(err)
	}

	announceURL, err := tor.TrackerURL(torrent.RandomPeerID(), torrent.Port)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := torrent.Announce(announceURL)
	if err != nil {
		log.Fatal(err)
	}

	peers, err := resp.PeerAddresses()
	if err != nil {
		log.Fatal(err)
	}

	log.Default().Print(peers[0])
}

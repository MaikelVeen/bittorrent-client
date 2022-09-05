package torrent

import "crypto/rand"

// RandomPeerID return a 20-byte slice that can
// be used as an peer id.
func RandomPeerID() [20]byte {
	var peerID [20]byte
	_, err := rand.Read(peerID[:])
	if err != nil {
		panic("error while generating random peer id")
	}

	return peerID
}

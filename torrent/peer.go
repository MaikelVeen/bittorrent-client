package torrent

import (
	"crypto/rand"
	"encoding/binary"
	"net"
)

// PeerBinarySize: peer IP addresses are made out of groups of six bytes.
const PeerBinarySize = 6

// Unmarshal parses peer IP addresses and ports from a binary buffer.
func UnmarshalPeers(p []byte) ([]net.TCPAddr, error) {
	/*if len(p)%PeerBinarySize != 0 {
		return nil, errors.New("malformed peers buffer")
	}*/

	peerLen := len(p) / PeerBinarySize
	peers := make([]net.TCPAddr, peerLen)

	for i := 0; i < peerLen; i++ {
		offset := i * PeerBinarySize
		peers[i].IP = net.IP(p[offset : offset+4])
		peers[i].Port = int(binary.BigEndian.Uint16(p[offset+4 : offset+6]))
	}

	return peers, nil
}

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

package torrent

import (
	"errors"
	"net"
	"net/http"

	"github.com/MaikelVeen/bittorrent-client/bencode"
)

// Announce announces our presence as peer to the tracker.
func Announce(url string) (*TrackerResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	trackerResp := &TrackerResponse{}
	if err := bencode.Decode(resp.Body, trackerResp); err != nil {
		return nil, err
	}

	return trackerResp, nil
}

type TrackerResponse struct {
	// FailureReason is a human-readable error message as to why the request failed.
	FailureReason string `json:"failure reason"`
	// Interval in seconds that the client should wait between sending regular requests to the tracker
	Interval int `json:"interval"`
	// Peers is a string consisting of multiples of 6 bytes. First 4 bytes are the IP address and last 2 bytes are the port number.
	// All in network (big endian) notation.
	Peers string `json:"peers"`
}

func (t *TrackerResponse) Error() string {
	return t.FailureReason
}

func (t *TrackerResponse) PeerAddresses() ([]net.TCPAddr, error) {
	if t.Peers == "" {
		return nil, errors.New("cannot unmarshal empty peers string")
	}

	return UnmarshalPeers([]byte(t.Peers))
}

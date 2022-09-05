package torrent

import (
	"crypto/sha1"
	"encoding/json"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/MaikelVeen/bittorrent-client/bencode"
)

// Port to listen on
const Port uint16 = 6881

// TorrentFile encapsulates the data fields that are present in
// a .torrent file.
type TorrentFile struct {
	// Announce is the URL of the tracker.
	Announce string `json:"announce"`
	// Comment is an optional field containing free-form textual comments of the torrent author.
	Comment string `json:"comment"`
	// CreatedAt is an optional field containing the creation time of the torrent, in standard UNIX epoch format.
	CreatedAt int `json:"creation date"`
	// CreatedBy is an optional field containing the name and version of the program used to create the .torrent. Optional.
	CreatedBy string `json:"created by"`
	// Encoding is an optional field containing the string encoding format used to generate the pieces part of the info dictionary in the .torrent metafile
	Encoding string      `json:"encoding"`
	Info     TorrentInfo `json:"info"`
}

// Open returns a new torrent instance.
func Open(name string) (*TorrentFile, error) {
	reader, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := bencode.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: move to unmarshal to bencode package.
	// Convert map to json string
	jsonStr, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}

	// Convert json string to struct
	var torrent TorrentFile
	if err := json.Unmarshal(jsonStr, &torrent); err != nil {
		return nil, err
	}

	return &torrent, nil
}

// InfoHash returns the SHA1 hash of the Info dictionary.
func (t *TorrentFile) InfoHash() ([20]byte, error) {
	data, err := json.Marshal(t.Info)
	if err != nil {
		return sha1.Sum([]byte("")), nil
	}

	return sha1.Sum(data), nil
}

// TrackerURL returns the URL of the tracker.
func (t *TorrentFile) TrackerURL(peerID [20]byte, port uint16) (string, error) {
	baseURL, err := url.Parse(t.Announce)
	if err != nil {
		return "", err
	}

	infoHash, err := t.InfoHash()
	if err != nil {
		return "", err
	}

	params := url.Values{
		"info_hash":  []string{string(infoHash[:])},
		"peer_id":    []string{string(peerID[:])},
		"port":       []string{strconv.Itoa(int(Port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(t.Info.Length)},
	}

	baseURL.RawQuery = params.Encode()
	return baseURL.String(), nil
}

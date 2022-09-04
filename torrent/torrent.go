package torrent

import (
	"encoding/json"
	"log"
	"os"

	"github.com/MaikelVeen/bittorrent-client/bencode"
)

// Torrent encapsulates the data fields that are present in
// a .torrent file.
type Torrent struct {
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

// ReadFile returns a new torrent instance.
func ReadFile(name string) (*Torrent, error) {
	reader, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := bencode.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Convert map to json string
	jsonStr, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}

	// Convert json string to struct
	var torrent Torrent
	if err := json.Unmarshal(jsonStr, &torrent); err != nil {
		return nil, err
	}

	return &torrent, nil
}

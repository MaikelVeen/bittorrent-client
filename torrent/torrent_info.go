package torrent

// TorrentInfo is a dictionary that describes the file(s) of the torrent.
type TorrentInfo struct {
	Length      int    `json:"length"`
	Name        string `json:"name"`
	PieceLength int    `json:"piece length"`
	Pieces      string `json:"pieces"`
}

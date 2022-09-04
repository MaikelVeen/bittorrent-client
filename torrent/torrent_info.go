package torrent

// TorrentInfo is a dictionary that describes the file(s) of the torrent.
type TorrentInfo struct {
	Length      int    `json:"length"`
	Name        string `json:"name"`
	Pieces      string `json:"pieces"`
	PieceLength int    `json:"piece length"`
}

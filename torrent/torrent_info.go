package torrent

// TorrentInfo is a dictionary that describes the file(s) of the torrent.
type TorrentInfo struct {
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
}

package torrent

type TrackerResponse struct {
	// FailureReason is a human-readable error message as to why the request failed.
	FailureReason string `json:"failure reason"`
	// Interval in seconds that the client should wait between sending regular requests to the tracker
	Interval int `json:"interval"`
	// Peers is a string consisting of multiples of 6 bytes. First 4 bytes are the IP address and last 2 bytes are the port number. All in network (big endian) notation.
	Peers string `json:"peers"`
}

func (t *TrackerResponse) Error() string {
	return t.FailureReason
}

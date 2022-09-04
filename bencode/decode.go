package bencode

import (
	"bufio"
	"io"
	"strconv"
)

const (
	Dictionary       byte = 'd'
	Integer          byte = 'i'
	IntegerEndMarker byte = 'e'
	List             byte = 'l'
)

// Decode parses the reader stream and returns
// the Go type representation of a bencode encoded file.
func Decode(r io.Reader) (interface{}, error) {
	return decode(bufio.NewReader(r))
}

func decode(r *bufio.Reader) (interface{}, error) {
	c, err := r.ReadByte()
	if err != nil {
		return nil, err
	}

	switch c {
	case Integer:
		integerSlice, err := r.ReadBytes(IntegerEndMarker)
		if err != nil {
			return nil, err
		}

		buf := integerSlice[:len(integerSlice)-1]
		integer, err := strconv.Atoi(string(buf))
		if err != nil {
			return nil, err
		}

		return integer, nil
	}

	return nil, nil
}

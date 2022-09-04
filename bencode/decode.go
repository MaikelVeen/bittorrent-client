package bencode

import (
	"bufio"
	"io"
	"strconv"
)

const (
	Dictionary            byte = 'd'
	Integer               byte = 'i'
	EndDelimiter          byte = 'e'
	List                  byte = 'l'
	StringLengthDelimiter byte = ':'
)

// Decode parses the reader stream and returns
// the Go type representation of a bencode encoded file.
func Decode(r io.Reader) (interface{}, error) {
	i, err := Unmarshal(bufio.NewReader(r))
	if err != nil {
		return nil, err
	}

	return i, nil
}

func Unmarshal(r *bufio.Reader) (interface{}, error) {
	c, err := r.ReadByte()
	if err != nil {
		return nil, err
	}

	switch c {
	case Integer:
		return unmarshalInt(r)
	// TODO: Dictionaries.
	// TODO: Lists.
	case List:
		return unmarshalList(r)
	default:
		return unmarshalString(r)
	}
}

func unmarshalInt(r *bufio.Reader) (int, error) {
	integerSlice, err := r.ReadBytes(EndDelimiter)
	if err != nil {
		return -1, err
	}

	buf := integerSlice[:len(integerSlice)-1]
	integer, err := strconv.Atoi(string(buf))
	if err != nil {
		return -1, err
	}

	return integer, nil
}

func unmarshalString(r *bufio.Reader) (string, error) {
	if err := r.UnreadByte(); err != nil {
		return "", err
	}

	lengthSlice, err := r.ReadBytes(StringLengthDelimiter)
	if err != nil {
		return "", err
	}

	buf := lengthSlice[:len(lengthSlice)-1]
	length, err := strconv.Atoi(string(buf))
	if err != nil {
		return "", err
	}

	stringBuf := make([]byte, length)
	for i := 0; i < length; i++ {
		byt, err := r.ReadByte()
		if err != nil {
			return "", err
		}
		stringBuf[i] = byt
	}

	return string(stringBuf), nil
}

func unmarshalList(r *bufio.Reader) ([]interface{}, error) {
	list := []interface{}{}

	for {
		b, err := r.Peek(1)
		if err != nil {
			return nil, err
		}

		if b[0] == EndDelimiter {
			_, err := r.ReadByte()
			return list, err
		}

		val, err := Unmarshal(r)
		if err != nil {
			return nil, err
		}

		list = append(list, val)
	}
}

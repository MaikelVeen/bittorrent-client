package bencode

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Parallel()

	type args struct {
		r io.Reader
	}

	tests := map[string]struct {
		args    args
		want    interface{}
		wantErr bool
	}{
		"decodes an single integer corectly": {
			args: args{
				r: strings.NewReader("i420e"),
			},
			want:    int(420),
			wantErr: false,
		},
		"returns an error if an single integer is malformed": {
			args: args{
				r: strings.NewReader("i4r20e"),
			},
			wantErr: true,
		},
		"decodes a single string correctly": {
			args: args{
				r: strings.NewReader("10:loremipsum"),
			},
			wantErr: false,
			want:    "loremipsum",
		},
		"returns an error if string length is incorrect": {
			args: args{
				r: strings.NewReader("5:qwe"),
			},
			wantErr: true,
		},
		"decodes an empty list correctly": {
			args: args{
				r: strings.NewReader("le"),
			},
			want:    []interface{}{},
			wantErr: false,
		},
		"decodes a list of strings correctly": {
			args: args{
				r: strings.NewReader("l3:one3:two5:threee"),
			},
			want: []interface{}{
				"one", "two", "three",
			},
			wantErr: false,
		},
		"decodes a mixed list correctly": {
			args: args{
				r: strings.NewReader("l3:one3:two5:threei10ee"),
			},
			want: []interface{}{
				"one", "two", "three", 10,
			},
			wantErr: false,
		},
		"decodes a nested list correctly": {
			args: args{
				r: strings.NewReader("l3:one3:two5:threel3:one3:two5:threeee"),
			},
			want: []interface{}{
				"one", "two", "three", []interface{}{"one", "two", "three"},
			},
			wantErr: false,
		},
		"decodes dictionaries correctly": {
			args: args{
				r: strings.NewReader("d6:lengthi351272960e4:name31:debian-10.2.0-amd64-netinst.iso11:piecelengthi262144ee"),
			},
			want: map[string]interface{}{
				"length":      351272960,
				"name":        "debian-10.2.0-amd64-netinst.iso",
				"piecelength": 262144,
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := Decode(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

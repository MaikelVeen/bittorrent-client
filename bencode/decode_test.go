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

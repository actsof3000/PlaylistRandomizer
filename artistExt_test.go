package main

import (
	"reflect"
	"testing"

	"github.com/zmb3/spotify"
)

func makeArtist(id string) *spotify.SimpleArtist {
	artist := spotify.SimpleArtist{ID: spotify.ID(id)}
	return &artist
}

func Test_unique(t *testing.T) {
	type args struct {
		artists []spotify.SimpleArtist
	}
	tests := []struct {
		name string
		args args
		want []spotify.ID
	}{
		// TODO: Add test cases.
		{"Test Empty", args{artists: []spotify.SimpleArtist{}}, []spotify.ID{}},
		{"Test Single", args{artists: []spotify.SimpleArtist{*makeArtist("123")}}, []spotify.ID{spotify.ID("123")}},
		{"Test Multiple", args{artists: []spotify.SimpleArtist{*makeArtist("123"), *makeArtist("456")}}, []spotify.ID{spotify.ID("123"), spotify.ID("456")}},
		{"Test Repeating", args{artists: []spotify.SimpleArtist{*makeArtist("123"), *makeArtist("456"), *makeArtist("456")}}, []spotify.ID{spotify.ID("123"), spotify.ID("456")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unique(tt.args.artists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"

	"github.com/zmb3/spotify"
)

func makeMockClient() *spotify.Client {

}

func Test_getTopTracksForArtists(t *testing.T) {
	type args struct {
		client    *spotify.Client
		artistIds []spotify.ID
	}
	tests := []struct {
		name string
		args args
		want *map[spotify.ID][]*spotify.SimpleTrack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTopTracksForArtists(tt.args.client, tt.args.artistIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTopTracksForArtists() = %v, want %v", got, tt.want)
			}
		})
	}
}

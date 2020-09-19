package main

import "github.com/zmb3/spotify"
import "strings"

func unique(artists []spotify.SimpleArtist) []spotify.ID {
	uniqArtistIDs := make([]spotify.ID, 0)
	for _, artist := range artists {
		if(!contains(uniqArtistIDs, artist.ID)) {
			uniqArtistIDs = append(uniqArtistIDs, artist.ID)
		}
	}
	return uniqArtistIDs
}

func contains(slice []spotify.ID, id spotify.ID) bool {
	for _, a := range slice {
		if a == id {
			return true
		}
	}
	return false
}

func addToPlaylist(playlist []spotify.ID, track *spotify.SimpleTrack) bool {
	return !strings.Contains(track.Name, "Remix") && !contains(playlist, track.ID)
}
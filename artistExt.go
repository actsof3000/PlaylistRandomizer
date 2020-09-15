package main

import "github.com/zmb3/spotify"

func unique(artists []spotify.SimpleArtist) []spotify.ID {
	uniqArtistIDs := make([]spotify.ID, 0)
	for _, artist := range artists {
		if(!contains(uniqArtistIDs, artist.ID)) {
			uniqArtistIDs = append(uniqArtistIDs, artist.ID)
		}
	}
	return uniqArtistIDs
}

func contains(slice []spotify.ID, artist spotify.ID) bool {
	for _, a := range slice {
		if a == artist {
			return true
		}
	}
	return false
}
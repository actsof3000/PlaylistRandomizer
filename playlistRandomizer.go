package main

import (
	"fmt"
	"log"
	"github.com/zmb3/spotify"
)

func randomizePlaylist(client *spotify.Client, playist, randomizedPlaylist spotify.SimplePlaylist)  {
	fields := "items(track(artists(id),name))"
	playlistTrackPage, err := client.GetPlaylistTracksOpt(playist.ID, nil, fields)

	if err != nil {
		log.Fatalln(err)
		return
	}

	var uniqArtistIDs []spotify.ID
	// Get Unique artists
	for _, playlistTrack := range playlistTrackPage.Tracks {
		track := playlistTrack.Track
		for  _, trackArtist := range track.Artists {
			if(!contains(uniqArtistIDs, trackArtist.ID)) {
				uniqArtistIDs = append(uniqArtistIDs, trackArtist.ID)
			}
		}
	}

	fmt.Println("Getting Artist top tracks")
	topTracksByArtistID, err := getTopTracksForArtists(client, uniqArtistIDs)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var tracksForRandPL []spotify.ID
	for _, topTracks := range *topTracksByArtistID {
		for _, track := range topTracks {
			if(addToPlaylist(tracksForRandPL, &track)) {
				tracksForRandPL = append(tracksForRandPL, track.ID)
			}
		}
	}

	fmt.Println("Adding", len(tracksForRandPL),  "tracks to", randomizedPlaylist.Name)
	if len(tracksForRandPL) > 100 {
		trackMax := 100
		trackOffset := 0
		for trackOffset < len(tracksForRandPL) {
			page := tracksForRandPL[trackOffset:trackMax]

			_, err = client.AddTracksToPlaylist(randomizedPlaylist.ID, page...)
			if err != nil {
				log.Fatalln(err)
			}

			trackOffset = trackMax
			trackMax += 100
			if trackMax > len(tracksForRandPL) {
				trackMax = len(tracksForRandPL)
			}
		}
	} else {
		_, err = client.AddTracksToPlaylist(randomizedPlaylist.ID, tracksForRandPL...)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// Gets the top tracks for each artist.
// @param client  The spotify client
// @param artistIds  The ids for the artists to retrieve top tracks from
// @return a map of track arrays by their artistId
func getTopTracksForArtists(client *spotify.Client, artistIds []spotify.ID) (*map[spotify.ID][]spotify.SimpleTrack, error) {
	m := make(map[spotify.ID][]spotify.SimpleTrack)

	for _, artistID := range artistIds {
		topTracks, err := client.GetArtistsTopTracks(artistID, "US")

		if err != nil {
			return nil, err
		}

		var simpleTopTracks []spotify.SimpleTrack
		for _, track := range topTracks {
			simpleTopTracks = append(simpleTopTracks, track.SimpleTrack)
		}
		m[artistID] = simpleTopTracks
	}
	return &m, nil
}
package main

import (
      "bufio"
      "os"
      "os/exec"
	"log"
      "fmt"
      "strings"
      "runtime"
      "strconv"
	"github.com/zmb3/spotify"
      "net/http"
      uuid "github.com/google/uuid"
)

var state = "PR-" + uuid.New().String()
var auth spotify.Authenticator
var clientChan = make(chan *spotify.Client)

func main() {
      scope := []string{spotify.ScopePlaylistReadPrivate, spotify.ScopePlaylistModifyPrivate}
      clientID := "adf52684ca864e35842a9ce2b27939fb"
      secretKey := "3efa7cc5cf7949f08b94cf9107bd26b0"
      redirectURL := "http://localhost:8080/callback"

      http.HandleFunc("/callback", redirectHandler)

      go http.ListenAndServe(":8080", nil)

      authURL := getAuthURL(clientID, secretKey, redirectURL, scope...)
      fmt.Println("Opening:", authURL)
      openbrowser(authURL)

      client := <-clientChan

      user, err := client.CurrentUser()
      if(err != nil) {
            log.Fatalln(err)
      }

      fmt.Println("Logged in as:", user.ID)

      reader := bufio.NewReader(os.Stdin)
      simplePlaylistPage, err := client.GetPlaylistsForUser(user.ID)
      for i, playlist := range simplePlaylistPage.Playlists {
            fmt.Println(i + 1, ". ", playlist.Name)
      }
      fmt.Print("-> ")
      text, _ := reader.ReadString('\n')

      choice, err := strconv.Atoi(strings.Replace(text, "\n", "", -1))
      if(err != nil) {
            log.Fatalln(err)
      }

      playlistToRandomize := simplePlaylistPage.Playlists[choice - 1]
      fmt.Println("Randomizing:", playlistToRandomize.Name)

      outputPlaylist, err := client.CreatePlaylistForUser(user.ID, fmt.Sprintf("Randomizing %s", playlistToRandomize.Name), fmt.Sprintf("Randomized playlist made from each artist from %s", playlistToRandomize.Name), false)
      if(err != nil) {
            log.Fatalln(err)
      }
      randomizePlaylist(client, playlistToRandomize, outputPlaylist.SimplePlaylist)
}


func getAuthURL (clientID, secretKey, redirectURL string, scope ...string) string {
	auth = spotify.NewAuthenticator(redirectURL, scope...)
	auth.SetAuthInfo(clientID, secretKey)

	return auth.AuthURL(state)
}

// the user will eventually be redirected back to your redirect URL
// typically you'll have a handler set up like the following:
func redirectHandler(w http.ResponseWriter, r *http.Request) {
      // use the same state string here that you used to generate the URL
      token, err := auth.Token(state, r)
      if err != nil {
            http.Error(w, "Couldn't get token", http.StatusNotFound)
            return
      }
      // create a client using the specified token
      client := auth.NewClient(token)

      fmt.Println("Login complete!")

      clientChan <- &client

      // the client can now be used to make authenticated requests
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
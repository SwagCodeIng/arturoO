package main

import (
      "fmt"
      "github.com/julienschmidt/httprouter"
      "net/http"
      "log"
)

func getSongByID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      fmt.Fprint(w, "songID, songTitle, songArtist, songAlbum, songYear")
}

func main() {
  router := httprouter.New()
  router.GET("/getSong", getSongByID)
  log.Fatal(http.ListenAndServe(":8080", router))
}

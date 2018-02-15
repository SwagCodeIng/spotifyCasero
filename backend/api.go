package main

import (
      "fmt"
      "github.com/julienschmidt/httprouter"
      "net/http"
      "log"
      "encoding/json"
)

type Song struct {
    ID string
    Title string
    Artist string
    Album string
    Year int
    ImageURI string
}
type Artist struct {
    ID string
    Name string
}

func getSongByID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      song1 := &Song{ID: "1", Title: "Wasted Years", Artist: "Iron Maiden", Album: "Caught Somewhere In Time", Year: 1988, ImageURI: "https://images-na.ssl-images-amazon.com/images/I/61jEIt6vjUL.jpg"}
      data, err := json.Marshal(song1)
      if err != nil {
          fmt.Println(err)
          return
      }
      fmt.Fprint(w, string(data))
}

func getArtists(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      var artists = []Artist {
        Artist {
          ID: "1", Name: "Iron Maiden",
        },
        Artist {
          ID: "2", Name: "Post Malone",
        },
        Artist {
          ID: "3", Name: "Kendrick Lamar",
        },
        Artist {
          ID: "4", Name: "Wiz Khalifa",
        },
      }
      data, err := json.Marshal(artists)
      if err != nil {
          fmt.Println(err)
          return
      }
      fmt.Fprint(w, string(data))
}

func main() {

  router := httprouter.New()
  router.GET("/getSong", getSongByID)
  router.GET("/getArtists", getArtists)
  log.Fatal(http.ListenAndServe(":3011", router))

}

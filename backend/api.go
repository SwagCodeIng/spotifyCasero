package main

import (
      "fmt"
      "github.com/julienschmidt/httprouter"
      "net/http"
      "log"
      // "encoding/json"
      "gopkg.in/mgo.v2"
      "gopkg.in/mgo.v2/bson"
)

type Song struct {
    Id bson.ObjectId `json:"id" bson:"_id,omitempty"` //checar para las demas variables and the reason of mongo collection fields naming! ------
    InitTime string
    Name string
    Artist string
    Album string
    Year string
}
type Artist struct {
    Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Name string
}
type Person struct {
        Name string
        Phone string
}

// func getArtists(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//       var artists = []Artist {
//         Artist {
//           ID: "1", Name: "Iron Maiden",
//         },
//         Artist {
//           ID: "2", Name: "Post Malone",
//         },
//         Artist {
//           ID: "3", Name: "Kendrick Lamar",
//         },
//         Artist {
//           ID: "4", Name: "Wiz Khalifa",
//         },
//       }
//       data, err := json.Marshal(artists)
//       if err != nil {
//           fmt.Println(err)
//           return
//       }
//       fmt.Fprint(w, string(data))
// }

func main() {

  fmt.Println("Running server =3")

  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  c := session.DB("SpotiCloneTest").C("Songs")
  fmt.Println("Collection: ",c)
  err = c.Insert(&Song{"", "0.0", "Wasted Years", "Iron Maiden", "Caught Somewhere In Time", "1982"},
	               &Song{"", "0.0", "The Number Of The Beast", "Iron Maiden", "The Number Of The Beast", "1986"})
  if err != nil { //check why != nil and not == si se supone que tiene algo -----
    log.Fatal(err)
  }

  result := Song{}
        err = c.Find(bson.M{"name": "Wasted Years"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("ID:", result.Id.Hex(), "InitTime:", result.InitTime, "Name:", result.Name, "Artist:",
        result.Artist, "Album:", result.Album, "Year:", result.Year)

    c1 := session.DB("SpotiCloneTest").C("Artists")
    fmt.Println("Collection: ",c1)
    err = c1.Insert(&Artist{"", "Iron Maiden"},&Artist{"", "Post Malone"},&Artist{"", "Wiz Khalifa"},
    &Artist{"", "Lil Wayne"},&Artist{"", "Kendrick Lamar"})

  result1 := Artist{}
        err = c1.Find(bson.M{"name": "Post Malone"}).One(&result1)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("ID:", result1.Id.Hex(), "Name:", result1.Name)


  router := httprouter.New()
  // router.GET("/getSong", getSongByID)
  // router.GET("/getArtists", getArtists)
  log.Fatal(http.ListenAndServe(":8080", router))


}

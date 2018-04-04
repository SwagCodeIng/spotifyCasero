//Asi o separar los .go en varios (?)
package main

import (
      "fmt"
      "github.com/julienschmidt/httprouter"
      "net/http"
      "log"
      "encoding/json"
      "github.com/globalsign/mgo"
      "github.com/globalsign/mgo/bson"
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
type Album struct {
    Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Name string
    Artist string
    Genre string
    Year string
}

func getSong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      fmt.Println("Invocacion de getSong")
      ///Establecemos conexion con la db en mongo
      session, err := mgo.Dial("localhost:27017")
      if err != nil {
        panic(err)
      }
      defer session.Close()

      c := session.DB("SpotiCloneTest").C("Songs")
      test := "was"
      var songs []Song
          // err = c.Find(bson.M{"name": bson.RegEx{"was", "i"}}).All(&songs)
          err = c.Find(bson.M{"name": bson.RegEx{test, "i"}}).All(&songs)

            if err != nil {
                    log.Fatal(err)
            }
      data, err := json.Marshal(songs)
      if err != nil {
          fmt.Println(err)
          return
      }
      fmt.Fprint(w, string(data))
}

func getAllSongs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      fmt.Println("Invocacion de getAllSong")
      session, err := mgo.Dial("localhost:27017")
      if err != nil {
        panic(err)
      }
      defer session.Close()

      c := session.DB("SpotiCloneTest").C("Songs")
      var songs []Song
        err = c.Find(nil).All(&songs)

          if err != nil {
                  log.Fatal(err)
                }
                data, err := json.Marshal(songs)
                if err != nil {
                  fmt.Println(err)
                  return
                }
                fmt.Fprint(w, string(data))

  }

func getArtist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      fmt.Println("Invocacion de getArtist")
      ///Establecemos conexion con la db en mongo
      session, err := mgo.Dial("localhost:27017")
      if err != nil {
        panic(err)
      }
      defer session.Close()

      c := session.DB("SpotiCloneTest").C("Artists")
      var artists []Artist
            err = c.Find(bson.M{"name": bson.RegEx{"Post Malone", "i"}}).All(&artists) //Hardcodeo para la query de cierto artista
            if err != nil {
                    log.Fatal(err)
            }
      data, err := json.Marshal(artists)
      if err != nil {
          fmt.Println(err)
          return
      }
      fmt.Fprint(w, string(data))
}

func getAllArtists(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      fmt.Println("Invocacion de getAllArtists")
      session, err := mgo.Dial("localhost:27017")
      if err != nil {
        panic(err)
      }
      defer session.Close()

      c := session.DB("SpotiCloneTest").C("Artists")
      var artists []Artist
        err = c.Find(nil).All(&artists)

          if err != nil {
                  log.Fatal(err)
                }
                data, err := json.Marshal(artists)
                if err != nil {
                  fmt.Println(err)
                  return
                }
                fmt.Fprint(w, string(data))

  }

func getAlbum(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      fmt.Println("Invocacion de getAlbum")
      ///Establecemos conexion con la db en mongo
      session, err := mgo.Dial("localhost:27017")
      if err != nil {
        panic(err)
      }
      defer session.Close()

      c := session.DB("SpotiCloneTest").C("Albums")
      var albums []Album
            err = c.Find(bson.M{"name": bson.RegEx{"ston", "i"}}).All(&albums)
            if err != nil {
                    log.Fatal(err)
            }

      data, err := json.Marshal(albums)
      if err != nil {
          fmt.Println(err)
          return
      }
      fmt.Fprint(w, string(data))
}

func getAllAlbums(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      fmt.Println("Invocacion de getAllAlbums")
      session, err := mgo.Dial("localhost:27017")
      if err != nil {
        panic(err)
      }
      defer session.Close()

      c := session.DB("SpotiCloneTest").C("Albums")
      var albums []Album
        err = c.Find(nil).All(&albums)

          if err != nil {
                  log.Fatal(err)
                }
                data, err := json.Marshal(albums)
                if err != nil {
                  fmt.Println(err)
                  return
                }
                fmt.Fprint(w, string(data))
  }


func main() {

  fmt.Println("Running server =3")

  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  //El codigo que se encuentra dentro de las barras generadas con '/' pegados juntos (como la que est abajo
  //de este comentario es codigo completamente TEMPORAL y de PRUEBA mientras Mayo desarrolla la BD real)
  ////////////////////////////////////////////////////////////////////////////////////////////

  c := session.DB("SpotiCloneTest").C("Songs")
  fmt.Println("Collection: ",c)
  err = c.Insert(&Song{"", "0.0", "Wasted Years", "Iron Maiden", "Caught Somewhere In Time", "1982"},
	               &Song{"", "0.0", "The Number Of The Beast", "Iron Maiden", "The Number Of The Beast", "1986"},
                 &Song{"", "0.0", "Wasting Love", "Iron Maiden", "Fear Of The Dark", "1990"})
  if err != nil { //check why != nil and not == si se supone que tiene algo -----
    log.Fatal(err)
  }
  result := Song{}
  var songs []Song
        err = c.Find(bson.M{"name": bson.RegEx{"was", "i"}}).All(&songs)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("------------------------------------------------------debug1: ")
        fmt.Println(result)
        fmt.Println("ID:", result.Id.Hex(), "InitTime:", result.InitTime, "Name:", result.Name, "Artist:",
        result.Artist, "Album:", result.Album, "Year:", result.Year)
        fmt.Printf("finded songs %v\n")
        for _, song := range songs {
          fmt.Println("ID:", song.Id.Hex(), "InitTime:", song.InitTime, "Name:", song.Name, "Artist:",
          song.Artist, "Album:", song.Album, "Year:", song.Year)
        }
        fmt.Println("------------------------------------------------------debug1: ")


    //-----------------------------------------------------------------------------------------------------------

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

   //-----------------------------------------------------------------------------------------------------------

   c2 := session.DB("SpotiCloneTest").C("Albums")
   fmt.Println("Collection: ",c2)
   err = c2.Insert(&Album{"", "The Number Of The Beast", "Iron Maiden", "Heavy Metal", "1800"},
   &Album{"", "Stoney", "Post Malone", "Hip Hop", "2016"},
   &Album{"", "Caught Somewhere In Time", "Iron Maide", "Heavy Metal", "1985"})
   result2 := Album{}
       err = c2.Find(bson.M{"name": "Caught Somewhere In Time"}).One(&result2)
       if err != nil {
               log.Fatal(err)
       }
       fmt.Println("ID:", result2.Id.Hex(), "Name:", result2.Name, "Genre:", result2.Genre,
     "Year:", result2.Year)

   ///////////////////////////////////////////////////////////////////////////////////////////
     fmt.Println("Data inserted to db for testing purposes: true =3")


  router := httprouter.New()
  router.GET("/getSong", getSong)
  router.GET("/getArtist", getArtist)
  router.GET("/getAlbum", getAlbum)
  router.GET("/getAllSongs", getAllSongs)
  router.GET("/getAllArtists", getAllArtists)
  router.GET("/getAllAlbums", getAllAlbums)

  log.Fatal(http.ListenAndServe(":8080", router))


}

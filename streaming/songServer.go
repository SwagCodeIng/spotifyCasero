package main

import (
      "fmt"
      "github.com/julienschmidt/httprouter"
      "net/http"
      "log"
      "io"
      "github.com/mewkiz/flac"
      "encoding/json"
)

type songHeader struct {
  Name string
  Frames int

}

type songResponse struct {
  Frame int
  SubFrame int
  Frequencies []int32
}


func readSong(name string){
  songURI := "flacFiles/" + name + ".flac"

  stream, err := flac.Open(songURI)
  if err != nil{
    log.Fatal(err)
  }
  defer stream.Close()


  for{
    frame, err := stream.ParseNext()
    if err != nil{
      if err == io.EOF {
        break
      }//end if
      log.Fatal(err)
    }//end if
    if frame.Num <1 {
      fmt.Printf("frame %d\n", frame.Num)
      for i, subframe := range frame.Subframes{
      sResponse := &songResponse{
          Frame: int(frame.Num),
          SubFrame: i,
          Frequencies: subframe.Samples}
      fResponse, _  := json.Marshal(sResponse)
      fmt.Print(string(fResponse))
      }//end for
    }//end if
  }//end for
}//end readSong



func returnSong(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  test := readSong(string(p.ByName("name")))
  readSong(string(p.ByName("name")))
  data, err := json.Marshal(test)
  if err != nil {
      fmt.Println(err)
      return
  }
  fmt.Fprint(w, string(data))

  // response := "this is oging to be a json with the song frequencias"
  // data, err := json.Marshal(response)
  // if err != nil {
  //     fmt.Println(err)
  //     return
  // }
  // fmt.Fprint(w, p.ByName("name"))
  // fmt.Fprint(w, string(data))
}

func main() {

  router := httprouter.New()
  router.GET("/gimmeSong/:name", returnSong)

  log.Fatal(http.ListenAndServe(":8000", router))

}

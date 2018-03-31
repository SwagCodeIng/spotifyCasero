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
  Id int
}

type songResponse struct {
  Frame int
  SubFrame int
  Frequencies []int32
}

func getHeader(name string) string{
  songURI := "flacFiles/" + name + ".flac"
  stream, err := flac.Open(songURI)
  if err != nil{
    log.Fatal(err)
  }
  defer stream.Close()

  frameCount := 0
  for{
    _, err := stream.ParseNext()
    if err != nil{
      if err == io.EOF {
        break
      }//end if
      log.Fatal(err)
    }
    frameCount = frameCount + 1
  }//end for

  jsonResponse := &songHeader{
    Name: string(name),
    Frames: int(frameCount),
    Id: 1}
  data, err := json.Marshal(jsonResponse)
  if err != nil {
      fmt.Println(err)
      return "error"
  }
  return string(data)
}//end getHeader()

func returnSong(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  songURI := "flacFiles/" + p.ByName("name") + ".flac"

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
    // if frame.Num <2 {
      fmt.Printf("frame %d\n", frame.Num)
      for i, subframe := range frame.Subframes{
      sResponse := &songResponse{
          Frame: int(frame.Num),
          SubFrame: i,
          Frequencies: subframe.Samples}
      // fResponse, _  := json.Marshal(sResponse)
      data, err := json.Marshal(sResponse)
      if err != nil {
          fmt.Println(err)
          return
      }
      fmt.Fprint(w, string(data))
      }//end for
    // }//end if
  }//end for
}//end readSong()

func identify(w http.ResponseWriter, r *http.Request, p httprouter.Params){

  header := getHeader(p.ByName("name"))
  data, err := json.Marshal(header)
  if err != nil {
      fmt.Println(err)
      return
  }
  fmt.Fprint(w, string(data))
}//end identify()

func main() {
  router := httprouter.New()
  router.GET("/get-song/:name", returnSong)
  router.GET("/start-song/:name/:id", identify)

  log.Fatal(http.ListenAndServe(":8000", router))
}


// func readSong(name string) string{
//   songURI := "flacFiles/" + name + ".flac"
//
//   stream, err := flac.Open(songURI)
//   if err != nil{
//     log.Fatal(err)
//   }
//   defer stream.Close()
//
//
//   for{
//     frame, err := stream.ParseNext()
//     if err != nil{
//       if err == io.EOF {
//         break
//       }//end if
//       log.Fatal(err)
//     }//end if
//     if frame.Num <2 {
//       fmt.Printf("frame %d\n", frame.Num)
//       for i, subframe := range frame.Subframes{
//       sResponse := &songResponse{
//           Frame: int(frame.Num),
//           SubFrame: i,
//           Frequencies: subframe.Samples}
//       fResponse, _  := json.Marshal(sResponse)
//       fmt.Print(string(fResponse))
//       return string(fResponse)
//       }//end for
//     }//end if
//   }//end for
//   return ""
// }//end readSong

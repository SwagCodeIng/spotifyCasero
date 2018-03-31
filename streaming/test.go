package main

import (
  // "bytes"
  // "crypto/md5"
  "fmt"
  // "bufio"
  "log"
  "io"
  "github.com/mewkiz/flac"
  "encoding/json"
)

type songResponse struct {
  Frame int
  SubFrame int
  Frequencies []int32
}

func main(){
  stream, err := flac.Open("flacFiles/stairway-to-heaven.flac")
  if err != nil{
    log.Fatal(err)
  }
  defer stream.Close()



  // md5sum := md5.New()


    for{
      frame, err := stream.ParseNext()
      if err != nil{
        if err == io.EOF {
          break
        }//end if
        log.Fatal(err)
      }//end if
      // if frame.Num <1 {
        fmt.Printf("frame %d\n", frame.Num)
        for i, subframe := range frame.Subframes{
        sResponse := &songResponse{
            Frame: int(frame.Num),
            SubFrame: i,
            Frequencies: subframe.Samples}
        fResponse, _  := json.Marshal(sResponse)
        fmt.Print(string(fResponse))
        }//end for
      //}//end if
    }//end for
  fmt.Println()



//
// got, want := md5sum.Sum(nil), stream.Info.MD5sum[:]
// fmt.Println("decoded audio md5sum valid:", bytes.Equal(got, want))

}

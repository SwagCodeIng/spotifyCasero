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

type songFrame struct {
  Frame int
  LeftAudioChannel []int32
  RightAudioChannel []int32
}

func main(){
  var lac []int32
  var rac []int32
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
      if frame.Num <1 {
        fmt.Printf("frame %d\n", frame.Num)

        for i, subframe := range frame.Subframes{
          if i == 0{
            lac = subframe.Samples
          }
          if i == 1 {
            rac = subframe.Samples
          }
        }//end for

        frameResponse := &songFrame{
            Frame: int(frame.Num),
            LeftAudioChannel: lac,
            RightAudioChannel: rac}
        fResponse, _  := json.Marshal(frameResponse)
        fmt.Print(string(fResponse))

      }//end if
    }//end for
  fmt.Println()



//
// got, want := md5sum.Sum(nil), stream.Info.MD5sum[:]
// fmt.Println("decoded audio md5sum valid:", bytes.Equal(got, want))

}

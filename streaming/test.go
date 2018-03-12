package main

import (
  // "bytes"
  // "crypto/md5"
  "fmt"
  // "bufio"
  "log"
  "io"
  "github.com/mewkiz/flac"
)

func main(){
  stream, err := flac.Open("flacFiles/stairway-to-heaven.flac")
  if err != nil{
    log.Fatal(err)
  }
  defer stream.Close()

  // fmt.Print(stream.Info)
  //
  // var b bytes.Buffer
  // writer := bufio.NewWriter(&b)
  // fmt.Print(writer)


  // flac.Encode(writer, stream)

  // fmt.Print(writer)
  // fmt.Print(stream)


  // md5sum := md5.New()
  for{
    frame, err := stream.ParseNext()
    if err != nil{
      if err == io.EOF {
        break
      }//end if
      log.Fatal(err)
    }//end if
    // frame.Hash(md5sum)

    // fmt.Print(frame)

    if frame.Num <3 {
      fmt.Printf("frame %d\n", frame.Num)
      for i, subframe := range frame.Subframes{
        fmt.Printf(" subframe %d\n", i)
        fmt.Println("////////////////////////////////////")
        fmt.Println(subframe.NSamples)
        for j, sample := range subframe.Samples {
          if j >= 50{
            break
          }//end if
          fmt.Printf("   sample %d: %v\n", j, sample)

        }//end for
      }//end for
    }//end if
  }//end for
  fmt.Println()



//
// got, want := md5sum.Sum(nil), stream.Info.MD5sum[:]
// fmt.Println("decoded audio md5sum valid:", bytes.Equal(got, want))

}

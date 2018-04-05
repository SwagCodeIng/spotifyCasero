// Mauricio Figueroa Pereda 151504
// Manuel Escobar Hernandez 149688
package main

import (
      "fmt"
      "github.com/julienschmidt/httprouter"
      "net/http"
      "log"
      "encoding/json"
)


func showSomething(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      hello := "hello, server running and ready to be attacked :)"
      data, err := json.Marshal(hello)
      if err != nil {
          fmt.Println(err)
          return
      }
      fmt.Fprint(w, string(data))
}



func main() {

  router := httprouter.New()
  router.GET("/", showSomething)

  log.Fatal(http.ListenAndServe(":8000", router))

}

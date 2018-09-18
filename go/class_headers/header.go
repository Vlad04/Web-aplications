package main

import (
   "fmt"
   "net/http"
)

func main() {
  
   url := "http://www.lemonde.fr"

   respuesta, err := http.Head(url)
   if err != nil {
      fmt.Printf("Error URL. ", err)
   }

   for key, value := range respuesta.Header {
      fmt.Printf("%s: %s\n", key, value[0])
   }
} 
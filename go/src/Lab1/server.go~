package main
import 
(
  "fmt"
  "log"
  "net/http"
)

func darMensaje(w http.ResponseWriter, r *http.Request) {
  
  /*w.WriteHeader(http.StatusNotImplemented)
  w.Write([]byte("501 Not Implementes "))*/
  w.Header().Set("http.StatusNotImplemented", "501 Not Implementes ")
  fmt.Fprintf(w, "Laboratory #1")


}


func miNombre(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Header2", "CREATED: 201")
  /*w.WriteHeader(http.StatusOK)
  w.Write([]byte("200 OK "))*/
      w.Header().Set("http.StatusOK", "200 OK ")
  fmt.Fprintf(w, "Vladimir Rodriguez Bahena")
}


func main() {
  http.HandleFunc("/", darMensaje)
  http.HandleFunc("/home/vladimir/Web_aplications/go/src/Lab1", miNombre)
  err := http.ListenAndServe("localhost"+":"+"8080", nil)
  if err != nil {
    log.Fatal("error en el servidor : ", err)
    return
  }
}


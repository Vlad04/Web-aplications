package main
import 
(
  "fmt"
  "log"
  "net/http"
  "compress/gzip"
  "io"
  "strings"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func makeGzipHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		fn(gzr, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is a test."))
}

func darMensaje(w http.ResponseWriter, r *http.Request) {
  	w.Header().Set("Content-Type", "text/plain")
  
    w.WriteHeader(http.StatusNotImplemented)
  	w.Write([]byte("501 Not Implementes"))
    
    /*w.WriteHeader(http.StatusCreated)
  	w.Write([]byte("201 Created "))*/

  fmt.Fprintf(w, "Laboratory #1")


}


func miNombre(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Vladimir Rodriguez Bahena")
}


func main() {
  http.HandleFunc("/", darMensaje)
  //http.HandleFunc("/home/vladimir/Web_aplications/go/src/Lab1", miNombre)
  http.HandleFunc("/zip", makeGzipHandler(handler))
  err := http.ListenAndServe("localhost"+":"+"8080", nil)
  //http.ListenAndServe(":8080", makeGzipHandler(darMensaje))
  if err != nil {
    log.Fatal("error en el servidor : ", err)
    return
  }
}


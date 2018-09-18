package main

import (
  "fmt"
  "html/template"
  "net/http"
  "os"
  "io"
)

type Libro struct {
  Titulo string;
  Autor string;
}

type Nazi struct{
  Nombre string;
  Campo string;
  Perro string;
}

func darMensaje(w http.ResponseWriter, r *http.Request) {
  libro := Libro{Titulo: "En busca del tiempo", Autor: "Ar Turo"}
  
  pTemplate, _ := template.ParseFiles("plantilla.html")
  err := pTemplate.Execute(w, libro)
  if err != nil {
    fmt.Printf("Error: ", err)
    return
  }
}

func darOtroMensaje(w http.ResponseWriter, r *http.Request) {
  nazi := Nazi{Nombre: "Adolf", Campo: "Auschwitz", Perro: "Aria"}
  
  pTemplate, _ := template.ParseFiles("plantillaNazi.html")
  err := pTemplate.Execute(w, nazi)
  if err != nil {
    fmt.Printf("Error: ", err)
    return
  }
}

func irAForma(w http.ResponseWriter, r *http.Request){

  if (r.Method == "POST"){
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Fprint(w, r.FormValue("lastname"))
  } else {
    pTemplate, _ := template.ParseFiles("forma.html")
    err := pTemplate.Execute(w, nil)
    if err != nil {
      fmt.Printf("Error: ", err)
      return
    }
  }
}

func setGalleta(w http.ResponseWriter, r *http.Request){
  g1 := http.Cookie{
    Name: "Alan",
    Value: "80",
  }

  g2 := http.Cookie{
    Name: "Vladimir",
    Value: "90",
  }

  g3 := http.Cookie{
    Name: "Aranza",
    Value: "20",
  }

  http.SetCookie(w, &g1)
  http.SetCookie(w, &g2)
  http.SetCookie(w, &g3)
}

func getGalleta (w http.ResponseWriter, r*http.Request){

    valor,_:=r.Cookie("Alan")
    fmt.Fprint(w,valor,"\n")


    
    valor_b,_:=r.Cookie("Vladimir")
    fmt.Fprint(w,valor_b,"\n")


    
    valor_c,_:=r.Cookie("Aranza")
    fmt.Fprint(w,valor_c,"\n")


}

func procesarArchivo (w http.ResponseWriter, r*http.Request){

    if r.Method=="POST"{
        r.ParseMultipartForm(4096)
        archivo,header,err:=r.FormFile("subir")
        if err != nil{
            fmt.Printf("Error archivo")
            return
            }
            
            defer archivo.Close()
            fmt.Fprintf(w,"%v", header.Header)
            file,err:=os.OpenFile(header.Filename,  os.O_APPEND|os.O_CREATE|os.O_WRONLY , 0644)
            if err != nil{
            fmt.Printf("Error archivo")
            return
            }
            defer file.Close()
            
            io.Copy(file,archivo)
            
    }else{
        pTemplate,_:=template.ParseFiles("archivo.html")
        err:=pTemplate.Execute(w,nil)
        if err != nil{
            fmt.Printf("Error: ",err)
            return
            }
     }
}


func main() {
  http.HandleFunc("/libro", darMensaje)
  http.HandleFunc("/nazi", darOtroMensaje)
  http.HandleFunc("/forma", irAForma)
  http.HandleFunc("/setGalleta", setGalleta)
  http.HandleFunc("/getGalleta", getGalleta)
  http.HandleFunc("/procesarArchivo", procesarArchivo)
  err := http.ListenAndServe("localhost"+":"+"8080", nil)
  if err != nil {
    return
  }
}

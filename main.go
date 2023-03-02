package main

import (
    "net/http"
    "log"
    "latihan-web/controllers"
)

func main() {
    
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", controllers.Home)
    mux.HandleFunc("/about", controllers.About)
    mux.HandleFunc("/form", controllers.Form)
    mux.HandleFunc("/message", controllers.Message)
    
    fileServer := http.FileServer(http.Dir("assets"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))
    
    server := http.Server{
        Addr: "localhost:8080", 
        Handler: mux, 
    }
    
    log.Println("Menjalankan Web Server...")
    err := server.ListenAndServe()
    log.Fatal(err)
}
package controllers 

import (
    "log"
    "net/http"
    "path"
    "html/template"
)

func Message(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Something Went Wrong", http.StatusBadRequest)
        return
    }
    
    err := r.ParseForm()
    if err != nil {
        log.Println(err)
        http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
        return
    }
    
    nama := r.Form.Get("nama")
    pesan := r.Form.Get("pesan")
    
    data := map[string]any{
        "nama": nama, 
        "pesan": pesan, 
    }
    
    tmpl, err := template.ParseFiles(path.Join("views", "message.html"), path.Join("views", "layout.html"))
    if err != nil {
        log.Println(err)
        http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
        return
    }
    
    err = tmpl.Execute(w, data)
    if err != nil {
        log.Println(err)
        http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
        return
    }
}
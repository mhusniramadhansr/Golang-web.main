package controllers

import (
    "net/http"
    "html/template"
    "path"
    "log"
)

func Home(w http.ResponseWriter, r *http.Request) {
    
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    
    tmpl, err := template.ParseFiles(path.Join("views", "home.html"), path.Join("views", "layout.html"))
    if err != nil {
        log.Println(err)
        http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
        return
    }
    
    err = tmpl.Execute(w, nil)
    if err != nil {
        log.Println(err)
        http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
        return
    }
}
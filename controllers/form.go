package controllers

import (
    "log"
    "net/http"
    "html/template"
    "path"
)

func Form(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
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
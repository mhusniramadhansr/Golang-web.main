package models

import (
    "latihan-web/entity"
)

func GetAllData() []entity.Posts {
    data := []entity.Posts{
        {
            Id: "judul-pertama", 
            Title: "Judul Pertama", 
            Desc: "Lorem ipsum dolor sit amet", 
        },
        {
            Id: "judul-kedua", 
            Title: "Judul kedua", 
            Desc: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatem quae a incidunt nobis in porro modi rerum ea laborum consequatur.", 
        },
        {
            Id: "judul-ketiga", 
            Title: "Judul ketiga", 
            Desc: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatem quae a incidunt nobis in porro modi rerum ea laborum consequatur, eveniet non vero molestias saepe fugit, deserunt quia, voluptatum esse.", 
        },
    }
    
    return data
}
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.Static("/static", "./static")
    r.LoadHTMLGlob("templates/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "home.tmpl", gin.H{
            "title": "Cyber<span style='color:#dc3545'>Security</span> Manager",
            "active": "home",
        })
    })

    r.GET("/about", func(c *gin.Context) {
        c.HTML(200, "about.tmpl", gin.H{
            "title": "About Me – Cybersecurity Manager",
            "active": "about",
        })
    })

    r.Run(":8080")
}
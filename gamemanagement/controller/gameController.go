package controller

import (
    "../model"
    "github.com/gin-gonic/gin"
    "net/http"
)

func ShowHome(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl.html", gin.H{})
}


func selectDescriptionByName(gameName string) string {
    g := model.Games{}
    GetDB().Where(&model.Games{Name: gameName}).First(&g)
    return g.Description
}

func insertGame(data map[string]string) {
    g := model.Games{
        Name: data["name"],
		Description: data["description"],
		IDPlatform: 2,
		Plus: 0,
	}
	GetDB().Create(&g)
}

func insertPlatform(name string) {
	p := model.Platform {
		Name: name,
	}
	GetDB().Create(&p)
}

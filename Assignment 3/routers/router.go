package routers

import (
	"assignment3/controllers"
	"assignment3/models"
	"html/template"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	status := make(chan models.Status)
	go controllers.UpdateStatus(status)

	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		latestStatus := <-status

		waterStatus, windStatus := controllers.GetStatusText(latestStatus)

		tpl, err := template.ParseFiles("template.html")
		if err != nil {
			c.String(500, err.Error())
			return
		}

		err = tpl.Execute(c.Writer, map[string]interface{}{
			"Water":       latestStatus.Water,
			"WaterStatus": waterStatus,
			"Wind":        latestStatus.Wind,
			"WindStatus":  windStatus,
		})
		if err != nil {
			c.String(500, err.Error())
			return
		}
	})
	return router
}

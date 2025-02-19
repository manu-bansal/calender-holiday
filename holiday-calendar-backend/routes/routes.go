package routes

import (
	"holiday-calendar/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) { // Now it accepts *gin.Engine as an argument
	r.POST("/api/holidays", controllers.AddHoliday)
	r.GET("/api/holidays", controllers.GetHolidays)
	r.DELETE("/api/holidays/:id", controllers.DeleteHoliday)
}

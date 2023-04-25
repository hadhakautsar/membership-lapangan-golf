package routes

import (
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"membership-lapangan-golf/controllers"
	"membership-lapangan-golf/middleware"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	memberController := controllers.NewMemberController(db)
	adminController := controllers.NewAdminController(db)
	teeTimeController := controllers.NewTeeTimeController(db)

	// Member routes
	e.POST("/members/register", memberController.Register)
	e.POST("/members/login", memberController.Login)

	// BasicAuth middleware for authentication
	auth := e.Group("/auth")
	auth.Use(mid.BasicAuth(middleware.BasicAuthDB))
	auth.POST("/members/teetimes", memberController.BookTeeTime)

	// Admin routes
	e.POST("/admin/members", adminController.Create)
	e.GET("/admin/members", adminController.ReadAll)
	e.GET("/admin/members/:id", adminController.Read)
	e.PUT("/admin/members/:id", adminController.Update)
	e.DELETE("/admin/members/:id", adminController.Delete)
	e.GET("/admin/teetimes", adminController.GetTeeTimes)

	// Tee time routes
	e.GET("/teetimes", teeTimeController.GetAllTeeTimes)
	e.GET("/teetimes/available", teeTimeController.GetAvailableTeeTimes)
	e.POST("/teetimes", teeTimeController.CreateTeeTime)
	e.PUT("/teetimes/:id", teeTimeController.UpdateTeeTime)
	e.DELETE("/teetimes/:id", teeTimeController.DeleteTeeTime)
}

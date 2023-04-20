package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"membership-lapangan-golf/controllers"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	memberController := controllers.NewMemberController(db)
	adminController := controllers.NewAdminController(db)
	teeTimeController := controllers.NewTeeTimeController(db)

	// Member routes
	e.POST("/members/register", memberController.Register)
	e.POST("/members/login", memberController.Login)

	// Admin routes
	// e.GET("/admin/members", adminController.GetAllMembers)
	e.GET("/admin/members/:id", adminController.Read)
	// e.PUT("/admin/members/:id", adminController.UpdateMember)
	e.DELETE("/admin/members/:id", adminController.Delete)

	// Tee time routes
	e.GET("/teetimes", teeTimeController.GetAllTeeTimes)
	e.GET("/teetimes/available", teeTimeController.GetAvailableTeeTimes)
	e.POST("/teetimes", teeTimeController.CreateTeeTime)
	e.PUT("/teetimes/:id", teeTimeController.UpdateTeeTime)
	e.DELETE("/teetimes/:id", teeTimeController.DeleteTeeTime)
}

package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"membership-lapangan-golf/models"
)

type TeeTimeController struct {
	db *gorm.DB
}

// Constructor for TeeTimeController
func NewTeeTimeController(db *gorm.DB) *TeeTimeController {
	return &TeeTimeController{db}
}

// Get all tee times
func (ttc *TeeTimeController) GetAllTeeTimes(c echo.Context) error {
	var teeTimes []models.TeeTime
	if err := ttc.db.Preload("Member").Find(&teeTimes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get tee times")
	}

	return c.JSON(http.StatusOK, teeTimes)
}

// Create a new tee time
func (ttc *TeeTimeController) CreateTeeTime(c echo.Context) error {
	teeTime := new(models.TeeTime)
	if err := c.Bind(teeTime); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	member := new(models.Member)
	if err := ttc.db.First(member, teeTime.MemberID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to book tee time")
	}

	teeTime.Member = *member

	if err := ttc.db.Create(&teeTime).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create tee time")
	}

	return c.JSON(http.StatusCreated, teeTime)
}

// Update an existing tee time
func (ttc *TeeTimeController) UpdateTeeTime(c echo.Context) error {
	teeTimeID := c.Param("id")

	var teeTime models.TeeTime
	if err := ttc.db.First(&teeTime, teeTimeID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Tee time not found")
	}

	if err := c.Bind(&teeTime); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := ttc.db.Save(&teeTime).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update tee time")
	}

	return c.JSON(http.StatusOK, teeTime)
}

// Delete a tee time
func (ttc *TeeTimeController) DeleteTeeTime(c echo.Context) error {
	teeTimeID := c.Param("id")

	var teeTime models.TeeTime
	if err := ttc.db.First(&teeTime, teeTimeID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Tee time not found")
	}

	if err := ttc.db.Delete(&teeTime).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete tee time")
	}

	return c.NoContent(http.StatusNoContent)
}

// Get available tee times
func (ttc *TeeTimeController) GetAvailableTeeTimes(c echo.Context) error {
	dateStr := c.QueryParam("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid date format")
	}

	var teeTimes []struct {
		ID       uint      `json:"id"`
		MemberID uint      `json:"memberId"`
		Time     time.Time `json:"time"`
	}
	if err := ttc.db.Model(&models.TeeTime{}).
		Select("id, member_id, time").
		Where("time >= ? AND time < ?", date, date.AddDate(0, 0, 1)).
		Scan(&teeTimes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get tee times")
	}

	for i := range teeTimes {
		teeTimes[i].Time = teeTimes[i].Time.UTC().Truncate(24 * time.Hour)
	}

	return c.JSON(http.StatusOK, teeTimes)
}

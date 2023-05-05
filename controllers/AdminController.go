package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"membership-lapangan-golf/models"
)

type AdminController struct {
	db *gorm.DB
}

// Constructor for AdminController
func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{db}
}

// Create a member
func (ac *AdminController) Create(c echo.Context) error {
	member := new(models.Member)
	if err := c.Bind(member); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := ac.db.Create(&member).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create member")
	}

	member.Password = ""

	return c.JSON(http.StatusCreated, member)
}

// Read all member
func (ac *AdminController) ReadAll(c echo.Context) error {
	var members []models.Member
	if err := ac.db.Find(&members).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get members")
	}

	for i := range members {
		members[i].Password = ""
	}

	return c.JSON(http.StatusOK, members)
}

// Read a member
func (ac *AdminController) Read(c echo.Context) error {
	id := c.Param("id")

	var member models.Member
	if err := ac.db.First(&member, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Member not found")
	}

	member.Password = ""

	return c.JSON(http.StatusOK, member)
}

// Update a member
func (ac *AdminController) Update(c echo.Context) error {
	id := c.Param("id")

	var member models.Member
	if err := ac.db.First(&member, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Member not found")
	}

	newMember := new(models.Member)
	if err := c.Bind(newMember); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if newMember.Username != "" {
		member.Username = newMember.Username
	}
	if newMember.Email != "" {
		member.Email = newMember.Email
	}
	if newMember.Password != "" {
		member.Password = newMember.Password
	}
	if newMember.Handicap != 0 {
		member.Handicap = newMember.Handicap
	}
	if newMember.Score != 0 {
		member.Score = newMember.Score
	}

	if err := ac.db.Save(&member).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update member")
	}

	member.Password = ""

	return c.JSON(http.StatusOK, member)
}

// Delete a member
func (ac *AdminController) Delete(c echo.Context) error {
	id := c.Param("id")

	var member models.Member
	if err := ac.db.First(&member, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Member not found")
	}

	if err := ac.db.Delete(&member).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete member")
	}

	return c.JSON(http.StatusOK, "Member deleted successfully")
}

// Get all tee times
func (ac *AdminController) GetTeeTimes(c echo.Context) error {
	var teeTimes []models.TeeTime
	if err := ac.db.Find(&teeTimes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get tee times")
	}

	return c.JSON(http.StatusOK, teeTimes)
}

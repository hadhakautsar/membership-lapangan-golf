package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"membership-lapangan-golf/models"
)

type MemberController struct {
	db *gorm.DB
}

// Constructor for MemberController
func NewMemberController(db *gorm.DB) *MemberController {
	return &MemberController{db}
}

// Register a new member
func (mc *MemberController) Register(c echo.Context) error {
	member := new(models.Member)
	if err := c.Bind(member); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := mc.db.Create(&member).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create member")
	}

	member.Password = ""

	return c.JSON(http.StatusCreated, member)
}

// Login a member
func (mc *MemberController) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	var member models.Member
	if err := mc.db.Where("email = ? AND password = ?", email, password).First(&member).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}

	member.Password = ""

	return c.JSON(http.StatusOK, member)
}

// Book a tee time
func (mc *MemberController) BookTeeTime(c echo.Context) error {
	teeTime := new(models.TeeTime)
	if err := c.Bind(teeTime); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	member := new(models.Member)
	if err := mc.db.First(member, teeTime.MemberID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to book tee time")
	}

	teeTime.Member = *member

	if err := mc.db.Create(&teeTime).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to book tee time")
	}

	teeTime.Member.Password = ""

	return c.JSON(http.StatusCreated, teeTime)
}

// Get member's score
func (mc *MemberController) GetScore(c echo.Context) error {
	memberID := c.Get("user").(models.Member).ID

	var member models.Member
	if err := mc.db.First(&member, memberID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Member not found")
	}

	return c.JSON(http.StatusOK, member.Score)
}

// Get member's handicap
func (mc *MemberController) GetHandicap(c echo.Context) error {
	memberID := c.Get("user").(models.Member).ID

	var member models.Member
	if err := mc.db.First(&member, memberID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Member not found")
	}

	return c.JSON(http.StatusOK, member.Handicap)
}

// Get member's ranking
func (mc *MemberController) GetRanking(c echo.Context) error {
	var members []models.Member
	if err := mc.db.Find(&members).Order("score desc").Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get members")
	}

	ranking := make(map[uint]int)
	for i, member := range members {
		ranking[member.ID] = i + 1
	}

	memberID := c.Get("user").(models.Member).ID
	return c.JSON(http.StatusOK, ranking[memberID])
}

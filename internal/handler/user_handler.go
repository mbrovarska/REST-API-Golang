package handler

import (
	"net/http"

	models "example.com/rest-api-notes/internal/domain"
	"example.com/rest-api-notes/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService *service.UserService
	logger *zap.Logger
}

func NewUserHandler(userService *service.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger: logger,
	}
}

// SignUp godoc
// @Summary      User Sign Up
// @Description  Register a new user account with email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      models.SignUpRequest  true  "Sign up credentials"
// @Success      201      {object}  models.AuthResponse
// @Failure      400      {object}  map[string]string "Invalid request body or validation error"
// @Failure      409      {object}  map[string]string "Email already registered"
// @Failure      500      {object}  map[string]string "Internal server error"
// @Router       /auth/signup [post]
func  (h *UserHandler) SignUp(c *gin.Context) {
	var req models.SignUpRequest

	//bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Invalid signup request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	//call service
	response, err := h.userService.SignUp(c.Request.Context(), req)
	if err != nil {
		// Handle specific errors
		if err.Error() == "email already registered" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Email already registered",
			})
			return
		}

		h.logger.Error("Failed to sign up user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create account. Please try again later.",
			"debug_error":  err.Error(),
		})
		return	
	}

	h.logger.Info("User signed up successfully",
		zap.Int64("user_id", response.User.ID),
		zap.String("email", response.User.Email),
	)
	c.JSON(http.StatusCreated, response)
}
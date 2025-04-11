package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/yuhari7/backend_supervision/internal/common/dto"
	"github.com/yuhari7/backend_supervision/internal/usecase/user"
	jwtutil "github.com/yuhari7/backend_supervision/pkg/jwt"
)

type UserController struct {
	Usecase user.UserUsecase
}

func NewUserController(u user.UserUsecase) *UserController {
	return &UserController{Usecase: u}
}

func (h *UserController) Register(c echo.Context) error {
	var input dto.CreateUserRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	// Validasi dengan validator
	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	newUser, err := h.Usecase.Register(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	response := dto.UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
		Role:  newUser.RoleID,
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "user registered successfully",
		"user":    response,
	})
}

func (h *UserController) Login(c echo.Context) error {
	var input dto.LoginRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user, err := h.Usecase.Login(input)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}

	accessToken, err := jwtutil.GenerateAccessToken(user.ID, user.Email, user.RoleID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate access token"})
	}

	refreshToken, err := jwtutil.GenerateRefreshToken(user.ID, user.Email, user.RoleID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate refresh token"})
	}

	response := dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.RoleID,
		},
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserController) RefreshToken(c echo.Context) error {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	claims, err := jwtutil.ParseRefreshToken(body.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid or expired refresh token"})
	}

	// generate new access token
	newAccessToken, err := jwtutil.GenerateAccessToken(claims.UserID, claims.Email, claims.RoleID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate access token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token": newAccessToken,
	})
}

// Users

func (h *UserController) GetAllUsers(c echo.Context) error {
	var pagination dto.PaginationQuery
	if err := c.Bind(&pagination); err != nil {
		pagination = dto.PaginationQuery{Page: 1, Limit: 10}
	}

	result, err := h.Usecase.GetAllUsers(pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *UserController) GetUserByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	user, err := h.Usecase.GetUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserController) CreateUser(c echo.Context) error {
	var input dto.CreateUserRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	// Validasi dengan validator
	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Bisa reuse usecase Register karena logikanya sama
	newUser, err := h.Usecase.Register(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "user created successfully",
		"user":    newUser,
	})
}

func (h *UserController) UpdateUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	var input dto.UpdateUserRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	updatedUser, err := h.Usecase.UpdateUser(uint(id), input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	response := dto.UserResponse{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
		Role:  updatedUser.RoleID,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "user updated successfully",
		"user":    response,
	})
}

func (h *UserController) DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	err = h.Usecase.DeleteUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "user deleted successfully"})
}

func (h *UserController) DeactivateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	err = h.Usecase.ToggleUserActive(uint(id), false)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
		}
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "user deactivated"})
}

func (h *UserController) ActivateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	err = h.Usecase.ToggleUserActive(uint(id), true)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
		}
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "user activated"})
}

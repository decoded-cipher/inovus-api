package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/decoded-cipher/inovus-api/internal"
	"github.com/decoded-cipher/inovus-api/models"
	"github.com/labstack/echo/v4"
)

// getAllUsers returns all users with pagination
func getAllUsers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// Get users
	var users []models.User
	err := queries.GetAllUsers.Select(&users, pageSize, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch users",
		})
	}

	// Get total count
	var total int
	err = queries.GetUsersCount.Get(&total)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to count users",
		})
	}

	response := internal.Paginate(users, page, pageSize, total)
	return c.JSON(http.StatusOK, response)
}

// getUser returns a single user by id, username, or email
func getUser(c echo.Context) error {
	// Try to get by ID first
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	// Get query parameters
	username := c.QueryParam("username")
	email := c.QueryParam("email")

	var user models.User
	err := queries.GetUser.Get(&user, id, username, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "User not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch user",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// createUser creates a new user
func createUser(c echo.Context) error {
	var req models.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// TODO: Hash password before storing
	var user models.User
	err := queries.CreateUser.Get(&user, req.Username, req.Email, req.PasswordHash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create user",
		})
	}

	return c.JSON(http.StatusCreated, user)
}

// updateUser updates an existing user
func updateUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID",
		})
	}

	var req models.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Get existing user first
	var existingUser models.User
	err = queries.GetUser.Get(&existingUser, id, "", "")
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "User not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch user",
		})
	}

	// Use existing values if not provided in request
	username := req.Username
	if username == "" {
		username = existingUser.Username
	}
	email := req.Email
	if email == "" {
		email = existingUser.Email
	}
	passwordHash := req.PasswordHash
	if passwordHash == "" {
		passwordHash = existingUser.PasswordHash
	}

	// Update user
	var user models.User
	err = queries.UpdateUser.Get(&user, username, email, passwordHash, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update user",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// deleteUser deletes a user by id
func deleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID",
		})
	}

	// Check if user exists
	var user models.User
	err = queries.GetUser.Get(&user, id, "", "")
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "User not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch user",
		})
	}

	// Delete user
	_, err = queries.DeleteUser.Exec(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete user",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}

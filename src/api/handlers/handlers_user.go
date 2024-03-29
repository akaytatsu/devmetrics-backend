package handlers

import (
	"app/entity"
	infra_jira "app/infrastructure/jira"
	infra_mongodb "app/infrastructure/mongodb"
	"app/infrastructure/repository"
	usecase_kanban "app/usecase/kanban"
	usecase_user "app/usecase/user"
	"net/http"
	"strconv"

	middleware "app/api/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPasswordData struct {
	Email           string `json:"email"`
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserHandlers struct {
	UsecaseUser usecase_user.IUsecaseUser
}

func NewUserHandler(usecaseUser usecase_user.IUsecaseUser) *UserHandlers {
	return &UserHandlers{UsecaseUser: usecaseUser}
}

func (h UserHandlers) LoginHandler(c *gin.Context) {

	var loginData LoginData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		handleError(c, err)
		return
	}

	user, err := h.UsecaseUser.LoginUser(loginData.Email, loginData.Password)

	if exception := handleError(c, err); exception {
		return
	}

	token, refreshToken, err := usecase_user.JWTTokenGenerator(*user)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"token": token, "refreshToken": refreshToken})
}

func (h UserHandlers) GetMeHandler(c *gin.Context) {
	user, err := h.UsecaseUser.GetUserByToken(c.GetHeader("Authorization"))

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, user)
}

func (h UserHandlers) CreateUserHandler(c *gin.Context) {

	var entityUser entity.EntityUser

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseUser.Create(&entityUser)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "User created successfully"})

}

func (h UserHandlers) UpdateUserHandler(c *gin.Context) {

	var entityUser entity.EntityUser

	id := strconv.Itoa(c.GetInt("id"))

	entityUser.ID = id

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseUser.Update(&entityUser)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h UserHandlers) DeleteUserHandler(c *gin.Context) {

	var entityUser entity.EntityUser

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseUser.Delete(&entityUser)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h UserHandlers) UpdatePasswordHandler(c *gin.Context) {

	var updatePasswordData UpdateUserPasswordData

	if err := c.ShouldBindJSON(&updatePasswordData); err != nil {
		handleError(c, err)
		return
	}

	id := c.Param("id")

	err := h.UsecaseUser.UpdatePassword(id, updatePasswordData.OldPassword, updatePasswordData.NewPassword, updatePasswordData.ConfirmPassword)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func (h UserHandlers) GetUsersHandler(c *gin.Context) {

	var filters entity.EntityUserFilters

	filters.Search = c.Query("search")
	filters.Active = c.Query("active")

	users, err := h.UsecaseUser.GetUsers(filters)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, users)
}

func (h UserHandlers) GetUserHandler(c *gin.Context) {

	id := c.Param("id")

	user, err := h.UsecaseUser.GetUser(id)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, user)
}

func (h UserHandlers) GetDataHandler(c *gin.Context) {

	usecaseKanban := usecase_kanban.NewService(
		repository.NewKanbanJira(
			infra_jira.Connect(),
			infra_mongodb.Connect(),
		),
	)

	usecaseKanban.UpdateIssues()
}

func MountUsersHandlers(gin *gin.Engine, conn *mongo.Database) {

	userHandlers := NewUserHandler(
		usecase_user.NewService(
			repository.NewUserPostgres(conn),
		),
	)

	gin.GET("/", HomeHandler)
	gin.GET("/sssssssss", userHandlers.GetDataHandler)
	gin.POST("/api/login", userHandlers.LoginHandler)

	gin.POST("/login", userHandlers.LoginHandler)

	// user
	group := gin.Group("/api/user")
	group.Use(middleware.AuthenticatedMiddleware(userHandlers.UsecaseUser))
	group.GET("/me", userHandlers.GetMeHandler)
	group.POST("/create", userHandlers.CreateUserHandler)
	group.PUT("/:id", userHandlers.UpdateUserHandler)
	group.DELETE("/:id", userHandlers.DeleteUserHandler)
	group.PUT("/password/:id", userHandlers.UpdatePasswordHandler)
	group.GET("/list", userHandlers.GetUsersHandler)
	group.GET("/:id", userHandlers.GetUserHandler)
}

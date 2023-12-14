package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-project/api/handler"
	"go-project/api/repository"
	"go-project/api/usecase"
	"go-project/db"
	"go-project/models"
	"go-project/utils/log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func SetUpHandlerUser() *handler.UserHandler {
	if err := db.InitConnectionDB(); err != nil {
		log.Log.Error("Error Connection DB ", err.Error())
	}
	dbPg, err := db.GetConnectionDB()
	if err != nil {
		fmt.Println("err connection", err)
		log.Log.Error("Error to Get Connection DB")
	}
	repoUser := &repository.UserRepository{DB: dbPg}
	userUsecase := &usecase.UserUsecase{URepository: *repoUser}
	userHandler := &handler.UserHandler{Uusecase: *userUsecase}

	return userHandler
}

func TestHealthBase(t *testing.T) {

	route := gin.New()
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Tech Company listing API with Golang"})
	})
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	route.ServeHTTP(rec, req)
	fmt.Println("hasil", rec)

	assert.Equal(t, http.StatusOK, rec.Code)

	// var response map[string]interface{}
	// err = json.Unmarshal(rec.Body.Bytes(), &response)
	// assert.NoError(t, err)

	// fmt.Println("status code", response)
	// assert.Equal(t, "ok", response["status"])
}

func TestFuncCoba(t *testing.T) {

	userUsecase := usecase.UserUsecase{}
	userHandler := handler.UserHandler{Uusecase: userUsecase}

	router := gin.Default()
	router.GET("/user/api/tes", userHandler.CobaTest)

	req, err := http.NewRequest("GET", "/user/api/tes", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	fmt.Println("hasil", rec)

	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestRegisterUser(t *testing.T) {
	handler := SetUpHandlerUser()
	router := gin.Default()
	router.POST("/user/api/register", handler.RegisterUser)

	reqBody := `{"name":"abdsamsulul","email":"woepwmksd5674582@mail.com","password" :"1234sp","role":"user"}`
	req, err := http.NewRequest("POST", "/user/api/register", bytes.NewBuffer([]byte(reqBody)))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "SUCCESS", response["respMessage"])

}

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) CreateUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func TestRegisterUserUseMock(t *testing.T) {
	MockuserUsecase1 := new(MockUserUsecase)

	MockuserUsecase1.On("CreateUser", mock.Anything).Return(models.User{
		Name:     "abdsamsulul",
		Email:    "woepwmksd5674582@mail.com",
		Password: "1234sp",
		Role:     "user",
	}, nil)

	handler := SetUpHandlerUser()

	router := gin.Default()
	router.POST("/user/api/register", handler.RegisterUser)

	tests := []struct {
		name           string
		reqBody        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful registration",
			reqBody:        `{"name":"abdsamsulul","email":"woepwmksd5674582@mail.com","password":"1234sp","role":"user"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"respMessage":"SUCCESS"}`, // Adjust the expected response based on your actual implementation
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the UserUseCase method
			MockuserUsecase1.On("CreateUser", mock.Anything).Return(models.User{}, nil)

			// Create a new HTTP request for testing
			req, err := http.NewRequest("POST", "/user/api/register", bytes.NewBuffer([]byte(tt.reqBody)))
			assert.NoError(t, err)

			// Create a response recorder to capture the response
			w := httptest.NewRecorder()

			// Serve the HTTP request to the router
			router.ServeHTTP(w, req)

			// Check the HTTP status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Check the response body
			var response map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			// Adjust the assertion based on your actual implementation
			assert.Equal(t, "SUCCESS", response["respMessage"])

			// Assert that the expected method on the mock was called
			MockuserUsecase1.AssertExpectations(t)
		})
	}
}

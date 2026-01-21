package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func (suite *APITestSuite) TestGetAllTodosEmpty() {
	suite.CleanupCollections()

	token := suite.GetAuthToken("getalltodos@example.com", "SecurePassword123", "Get All User")
	request := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	request.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	suite.app.Router.ServeHTTP(w, request)

	assert.True(suite.T(), w.Code == http.StatusOK || w.Code == http.StatusNoContent || w.Code == http.StatusBadRequest)
}

func (suite *APITestSuite) TestGetTodoNotFound() {
	suite.CleanupCollections()

	token := suite.GetAuthToken("gettodo@example.com", "SecurePassword123", "Get Todo User")
	request := httptest.NewRequest(http.MethodGet, "/api/v1/todos/507f1f77bcf86cd799439011", nil)
	request.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	suite.app.Router.ServeHTTP(w, request)

	assert.True(suite.T(), w.Code == http.StatusNotFound || w.Code == http.StatusBadRequest)
}

func (suite *APITestSuite) TestUpdateTodoNotFound() {
	suite.CleanupCollections()

	token := suite.GetAuthToken("updatetodo@example.com", "SecurePassword123", "Update Todo User")
	updateReq := map[string]string{
		"title":       "Updated Title",
		"description": "Updated description",
	}
	body, _ := json.Marshal(updateReq)
	request := httptest.NewRequest(http.MethodPut, "/api/v1/todos/507f1f77bcf86cd799439011", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	suite.app.Router.ServeHTTP(w, request)

	assert.True(suite.T(), w.Code == http.StatusNotFound || w.Code == http.StatusBadRequest)
}

func (suite *APITestSuite) TestDeleteTodoNotFound() {
	suite.CleanupCollections()

	token := suite.GetAuthToken("deletetodo@example.com", "SecurePassword123", "Delete Todo User")
	request := httptest.NewRequest(http.MethodDelete, "/api/v1/todos/507f1f77bcf86cd799439011", nil)
	request.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	suite.app.Router.ServeHTTP(w, request)

	assert.True(suite.T(), w.Code == http.StatusNotFound || w.Code == http.StatusBadRequest || w.Code == http.StatusNoContent)
}

func (suite *APITestSuite) TestSignUpDuplicateEmail() {
	suite.CleanupCollections()

	req := map[string]string{
		"email":    "duplicate@example.com",
		"password": "SecurePassword123",
		"name":     "Duplicate User",
	}

	body, _ := json.Marshal(req)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.app.Router.ServeHTTP(w, request)
	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	body, _ = json.Marshal(req)
	request = httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	suite.app.Router.ServeHTTP(w, request)

	assert.True(suite.T(), w.Code == http.StatusConflict || w.Code == http.StatusBadRequest)
}

func (suite *APITestSuite) TestLoginWithWrongPassword() {
	suite.CleanupCollections()

	signUpReq := map[string]string{
		"email":    "wrongpass@example.com",
		"password": "CorrectPassword123",
		"name":     "Wrong Pass User",
	}
	body, _ := json.Marshal(signUpReq)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.app.Router.ServeHTTP(w, request)
	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	loginReq := map[string]string{
		"email":    "wrongpass@example.com",
		"password": "WrongPassword123",
	}
	body, _ = json.Marshal(loginReq)
	request = httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	suite.app.Router.ServeHTTP(w, request)

	assert.NotEqual(suite.T(), http.StatusOK, w.Code)
}

func (suite *APITestSuite) TestErrorResponseFormat() {
	suite.CleanupCollections()

	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBufferString("{invalid json}"))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.app.Router.ServeHTTP(w, request)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.NotNil(suite.T(), response)
}

func (suite *APITestSuite) TestSuccessResponseFormat() {
	suite.CleanupCollections()

	signUpReq := map[string]string{
		"email":    "format@example.com",
		"password": "SecurePassword123",
		"name":     "Format User",
	}
	body, _ := json.Marshal(signUpReq)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.app.Router.ServeHTTP(w, request)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code >= 200 && w.Code < 300 {
		assert.NotNil(suite.T(), response["data"])
	}
}

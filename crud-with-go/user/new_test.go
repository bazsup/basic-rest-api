package user_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bazsup/crud-with-go/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockNewUserClosure(u user.User) user.SaveFunc {
	return func(user.User) (user.User, error) {
		return u, nil
	}
}

func TestNewHandler(t *testing.T) {
	t.Run("create success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := gin.Default()

		createdUser := user.User{ID: 1, Name: "bas"}
		h := user.NewHandler(mockNewUserClosure(createdUser))
		r.POST("/users", h)

		payload := strings.NewReader(`{
			"name": "bas"
		}`)
		req := httptest.NewRequest(http.MethodPost, "/users", payload)
		req.Header.Add("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		res := w.Result()

		expected := `{"name":"bas"}`
		body, _ := ioutil.ReadAll(res.Body)

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, expected, string(body))
	})

	t.Run("create fail", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := gin.Default()

		createdUser := user.User{ID: 1, Name: "bas"}
		h := user.NewHandler(mockNewUserClosure(createdUser))
		r.POST("/users", h)

		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		r.ServeHTTP(w, req)
		res := w.Result()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

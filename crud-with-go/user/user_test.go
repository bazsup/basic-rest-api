package user_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bazsup/crud-with-go/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockFindAllClosure(u []user.User) func() ([]user.User, error) {
	return func() ([]user.User, error) {
		return u, nil
	}
}

func TestUserHandler(t *testing.T) {
	t.Run("http status should be 200", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := gin.Default()
		h := user.UserHandler(mockFindAllClosure([]user.User{}))
		r.GET("/users", h)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		r.ServeHTTP(w, req)
		res := w.Result()

		want := http.StatusOK
		get := res.StatusCode

		if want != get {
			t.Errorf("want http.status %q but get %q", want, get)
		}
	})

	t.Run("body should contain a user in list", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := gin.Default()
		uRes := []user.User{
			{Name: "bas"},
		}
		h := user.UserHandler(mockFindAllClosure(uRes))
		r.GET("/users", h)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		r.ServeHTTP(w, req)

		want, _ := json.Marshal(uRes)
		get := w.Body.Bytes()

		assert.Equal(t, want, get)
	})
}

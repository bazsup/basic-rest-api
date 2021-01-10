package user_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bazsup/crud-with-go/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockDeleteSuccess(int64) error { return nil }

func TestDeleteHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := gin.Default()
		w := httptest.NewRecorder()

		calledDelete := 0
		mockDeleteSuccess := func(int64) error {
			calledDelete += 1
			return nil
		}
		mockFindById := func(int64) (user.User, error) {
			return user.User{}, nil
		}
		h := user.DeleteHandler(mockDeleteSuccess, mockFindById)
		r.DELETE("/users/:id", h)

		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		r.ServeHTTP(w, req)
		res := w.Result()

		body, _ := ioutil.ReadAll(res.Body)

		assert.Equal(t, 1, calledDelete)
		assert.Equal(t, http.StatusNoContent, res.StatusCode)
		assert.Empty(t, string(body))
	})

	t.Run("invalid id", func(t *testing.T) {
		r := gin.Default()
		w := httptest.NewRecorder()

		mockFindById := func(int64) (user.User, error) {
			return user.User{}, fmt.Errorf("user not found")
		}
		h := user.DeleteHandler(mockDeleteSuccess, mockFindById)
		r.DELETE("/users/:id", h)

		req := httptest.NewRequest(http.MethodDelete, "/users/abc", nil)
		r.ServeHTTP(w, req)
		res := w.Result()

		expected := jsonCompact(`{
			"err": "\"id\" must be integer"

		}`)
		body, _ := ioutil.ReadAll(res.Body)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, expected, string(body))
	})

	t.Run("notfound", func(t *testing.T) {
		r := gin.Default()
		w := httptest.NewRecorder()

		mockFindById := func(int64) (user.User, error) {
			return user.User{}, fmt.Errorf("user not found")
		}
		h := user.DeleteHandler(mockDeleteSuccess, mockFindById)
		r.DELETE("/users/:id", h)

		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		r.ServeHTTP(w, req)
		res := w.Result()

		expected := jsonCompact(`{
			"err": "user not found"

		}`)
		body, _ := ioutil.ReadAll(res.Body)

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
		assert.Equal(t, expected, string(body))
	})
}

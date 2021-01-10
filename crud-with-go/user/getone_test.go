package user_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bazsup/crud-with-go/testhelper"
	"github.com/bazsup/crud-with-go/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockFindByIdClosure(u user.User, err error) user.FindByIdFunc {
	return func(int64) (user.User, error) {
		return u, err
	}
}

func TestGetOneHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := gin.Default()
		w := httptest.NewRecorder()

		now := testhelper.MockNow()
		expectedUser := user.User{ID: 1, Name: "alice", CreatedAt: now, UpdatedAt: now}
		h := user.GetOneHandler(mockFindByIdClosure(expectedUser, nil))
		r.GET("/users/:id", h)

		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		r.ServeHTTP(w, req)
		res := w.Result()

		expected := jsonCompact(`{
			"id": 1,
			"name": "alice",
			"created_at": "2021-01-10T15:00:00Z",
			"updated_at": "2021-01-10T15:00:00Z"
		}`)
		body, _ := ioutil.ReadAll(res.Body)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expected, string(body))
	})

	t.Run("invalid id", func(t *testing.T) {
		r := gin.Default()
		w := httptest.NewRecorder()

		h := user.GetOneHandler(mockFindByIdClosure(user.User{}, nil))
		r.GET("/users/:id", h)

		req := httptest.NewRequest(http.MethodGet, "/users/abc", nil)
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

		h := user.GetOneHandler(mockFindByIdClosure(user.User{}, fmt.Errorf("asdf")))
		r.GET("/users/:id", h)

		req := httptest.NewRequest(http.MethodGet, "/users/10", nil)
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

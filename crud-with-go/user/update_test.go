package user_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bazsup/crud-with-go/testhelper"
	"github.com/bazsup/crud-with-go/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockUpdateUserClosure(u user.User) user.UpdateFunc {
	return func(int64, user.User) error {
		return nil
	}
}

func mockNow() time.Time {
	now, _ := time.Parse(time.RFC3339, "2021-01-10T16:01:00Z")
	return now
}

func TestUpdateHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := gin.Default()

		createdUser := user.User{ID: 1, Name: "bas", CreatedAt: testhelper.MockNow()}

		h := user.UpdateHandler(
			mockUpdateUserClosure(createdUser),
			mockFindByIdClosure(createdUser, nil),
			mockNow,
		)
		r.PATCH("/users/:id", h)

		payload := strings.NewReader(`{
			"name": "bob"
		}`)
		req := httptest.NewRequest(http.MethodPatch, "/users/1", payload)
		req.Header.Add("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		res := w.Result()

		expected := jsonCompact(`{
			"id": 1,
			"name": "bob",
			"created_at": "2021-01-10T15:00:00Z",
			"updated_at": "2021-01-10T16:01:00Z"
		}`)
		body, _ := ioutil.ReadAll(res.Body)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expected, string(body))
	})

	t.Run("invalid id", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := gin.Default()

		createdUser := user.User{ID: 1, Name: "bas", CreatedAt: testhelper.MockNow()}

		h := user.UpdateHandler(
			mockUpdateUserClosure(createdUser),
			mockFindByIdClosure(createdUser, nil),
			mockNow,
		)
		r.PATCH("/users/:id", h)

		payload := strings.NewReader(`{
			"name": "bob"
		}`)
		req := httptest.NewRequest(http.MethodPatch, "/users/abcd", payload)
		req.Header.Add("Content-Type", "application/json")
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
		w := httptest.NewRecorder()
		r := gin.Default()

		h := user.UpdateHandler(
			mockUpdateUserClosure(user.User{}),
			mockFindByIdClosure(user.User{}, fmt.Errorf("not found")),
			mockNow,
		)
		r.PATCH("/users/:id", h)

		payload := strings.NewReader(`{
			"name": "bob"
		}`)
		req := httptest.NewRequest(http.MethodPatch, "/users/1", payload)
		req.Header.Add("Content-Type", "application/json")
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

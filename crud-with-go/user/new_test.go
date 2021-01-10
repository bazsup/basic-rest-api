package user_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bazsup/crud-with-go/testhelper"
	"github.com/bazsup/crud-with-go/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockNewUserClosure(u user.User) user.SaveFunc {
	return func(user.User) (int64, error) {
		return 1, nil
	}
}

func jsonCompact(str string) string {
	jsonByte := []byte(str)
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, jsonByte); err != nil {
		panic(err)
	}
	return buffer.String()
}

func TestNewHandler(t *testing.T) {
	t.Run("create success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := gin.Default()

		createdUser := user.User{ID: 1, Name: "bas"}

		h := user.NewHandler(mockNewUserClosure(createdUser), testhelper.MockNow)
		r.POST("/users", h)

		payload := strings.NewReader(`{
			"name": "bas"
		}`)
		req := httptest.NewRequest(http.MethodPost, "/users", payload)
		req.Header.Add("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		res := w.Result()

		expected := jsonCompact(`{
			"id": 1,
			"name": "bas",
			"created_at": "2021-01-10T15:00:00Z",
			"updated_at": "2021-01-10T15:00:00Z"
		}`)
		body, _ := ioutil.ReadAll(res.Body)

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, expected, string(body))
	})

	t.Run("create fail", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := gin.Default()

		createdUser := user.User{ID: 1, Name: "bas"}
		h := user.NewHandler(mockNewUserClosure(createdUser), testhelper.MockNow)
		r.POST("/users", h)

		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		r.ServeHTTP(w, req)
		res := w.Result()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

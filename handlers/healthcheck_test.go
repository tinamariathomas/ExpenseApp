package handlers_test

import (
	"testing"

	h "git.expense-app.com/ExpenseApp/handlers"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	h.HealthCheckHandler(w,r)

	assert.Equal(t, []byte("Server check. I'm alive!"), w.Body.Bytes())
}


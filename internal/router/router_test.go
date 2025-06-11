package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/harshau007/listmanager/internal/listmanager"
	"github.com/harshau007/listmanager/internal/models"
	"github.com/harshau007/listmanager/internal/router"
	"github.com/stretchr/testify/assert"
)

func setup() (*gin.Engine, *listmanager.Manager) {
	gin.SetMode(gin.TestMode)
	lm := listmanager.New()
	r := router.New(lm)
	return r, lm
}

func TestAddAPI(t *testing.T) {
	r, _ := setup()
	cases := []struct {
		input  float64
		expect []float64
	}{
		{5, []float64{5}},
		{10, []float64{5, 10}},
		{-6, []float64{9}},
	}
	for _, cse := range cases {
		b, _ := json.Marshal(models.AddRequest{Number: int(cse.input)})
		req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		var resp models.AddResponse
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Equal(t, cse.expect, resp.List)
	}
}

func TestListAPI(t *testing.T) {
	r, lm := setup()
	lm.Add(1)
	lm.Add(2)
	lm.Add(3)

	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var resp models.ListResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, []float64{1, 2, 3}, resp.List)
}

func TestResetAPI(t *testing.T) {
	r, lm := setup()
	lm.Add(1)
	lm.Add(2)

	req := httptest.NewRequest(http.MethodPost, "/reset", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	req2 := httptest.NewRequest(http.MethodGet, "/list", nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)
	var resp models.ListResponse
	json.Unmarshal(w2.Body.Bytes(), &resp)
	assert.Empty(t, resp.List)
}

func TestInvalidJSON(t *testing.T) {
	r, _ := setup()
	req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString("foo"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func BenchmarkAdd(b *testing.B) {
	lm := listmanager.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lm.Add(float64(i % 100))
	}
}

func BenchmarkReduce(b *testing.B) {
	lm := listmanager.New()
	for i := 1; i <= 100; i++ {
		lm.Add(float64(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lm.Add(float64(-(i%50 + 1)))
	}
}

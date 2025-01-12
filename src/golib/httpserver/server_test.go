package httpserver

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestNewServer(t *testing.T) {
	// Given:.
	ctx := context.Background()

	// When:.
	s, err := NewServer(ctx)

	// Then:.
	require.NoError(t, err)
	require.NotNil(t, s)

	// When:.
	s, err = NewServer(ctx, func(_ *Server, _ *chi.Mux) error {
		return errors.New("custom err")
	})

	// Then:.
	require.Error(t, err)
	require.Nil(t, s)
}

func TestServer_Start(t *testing.T) {
	// Given:.
	ctx := context.Background()

	srv, err := NewServer(ctx)
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	// When:.
	go func() {
		defer wg.Done()
		err = srv.Start(ctx)
	}()

	// Then:.
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
	require.NoError(t, err)
}

func TestWithPort(t *testing.T) {
	// Given:.
	s := &Server{srv: &http.Server{}}

	// When:.
	err := WithPort(-1)(s, nil)

	// Then:.
	require.Equal(t, errors.New("invalid port number"), err)

	// When:.
	err = WithPort(3)(s, nil)
	require.NoError(t, err)
	require.Equal(t, ":3", s.srv.Addr)
}

func TestWithReadTimeout(t *testing.T) {
	// Given:.
	s := &Server{srv: &http.Server{}}

	// When:.
	err := WithReadTimeout(-1)(s, nil)

	// Then:.
	require.Equal(t, errors.New("read timeout must be positive"), err)

	// When:.
	err = WithReadTimeout(3)(s, nil)
	require.NoError(t, err)
	require.Equal(t, time.Duration(3), s.srv.ReadTimeout)
}

func TestWithWriteTimeout(t *testing.T) {
	// Given:.
	s := &Server{srv: &http.Server{}}

	// When:.
	err := WithWriteTimeout(-1)(s, nil)

	// Then:.
	require.Equal(t, errors.New("write timeout must be positive"), err)

	// When:.
	err = WithWriteTimeout(3)(s, nil)
	require.NoError(t, err)
	require.Equal(t, time.Duration(3), s.srv.WriteTimeout)
}

func TestWithIdleTimeout(t *testing.T) {
	// Given:.
	s := &Server{srv: &http.Server{}}

	// When:.
	err := WithIdleTimeout(-1)(s, nil)

	// Then:.
	require.Equal(t, errors.New("idle timeout must be positive"), err)

	// When:.
	err = WithIdleTimeout(3)(s, nil)
	require.NoError(t, err)
	require.Equal(t, time.Duration(3), s.srv.IdleTimeout)
}

func TestWithGracefulShutdownTimeout(t *testing.T) {
	// Given:.
	s := &Server{srv: &http.Server{}}

	// When:.
	err := WithGracefulShutdownTimeout(-1)(s, nil)

	// Then:.
	require.Equal(t, errors.New("graceful shutdown timeout must be positive"), err)

	// When:.
	err = WithGracefulShutdownTimeout(3)(s, nil)
	require.NoError(t, err)
	require.Equal(t, time.Duration(3), s.gracefulShutdownTimeout)
}

func TestWithProfilingHandler(t *testing.T) {
	// Given:.
	m := chi.NewRouter()

	// When:.
	err := WithProfilingHandler()(nil, m)

	// Then:.
	require.NoError(t, err)

	var routesFound []string
	require.NoError(t, chi.Walk(
		m,
		func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
			routesFound = append(routesFound, method+" "+route)
			return nil
		},
	))
	sort.Strings(routesFound)

	routesExp := []string{
		"CONNECT /_/profile/*", "CONNECT /_/profile/allocs", "CONNECT /_/profile/block", "CONNECT /_/profile/cmdline", "CONNECT /_/profile/goroutine", "CONNECT /_/profile/heap", "CONNECT /_/profile/mutex", "CONNECT /_/profile/profile", "CONNECT /_/profile/symbol", "CONNECT /_/profile/threadcreate", "CONNECT /_/profile/trace",
		"DELETE /_/profile/*", "DELETE /_/profile/allocs", "DELETE /_/profile/block", "DELETE /_/profile/cmdline", "DELETE /_/profile/goroutine", "DELETE /_/profile/heap", "DELETE /_/profile/mutex", "DELETE /_/profile/profile", "DELETE /_/profile/symbol", "DELETE /_/profile/threadcreate", "DELETE /_/profile/trace",
		"GET /_/profile/*", "GET /_/profile/allocs", "GET /_/profile/block", "GET /_/profile/cmdline", "GET /_/profile/goroutine", "GET /_/profile/heap", "GET /_/profile/mutex", "GET /_/profile/profile", "GET /_/profile/symbol", "GET /_/profile/threadcreate", "GET /_/profile/trace",
		"HEAD /_/profile/*", "HEAD /_/profile/allocs", "HEAD /_/profile/block", "HEAD /_/profile/cmdline", "HEAD /_/profile/goroutine", "HEAD /_/profile/heap", "HEAD /_/profile/mutex", "HEAD /_/profile/profile", "HEAD /_/profile/symbol", "HEAD /_/profile/threadcreate", "HEAD /_/profile/trace",
		"OPTIONS /_/profile/*", "OPTIONS /_/profile/allocs", "OPTIONS /_/profile/block", "OPTIONS /_/profile/cmdline", "OPTIONS /_/profile/goroutine", "OPTIONS /_/profile/heap", "OPTIONS /_/profile/mutex", "OPTIONS /_/profile/profile", "OPTIONS /_/profile/symbol", "OPTIONS /_/profile/threadcreate", "OPTIONS /_/profile/trace",
		"PATCH /_/profile/*", "PATCH /_/profile/allocs", "PATCH /_/profile/block", "PATCH /_/profile/cmdline", "PATCH /_/profile/goroutine", "PATCH /_/profile/heap", "PATCH /_/profile/mutex", "PATCH /_/profile/profile", "PATCH /_/profile/symbol", "PATCH /_/profile/threadcreate", "PATCH /_/profile/trace",
		"POST /_/profile/*", "POST /_/profile/allocs", "POST /_/profile/block", "POST /_/profile/cmdline", "POST /_/profile/goroutine", "POST /_/profile/heap", "POST /_/profile/mutex", "POST /_/profile/profile", "POST /_/profile/symbol", "POST /_/profile/threadcreate", "POST /_/profile/trace",
		"PUT /_/profile/*", "PUT /_/profile/allocs", "PUT /_/profile/block", "PUT /_/profile/cmdline", "PUT /_/profile/goroutine", "PUT /_/profile/heap", "PUT /_/profile/mutex", "PUT /_/profile/profile", "PUT /_/profile/symbol", "PUT /_/profile/threadcreate", "PUT /_/profile/trace",
		"TRACE /_/profile/*", "TRACE /_/profile/allocs", "TRACE /_/profile/block", "TRACE /_/profile/cmdline", "TRACE /_/profile/goroutine", "TRACE /_/profile/heap", "TRACE /_/profile/mutex", "TRACE /_/profile/profile", "TRACE /_/profile/symbol", "TRACE /_/profile/threadcreate", "TRACE /_/profile/trace",
	}
	sort.Strings(routesExp)

	require.EqualValues(t, routesExp, routesFound)
}

func TestWithReadinessHandler(t *testing.T) {
	// Given:.
	m := chi.NewRouter()

	// When:.
	err := WithReadinessHandler(nil)(nil, m)

	// Then:.
	require.Equal(t, errors.New("readiness handler cannot be nil"), err)

	// When:.
	err = WithReadinessHandler(func(http.ResponseWriter, *http.Request) {})(nil, m)

	// Then:.
	require.NoError(t, err)
	require.NoError(t, chi.Walk(
		m,
		func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
			require.Equal(t, "GET", method)
			require.Equal(t, "/_/ready", route)
			return nil
		},
	))
}

func TestWithRESTHandler(t *testing.T) {
	// Given:.
	m := chi.NewRouter()

	// When:.
	err := WithRESTHandler(nil)(nil, m)

	// Then:.
	require.Equal(t, errors.New("rest handler cannot be nil"), err)

	// When:.
	err = WithRESTHandler(func(rtr chi.Router) {
		rtr.Get("/get", nil)
	})(nil, m)

	// Then:.
	require.NoError(t, err)
	require.NoError(t, chi.Walk(
		m,
		func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
			require.Equal(t, "GET", method)
			require.Equal(t, "/get", route)
			return nil
		},
	))
}

func TestWithGQLHandler(t *testing.T) {
	// Given:.
	m := chi.NewRouter()

	// When:.
	err := WithGQLHandler(nil)(nil, m)

	// Then:.
	require.Equal(t, errors.New("gql handler cannot be nil"), err)

	// When:.
	err = WithGQLHandler(http.NewServeMux())(nil, m)

	// Then:.
	require.NoError(t, err)
	var routesFound []string
	require.NoError(t, chi.Walk(
		m,
		func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
			routesFound = append(routesFound, method+" "+route)
			return nil
		},
	))
	sort.Strings(routesFound)

	routesExp := []string{
		"CONNECT /graph",
		"DELETE /graph",
		"OPTIONS /graph",
		"GET /graph",
		"HEAD /graph",
		"PATCH /graph",
		"POST /graph",
		"PUT /graph",
		"TRACE /graph",
	}
	sort.Strings(routesExp)

	require.EqualValues(t, routesExp, routesFound)
}

func Test_newRouter(t *testing.T) {
	// Given:.
	r := httptest.NewRequest(http.MethodGet, "/_/ping", nil)
	w := httptest.NewRecorder()
	m := newRouter()

	// When:.
	m.ServeHTTP(w, r)

	require.Equal(t, http.StatusOK, w.Result().StatusCode)
}

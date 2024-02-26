package ginrouter

import (
	"github.com/stretchr/testify/require"
	"mini-go-app/internal/util/envconf"
	"mini-go-app/internal/util/zaplog"
	"net/http"
	"net/http/httptest"
	"testing"
)

type test struct {
	path   string
	err    error
	status string
}

func TestRouter(t *testing.T) {
	log := zaplog.New()
	env := envconf.New()
	r := New(env, log)

	path := ""

	ts := httptest.NewServer(r)
	defer ts.Close()

	for _, tst := range [...]test{
		{
			path:   path + "/health",
			err:    nil,
			status: "200 OK",
		},
		{
			path:   path + "/notfound",
			err:    nil,
			status: "404 Not Found",
		},
	} {
		tc := tst
		t.Run(tc.path, func(t *testing.T) {
			res, err := http.Get(ts.URL + tc.path)
			require.Equal(t, tc.err, err)
			require.Equal(t, tc.status, res.Status)
		})
	}
}

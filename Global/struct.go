package Global

import (
	"github.com/BoredTape/httppool"
	"github.com/BoredTape/go-simple-sql"
	"net/http"
)

type spider struct {
	Pool *httppool.Pools
	DB   *go_simple_sql.CONN
	Err  error
}

type Request struct {
	Header    map[string]string
	Url       string
	Method    string
	Form      map[string]string
	Proxy     string
	Cookiejar []*http.Cookie
}
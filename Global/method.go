package Global

import (
	"github.com/BoredTape/httppool"
)

func (s *spider) Open(args Request) *httppool.Respond {
	var request httppool.Request
	request.Method = args.Method
	request.Header = args.Header
	request.Url = args.Url
	request.Form = args.Form
	request.Proxy = args.Proxy
	request.Cookiejar = args.Cookiejar
	return s.Pool.Open(request)
}

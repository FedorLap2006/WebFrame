package WebFrame

import (
	"bufio"
	"net/http"
)

type Context struct {
	RespWriter http.ResponseWriter
	Request    *http.Request
	IO         *bufio.ReadWriter
	Headers    http.Header
	Cookies    []*http.Cookie
	RemoteAddr string
}

func (this *Context) SetCookie(c http.Cookie) {
	http.SetCookie(this.RespWriter, &c)
}

func (this *Context) Redirect(url string, code int) {
	http.Redirect(this.RespWriter, this.Request, url, code)
}

// func (this *Context) GetRawWR() (http.ResponseWriter, *http.Request) {
// 	return this.RespWriter, this.Request
// }

// func UpgradeContext(c *Context) (http.ResponseWriter, *http.Request) {
// 	w := c.RespWriter
// 	r := c.Request
// 	return w, r
// }

// func (this *Context) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// }

type Handler func(*Context)

func HandleHTTP(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// .. get
		rio := bufio.NewReader(r)
		wio := bufio.NewWriter(w)
		c := Context{RespWriter: w, Request: r, IO: bufio.NewReadWriter(rio, wio), Headers: r.Header, Cookies: r.Cookies(), RemoteAddr: r.RemoteAddr}
		h(&c)

		// w, r = UpgradeContext(&c)
	}
}

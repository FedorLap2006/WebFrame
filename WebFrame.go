package WebFrame

import (
	"bufio"
	"io"
	"net/http"
	"html/template"
)

type Context struct {
	RespWriter http.ResponseWriter
	Request    *http.Request
	IO         *bufio.ReadWriter // reader from body and writer to conn
	Headers    http.Header
	Cookies    []*http.Cookie
	RemoteAddr string
}

func (this* Context) GetPage(filename ...string,ldelim string,rdelim string,exec bool) (*template.Template,error){
	tmp,err := template.New("").Delims(ldelim,rdelim).ParseFiles(filename...)
	return tmp,err
}

func (this *Context) WriteIO(b []byte) (int, error) {
	i, err := this.IO.Write(b)
	this.IO.Flush()
	return i, err
}
func (this *Context) WriteByteIO(b byte) error {
	err := this.IO.WriteByte(b)
	this.IO.Flush()
	return err
}
func (this *Context) WriteRuneIO(r rune) (int, error) {
	i, err := this.IO.WriteRune(r)
	this.IO.Flush()
	return i, err
}
func (this *Context) WriteStringIO(s string) (int, error) {
	i, err := this.IO.WriteString(s)
	this.IO.Flush()
	return i, err
}
func (this *Context) WriteToIO(w io.Writer) (int64, error) {
	i, err := this.IO.WriteTo(w)
	this.IO.Flush()
	return i, err
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
		rio := bufio.NewReader(r.Body)
		wio := bufio.NewWriter(w)
		// log.Println("ACTIVE")
		c := Context{RespWriter: w, Request: r, IO: bufio.NewReadWriter(rio, wio), Headers: r.Header, Cookies: r.Cookies(), RemoteAddr: r.RemoteAddr}
		// log.Println(c.IO)
		h(&c)

		// w, r = UpgradeContext(&c)
	}
}

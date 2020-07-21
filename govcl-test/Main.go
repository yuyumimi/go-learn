package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tjfoc/gmsm/sm3"
	"html/template"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)
func main() {
	url:= "http://localhost:8080/view/view"
	exec.Command(`cmd`, `/c`, `start`, url).Start()

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/sign/", makeHandler(signHandler))
	http.ListenAndServe(":8080", nil)
}


type Page struct {
	Title string
	Sign string
	Body  []byte
	SignParam  string
	Token  string
	Base64  string
	Client string
}
func genUUID() string{
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
	return u1.String()
}
func sign(body string) string {
	data := body
	h := sm3.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	hexStringData := hex.EncodeToString(sum)
	fmt.Printf("digest value is: %x\n", hexStringData)
	return hexStringData
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p :=&Page{}
	renderTemplate(w, "view", p)
}

func signHandler(w http.ResponseWriter, r *http.Request, title string) {
	if(title == "get"){
		signParam := r.FormValue("signParam")
		p :=getBodyValue(r)
		sign := sign(signParam)
		p.Sign=sign
		renderTemplate(w, "view", p)
		return
	}else if(title == "base64"){
		base64Handler(w,r)
		return
	}else if(title == "token"){
		uuidHandler(w,r)
		return
	}

}
func base64Handler(res http.ResponseWriter, req *http.Request) {
	client := req.FormValue("client")
	input := []byte(client)
	encodeString := base64.StdEncoding.EncodeToString(input)
	p :=getBodyValue(req)
    p.Base64= encodeString
	renderTemplate(res, "view", p)
}
func uuidHandler(res http.ResponseWriter, req *http.Request) {
	uid := genUUID()
	uid =strings.ReplaceAll(uid,"-","")
	uid =strings.ToUpper(uid)
	p :=getBodyValue(req)
	p.Token= uid
	renderTemplate(res, "view", p)
}
var templates = template.Must(template.ParseFiles( "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func getBodyValue(req *http.Request) *Page {
	sign := req.FormValue("sign")
	signParam := req.FormValue("signParam")
	token := req.FormValue("token")
	base64 := req.FormValue("base64")
	client := req.FormValue("client")
	p := &Page{SignParam:signParam,Sign:sign, Token: token,Base64:base64,
		Client:client}
	return p
}
var validPath = regexp.MustCompile("^/(edit|view|sign)/(" +
	"[a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		m := validPath.FindStringSubmatch(req.URL.Path)
		if m == nil {
			http.NotFound(res, req)
			return
		}
		fn(res, req, m[2])
	}
}


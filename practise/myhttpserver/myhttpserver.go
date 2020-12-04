package myhttpserver

import (
	"fmt"
	"html/template"
	"net/http"
)


func SampleTemplate(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		Title string
		Body string
	}{
		"Template",
		"Sample Tempalte",
	}
	var tmpl = template.Must(template.New("tmpl").Parse(`
<html>
<head>
<title>{{.Title}}</title>
</head>
<body>
	<h1>{{.Body}}</h1>
</body>
</html>
`))

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Printf("Error: SampleTemplate, %s", err)
	}
}


func SampleHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	data, err := fmt.Fprintf(w, "Method: %v\n", r.Method)
	data, err = fmt.Fprintf(w, "Path: %v\n", r.URL.Path)
	data, err = fmt.Fprintf(w, "Scheme: %v\n", r.URL.Scheme)
	data, err = fmt.Fprintf(w, "Hello client\n")
	if err != nil {
		fmt.Printf("Error SamplleHandler: %v, %v", data, err)
	}
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {

	/*cookies := r.Cookies()
	for index, cookie := range cookies {
		fmt.Printf("index:%d cookie:%#v\n", index, cookie)
	}*/
	// 获取cookie
	c, err := r.Cookie("sessionid")
	fmt.Printf("cookie:%#v, err:%v\n", c, err)

	// 设置cookie
	cookie := &http.Cookie{
		Name:   "sessionid",
		Value:  "lkjsdfklsjfklsfdsfdjslf",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}

	http.SetCookie(w, cookie)

	//在具体数据返回之前设置cookie，否则cookie种不上
	w.Write([]byte("hello"))
}

func Startup() {
	http.HandleFunc("/", SampleHandler)
	// http.HandleFunc("/cookie/", IndexHanle)  处理/cookie/*
	http.HandleFunc("/cookie", IndexHandle)
	http.HandleFunc("/template", SampleTemplate)
	err := http.ListenAndServe("0.0.0.0:8181", nil)
	if err != nil {
		fmt.Printf("Failed listen: %v", err)
		return
	}
}

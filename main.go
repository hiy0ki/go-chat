package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
		t.templ.Execute(w, nil)
	})
}

func main() {
	r := newRoom()
	http.Handle("/hoge", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルームを開始します
	go r.run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		<html>
			<head>
				<title>チャット</title>
			</head>
			<body>
				チャットしましょう
			</body>
		</html>
		`))
	})

	// webサーバを開始
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

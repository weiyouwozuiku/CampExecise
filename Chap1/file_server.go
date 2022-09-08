package chap1

import "net/http"

func FileServer() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":9999", nil)
}

package inject_http_handle

import (
	"fmt"
	"github.com/matishsiao/goInfo"
	"net/http"
)

func Init() {
	http.HandleFunc("/inject-http-handle", func(w http.ResponseWriter, r *http.Request) {
		gi := goInfo.GetInfo()
		fmt.Fprintf(w, "Os info %s", gi.String())
	})
}

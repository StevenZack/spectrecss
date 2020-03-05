package spectrecss

import (
	"net/http"
)

func handleCss(w http.ResponseWriter, r *http.Request, str []byte) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(str)
}

func HandleCss(w http.ResponseWriter, r *http.Request) {
	handleCss(w, r, Bytes_SpectreMinCss)
}
func HandleExp(w http.ResponseWriter, r *http.Request) {
	handleCss(w, r, Bytes_SpectreExpMinCss)
}
func HandleIcons(w http.ResponseWriter, r *http.Request) {
	handleCss(w, r, Bytes_SpectreIconsMinCss)
}

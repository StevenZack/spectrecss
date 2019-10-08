package spectrecss

import (
	"net/http"

	"github.com/StevenZack/spectrecss/views"
)

func handleCss(w http.ResponseWriter, r *http.Request, str string) {
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(str))
}

func HandleCss(w http.ResponseWriter, r *http.Request) {
	handleCss(w, r, views.Str_spectre)
}
func HandleExp(w http.ResponseWriter, r *http.Request) {
	handleCss(w, r, views.Str_spectre_exp)
}
func HandleIcons(w http.ResponseWriter, r *http.Request) {
	handleCss(w, r, views.Str_spectre_icons)
}

func GetStr_spectre() string {
	return views.Str_spectre
}

func GetStr_spectre_exp() string {
	return views.Str_spectre_exp
}

func GetStr_spectre_icons() string {
	return views.Str_spectre_icons
}

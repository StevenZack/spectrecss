package spectrecss

import (
	"github.com/StevenZack/fasthttp"
	"github.com/StevenZack/spectrecss/views"
)

func HandleCss(cx *fasthttp.RequestCtx) {
	cx.SetCssHeader()
	cx.SetContentLength(len(views.Str_spectre))
	cx.WriteString(views.Str_spectre)
}
func HandleExp(cx *fasthttp.RequestCtx) {
	cx.SetCssHeader()
	cx.SetContentLength(len(views.Str_spectre_exp))
	cx.WriteString(views.Str_spectre_exp)
}
func HandleIcons(cx *fasthttp.RequestCtx) {
	cx.SetCssHeader()
	cx.SetContentLength(len(views.Str_spectre_icons))
	cx.WriteString(views.Str_spectre_icons)
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

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

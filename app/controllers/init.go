package controllers
import (
	"github.com/revel/revel"
)
func init() {
	revel.InterceptMethod(App.Init, revel.BEFORE)
}
package controllers
import(
	r "github.com/revel/revel"
	//"fmt"
	//"github.com/RangelReale/osin"
	//"github.com/RangelReale/osin/example"
	//"net/http"
	//"net/url"
)
type BaseController struct{

	*r.Controller
	//Server *osin.Server
}

func (c *BaseController) Begin() r.Result {


	/*cfg := osin.NewServerConfig()
	cfg.AllowGetAccessRequest = true
	cfg.AllowClientSecretInParams = true

	c.Server = osin.NewServer(cfg, example.NewTestStorage())*/
	return nil
}
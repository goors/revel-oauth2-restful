package controllers

import (
	"github.com/revel/revel"
	"github.com/RangelReale/osin"
	"fmt"
	"golang.org/x/oauth2"
	"restful/app/helpers"
)
var cfg = osin.NewServerConfig()
var server = osin.NewServer(cfg, helpers.NewStorage())

type App struct {
	*revel.Controller
}

func (c App) Init() revel.Result{

	if c.Controller.Action != "App.Index" &&
	c.Controller.Action != "App.Token" &&
	c.Controller.Action != "App.GetToken"  {

		if c.Session["access_token"] != c.Params.Get("access_token"){
			mp := map[string]interface{}{
				"error":1,
			}
			return c.RenderJson(mp)
		}
	}

	return nil
}
func (c App) Index() revel.Result {

	r := c.Request.Request
	w := c.Response.Out

	resp := server.NewResponse()
	defer resp.Close()

	ar := server.HandleAuthorizeRequest(resp, r)
	ar.Authorized = true
	server.FinishAuthorizeRequest(resp, r, ar)
	osin.OutputJSON(resp, w, r)

	return nil
}

func (c App) Token() revel.Result {

	r := c.Request.Request
	w := c.Response.Out

	resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAccessRequest(resp, r); ar != nil {
		ar.Authorized = true
		server.FinishAccessRequest(resp, r, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	osin.OutputJSON(resp, w, r)

	return nil
}

func (c App) GetToken() revel.Result{

	code := c.Params.Get("code")

	githubConfig := &oauth2.Config{

		ClientID:     "nikola", // change this to yours
		ClientSecret: "nikola",
		RedirectURL:  "http://localhost:14000/appauth", // change this to your webserver adddress
		Scopes:       []string{"user:email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:14000/appauth",
			TokenURL: "http://localhost:14000/token",
		},
	}

	tok, _ := githubConfig.Exchange(oauth2.NoContext, code)

	mp := map[string]interface{}{
		"token":tok.AccessToken,
	}

	c.Session["access_token"] = tok.AccessToken

	return c.RenderJson(mp)
}

func (c App) Example() revel.Result  {
	mp := map[string]interface{}{
		"great":1,
	}

	return c.RenderJson(mp)
}
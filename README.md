# Revel oauth2 protected api with Golang oauth2 server

## Getting Started

Dependencies

<pre>
go get github.com/revel/revel
go get https://github.com/RangelReale/osin
go get golang.org/x/oauth2
</pre>

### Start api with server:

    revel run restful

### How to get token?

Open url <pre style="display: inline !important;">http://localhost:14000/?response_type=code&client_id=1234&redirect_uri=http://localhost:14000/appauth</pre>

You will get json:
<pre>
{
  "token": "bitv8UqUSiSOQ6bWe3TKjA"
}
</pre>

### How to get check if API is protected with token?
Open url <pre style="display: inline !important;">http://localhost:14000/example?access_token=somegarbagehere</pre> and you will see json
<pre>
{
  "error": 1
}
</pre>


### Project is using test storage but You can always use some database since storage is really straightforward.

<pre>
func NewTestStorage() *TestStorage {
	r := &TestStorage{
		clients:   make(map[string]osin.Client),
		authorize: make(map[string]*osin.AuthorizeData),
		access:    make(map[string]*osin.AccessData),
		refresh:   make(map[string]string),
	}

    //you can make storage of mysql database or similar
	r.clients["1234"] = &osin.DefaultClient{
		Id:          "1234",
		Secret:      "aabbccdd",
		RedirectUri: "http://localhost:14000/appauth",
	}

	return r
}
</pre>

I will add mysql storage when I get some free time. This test storage works great for now.
For help please check my profile <a href="https://murphy.rs/nikola">https://murphy.rs/nikola</a> and send me a message.
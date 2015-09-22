# Revel oauth2 protected api with Golang oauth2 server

## Getting Started

Dependencies

<pre>
go get github.com/revel/revel
go get https://github.com/RangelReale/osin
go get golang.org/x/oauth2
</pre>

#### Start api with server:

    revel run restful

#### How to get token?

Open url <pre style="display: inline !important;">http://localhost:14000/?response_type=code&client_id=1234&redirect_uri=http://localhost:14000/appauth</pre>

You will get json:
<pre>
{
  "token": "bitv8UqUSiSOQ6bWe3TKjA"
}
</pre>

#### How to get check if API is protected with token?
Open url <pre style="display: inline !important;">http://localhost:14000/example?access_token=somegarbagehere</pre> and you will see json
<pre>
{
  "error": 1
}
</pre>


#### Project is using mysql storage for client id and client secret.

<pre>
func NewStorage() *Storage {

	r := &Storage{
		clients:   make(map[string]osin.Client),
		authorize: make(map[string]*osin.AuthorizeData),
		access:    make(map[string]*osin.AccessData),
		refresh:   make(map[string]string),
	}

	db, _ := sql.Open("mysql", "oauth2:password@/oauth2?charset=utf8")
	rows, _ := db.Query("SELECT * FROM Oauth2Client")

	for rows.Next() {

		var Id int
		var Client string
		var Secret string
		var RedirectUrl string

		rows.Scan(&Id, &Client, &Secret, &RedirectUrl)

		log.Println(Client)
		log.Println(Secret)
		log.Println(RedirectUrl)
		r.clients[Client] = &osin.DefaultClient{
			Id:          Client,
			Secret:      Secret,
			RedirectUri: RedirectUrl,
		}

		log.Println(r.clients)
	}

	return r
}
</pre>

#### Do not forget to import db file located at sql/oauth2.sql

<pre>

CREATE TABLE `Oauth2Client` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Client` varchar(255) NOT NULL,
  `Secret` varchar(255) NOT NULL,
  `RedirectUrl` varchar(255) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

</pre>

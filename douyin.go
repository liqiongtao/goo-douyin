package goo_douyin

var (
	OAuth *oauth
	User  *user
	Video *video
)

func Init(clientKey, clientSecret string) {
	config := Config{
		clientKey:    clientKey,
		clientSecret: clientSecret,
	}

	OAuth = &oauth{config}
	User = &user{config}
	Video = &video{config}
}

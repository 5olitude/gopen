package shodan

const BaseURL = "https://api.shodan.io"

type client struct {
	apikey string
}

func New(apikey string) *client {
	return &client{apikey: apikey}
}

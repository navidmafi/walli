package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"navidmafi/walli/logger"
	"navidmafi/walli/networking"
	"net/url"
)

type unsplashProvider struct {
	Auth Authentication
	Provider
}

type randomResponse struct {
	Id   string `json:"id"`
	Slug string `json:"slug"`
	Urls struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	}
}

func (u *unsplashProvider) Authenticate(auth Authentication) error {
	if auth.Strategy != u.GetSupportedAuthStrategy() {
		return errors.ErrUnsupported
	}
	u.Auth = auth
	logger.Logger.Debug("Auth set")
	return nil
}

func (u *unsplashProvider) GetRandom(query string) (string, error) {
	queryUrl := fmt.Sprintf("https://api.unsplash.com/photos/random?query=%s", url.QueryEscape(query))
	byteRes := networking.Fetch(u.getNetworkCtx(queryUrl))
	var result *randomResponse = &randomResponse{}
	json.Unmarshal(byteRes, result)
	logger.Logger.Debug(result)
	return result.Urls.Full, nil
}

func (u *unsplashProvider) GetAssetURL(query string) (string, error) {
	url := u.getQueryURL(query)
	result := networking.FetchText(u.getNetworkCtx(url))
	fmt.Sprintln(result)
	return "asdf", nil
}

func (u *unsplashProvider) GetAuthURL() string {
	return "https://unsplash.com/oauth/applications"
}

func (u *unsplashProvider) GetAuthHelp() string {
	return "Create a new demo OAuth app, grab the Access Key and paste it here : "
}

func (u *unsplashProvider) GetSupportedAuthStrategy() AuthStrategy {
	return ClientID
}

func (u *unsplashProvider) getQueryURL(q string) string {
	return fmt.Sprintf("https://api.unsplash.com/photos/random?query=%s", url.QueryEscape(q))
}

func (u *unsplashProvider) getNetworkCtx(url string) networking.RequestCtx {
	ctx := networking.RequestCtx{
		URL: url,
		Headers: map[string]string{
			"authorization": fmt.Sprintf("Client-ID %s", u.Auth.Secret),
		},
	}

	return ctx
}

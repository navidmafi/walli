package providers

const (
	Unsplash ProviderName = "unsplash"
	Pexels   ProviderName = "pexels"
)

type ProviderName string

type AuthStrategy int

const (
	Cookie AuthStrategy = iota
	Bearer
	ClientID
)

type Authentication struct {
	Strategy AuthStrategy
	Secret   string
}

type Provider interface {
	// RequiresAuthorization() bool
	Authenticate(auth Authentication) error
	GetSupportedAuthStrategy() AuthStrategy
	GetAuthURL() string
	getQueryURL(q string) string
	GetAssetURL(q string) (string, error)
	GetRandom(query string) (string, error)
	GetNewest(query string) (string, error)
	GetAuthHelp() string
}

var Providers = map[ProviderName]Provider{
	Unsplash: &unsplashProvider{},
}

func GetAvailable() []string {
	var result []string

	for key := range Providers {
		result = append(result, string(key))
	}

	return result
}

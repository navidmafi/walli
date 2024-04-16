package networking

import (
	"io"
	"navidmafi/walli/logger"
	"net/http"
	"net/textproto"
)

func Fetch(ctx RequestCtx) []byte {

	req, err := http.NewRequest("GET", ctx.URL, nil)
	if err != nil {
		logger.Logger.Fatalf("Error creating request: %s", err)
	}

	if len(ctx.Headers) > 0 {
		for k, v := range ctx.Headers {
			logger.Logger.Debug(textproto.CanonicalMIMEHeaderKey(k))
			logger.Logger.Debug(v)
			req.Header.Set(textproto.CanonicalMIMEHeaderKey(k), v)
		}
	}

	if len((ctx.Cookies)) > 0 {
		for k, v := range ctx.Cookies {
			req.AddCookie(&http.Cookie{
				Name:  k,
				Value: v,
			})
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Logger.Fatalf("Error making request: %s", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.Fatalf("Error reading response: %s", err)
	}

	if resp.StatusCode != 200 {
		logger.Logger.Fatalf("Error: Non 200 response : %v", string(body))
	}

	return body

}

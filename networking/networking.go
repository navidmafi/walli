package networking

import (
	"bytes"
)

type RequestCtx struct {
	URL     string
	Headers map[string]string
	Cookies map[string]string
}

func FetchText(ctx RequestCtx) string {
	byteRes := Fetch(ctx)

	return string(byteRes)
}

func FetchBinary(ctx RequestCtx) bytes.Buffer {
	byteRes := Fetch(ctx)

	return *bytes.NewBuffer(byteRes)
}

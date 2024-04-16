package downloader

import "bytes"

type Downloader interface {
	DownloadToFile(remote string, local string) error
	Download(remote string, local string) (error, bytes.Buffer)
}

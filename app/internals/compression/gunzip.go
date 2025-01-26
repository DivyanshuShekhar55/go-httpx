package compression

import (
	"bytes"
	"compress/gzip"
)

func AddGZip(resBody string) string {
	var gzip_buff bytes.Buffer
	w := gzip.NewWriter(&gzip_buff)

	w.Write([]byte(resBody))
	w.Close()

	compressed_body := gzip_buff.String()
	return compressed_body
}

package property

import "io"

type Asset struct {
	Stream      io.ReadCloser
	FileName    string
	ContentType string
}

// Close ストリームへのアクセスを閉じる
func (a Asset) Close() error {
	return a.Stream.Close()
}

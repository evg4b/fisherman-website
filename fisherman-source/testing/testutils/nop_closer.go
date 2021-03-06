package testutils

import "io"

type nopCloser struct {
	wr io.Writer
}

func (w *nopCloser) Write(p []byte) (n int, err error) {
	return w.wr.Write(p)
}

func (w *nopCloser) Close() error {
	return nil
}

// NopCloser transforms io.Writer to io.WriteCloser.
func NopCloser(wr io.Writer) io.WriteCloser {
	return &nopCloser{wr: wr}
}

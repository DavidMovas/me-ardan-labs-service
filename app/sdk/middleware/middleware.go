package middleware

// Encoder defines behavior that can encode a data model and provide
// the content type for that encoding.
type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

// isError tests if the Encoder has an error inside of it.
func isError(e Encoder) error {
	err, isError := e.(error)
	if isError {
		return err
	}
	return nil
}

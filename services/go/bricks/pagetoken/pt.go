package pagetoken

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

// Encode encodes key to base64 string token.
func Encode(key interface{}) (string, error) {
	if key == nil {
		return "", nil
	}
	b := &bytes.Buffer{}
	e := gob.NewEncoder(b)
	if err := e.Encode(key); err != nil {
		return "", fmt.Errorf("failed to encode key: %w", err)
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

// Decode decodes base64 string token to key and sets the value to PageToken.Key field.
func Decode(token string, target interface{}) error {
	tokenBytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return fmt.Errorf("failed to decode token from base64: %w", err)
	}
	if err = gob.NewDecoder(bytes.NewReader(tokenBytes)).Decode(target); err != nil {
		return fmt.Errorf("failed to decode token to target: %w", err)
	}
	return nil
}

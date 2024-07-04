package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"godevtoy/pkg"
	"strings"
)

type JwtHandler struct {
	pkg.AbsCommonHandler
}

func (h *JwtHandler) InternalHandle(input string) (output string, err error) {
	parts := strings.Split(input, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid token, expected 3 parts got %d", len(parts))
	}

	headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", fmt.Errorf("failed to decode header: %v", err)
	}

	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to decode payload: %v", err)
	}

	var header map[string]interface{}
	var payload map[string]interface{}

	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return "", fmt.Errorf("failed to unmarshal header: %v", err)
	}

	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		return "", fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	m := map[string]any{
		"Header":    header,
		"Payload":   payload,
		"Signature": parts[2],
	}
	resultBytes, _err := json.MarshalIndent(m, "", "  ")
	if _err != nil {
		err = _err
		return
	}
	return string(resultBytes), nil
}

func NewJwtHandler() *JwtHandler {
	h := &JwtHandler{pkg.AbsCommonHandler{}}
	h.IHasparamHandler = h
	return h
}

func init() {
	pkg.Register("jwt", NewJwtHandler())
}

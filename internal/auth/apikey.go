package auth

import (
    "errors"
    "net/http"
    "strings"
)

func GetAPIKey(headers http.Header) (string, error) {
    authHeader := headers.Get("Authorization")
    if authHeader == "" {
        return "", errors.New("no Authorization header")
    }

    parts := strings.SplitN(authHeader, " ", 2)
    if len(parts) != 2 || parts[0] != "ApiKey" {
        return "", errors.New("invalid Authorization header format")
    }

    return strings.TrimSpace(parts[1]), nil
}

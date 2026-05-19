package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type JWTPayload struct {
	Iss  string `json:"iss"`      // Subject
	Name string `json:"username"` // Name
	Iat  int64  `json:"iat"`      // Issued At
	Exp  int64  `json:"exp"`      // Expiration Time
}

// parseJwtToStruct parses JWT into struct, returns success status
func ParseJwtToStruct(token string, payload *JWTPayload) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("jwt parse error: %v\n", r)
		}
	}()

	// Split token
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("jwt parse error: invalid token format")
		return false
	}

	// Get payload part
	base64Url := parts[1]

	// Replace URL-safe base64 characters
	base64Str := strings.ReplaceAll(base64Url, "-", "+")
	base64Str = strings.ReplaceAll(base64Str, "_", "/")

	// Add padding
	switch len(base64Str) % 4 {
	case 2:
		base64Str += "=="
	case 3:
		base64Str += "="
	}

	// Base64 decode
	decoded, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		fmt.Printf("jwt parse error: %v\n", err)
		return false
	}

	// Parse JSON into struct
	err = json.Unmarshal(decoded, payload)
	if err != nil {
		fmt.Printf("jwt parse error: %v\n", err)
		return false
	}

	return true
}

func main() {
	// testToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwYWFzLnBhc3Nwb3J0LmF1dGgiLCJleHAiOjE3NTU2ODU0MDUsImlhdCI6MTc1NTY4MTc0NSwidXNlcm5hbWUiOiJzdW5qaWFxaS50cGFzIiwidHlwZSI6InBlcnNvbl9hY2NvdW50IiwicmVnaW9uIjoiY24iLCJ0cnVzdGVkIjp0cnVlLCJ1dWlkIjoiY2YwZmYyYzctZDhiNS00NTVhLWE3OTUtZDFiNDY5OTU0OTQwIiwic2l0ZSI6Im9ubGluZSIsImJ5dGVjbG91ZF90ZW5hbnRfaWQiOiJieXRlZGFuY2UiLCJieXRlY2xvdWRfdGVuYW50X2lkX29yZyI6ImJ5dGVkYW5jZSIsInNjb3BlIjoiYnl0ZWRhbmNlIiwic2VxdWVuY2UiOiJUZXN0Iiwib3JnYW5pemF0aW9uIjoi5Lqn5ZOB56CU5Y-R5ZKM5bel56iL5p625p6ELeeUn-a0u-acjeWKoS3otKjph4_kv53pmpwt6LSo6YeP5p625p6EIiwid29ya19jb3VudHJ5IjoiQ0hOIiwibG9jYXRpb24iOiJDTiIsImF2YXRhcl91cmwiOiJodHRwczovL3MxLWltZmlsZS5mZWlzaHVjZG4uY29tL3N0YXRpYy1yZXNvdXJjZS92MS92M18wMGsxXzZjZDc2NTA1LWNhMGMtNGYzZi1hOGYxLWE4YjhmYWJmNjU4Z34_aW1hZ2Vfc2l6ZT1ub29wXHUwMDI2Y3V0X3R5cGU9XHUwMDI2cXVhbGl0eT1cdTAwMjZmb3JtYXQ9cG5nXHUwMDI2c3RpY2tlcl9mb3JtYXQ9LndlYnAiLCJlbWFpbCI6InN1bmppYXFpLnRwYXNAYnl0ZWRhbmNlLmNvbSIsImVtcGxveWVlX2lkIjoyNjI2MTIzLCJuZXdfZW1wbG95ZWVfaWQiOjI2MjYxMjN9.NpBri7Ju5MzdC-l9INzAH_cdApfYM8j4JFle6a_8aJMGyM5NvvfRfjUZz-jWUzxaWvL8SEDQEdskSv2U19tROunAzldnzZ95Qt1F0yVBNdb0RLBBth47r4_c7jOww0RJ0SwaJv894IkPY8QQWpHsK-oLENSTTA_xerKIKgcqYGk"
	testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	var payload JWTPayload
	if ParseJwtToStruct(testToken, &payload) {
		fmt.Printf("✅ JWT解析成功:\n")
		fmt.Printf("  Subject: %s\n", payload.Iss)
		fmt.Printf("  username: %s\n", payload.Name)
		fmt.Printf("  Issued At: %s\n", time.Unix(payload.Iat, 0).Format(time.RFC3339))
	} else {
		fmt.Println("❌ JWT解析失败")
	}
}

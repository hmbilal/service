package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"time"
)

const (
	salt = "PROJECT"
)

func main() {
	// prod US
	//accessKey := "2RE@13BkUddLpj&$17s1"
	//secret := "p5xgZgeVeN%SxJfVkVTQ9Z%XKmnvp2yT4fGSutxX"

	// prod EU
	accessKey := "PL&@jjrZ1*Nsj1r8UE8@"
	secret := "zdrwxP#41KgqD5JY@63M1SL7zE#szh#s6r57vecY"

	// staging US
	//accessKey := "a5ca6QNcxYjjC^^^HXvg"
	//secret := "zKfD5dm&d%atmvy^uzEyMUdUv"

	// staging EU
	//accessKey := "%aU@pZL2tjL#@NJ"
	//secret := "MTssaTUCwB&qF&UrJVM6UYszw4TtS4"

	timestamp := flag.Int64("t", time.Now().Unix(), "timestamp")
	flag.Parse()
	fmt.Printf("Authorization: %s:%s\n", accessKey, generateSignature(secret, *timestamp))
	fmt.Println("Timestamp: ", *timestamp)
}

func generateSignature(secret string, timestamp int64) string {
	partOne := hmacSha256(salt, secret)
	partTwo := hmacSha256(string(partOne), strconv.FormatInt(timestamp, 10))
	h := sha256.New()
	_, _ = h.Write(partTwo)

	return hex.EncodeToString(h.Sum(nil))
}

func hmacSha256(message, secret string) []byte {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	_, _ = h.Write([]byte(message))

	return h.Sum(nil)
}

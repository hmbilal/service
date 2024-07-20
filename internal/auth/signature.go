package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"time"
)

const (
	salt                 = "ACTIVITIES"
	allowedTimestampDiff = 60
)

type NordSignatureManager struct {
	salt            string
	allowedTimeDiff int64
}

type SignatureManager interface {
	Verify(project *Project, signature string, timestamp int64) error
}

func NewNordSignatureManager() SignatureManager {
	return &NordSignatureManager{
		salt:            salt,
		allowedTimeDiff: allowedTimestampDiff,
	}
}

func (s *NordSignatureManager) Verify(project *Project, signature string, timestamp int64) error {
	currentTimestamp := time.Now().Unix()
	if timestamp+s.allowedTimeDiff < currentTimestamp ||
		timestamp-s.allowedTimeDiff > currentTimestamp {
		return errors.New("invalid timestamp")
	}

	if signature != s.generateSignature(project, timestamp) {
		return errors.New("invalid signature")
	}

	return nil
}

func (s *NordSignatureManager) generateSignature(project *Project, timestamp int64) string {
	partOne := s.hmacSha256(salt, project.Secret)
	partTwo := s.hmacSha256(string(partOne), strconv.FormatInt(timestamp, 10))
	h := sha256.New()
	_, _ = h.Write(partTwo)

	return hex.EncodeToString(h.Sum(nil))
}

func (s *NordSignatureManager) hmacSha256(message, secret string) []byte {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	_, _ = h.Write([]byte(message))

	return h.Sum(nil)
}

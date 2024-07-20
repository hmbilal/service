package auth

import (
	"errors"
	nativeError "errors"
	"github.com/gofiber/fiber/v2"
	httpLib "github.com/hmbilal/gofiber-start/internal/http"
	"net/http"
	"strconv"
	"strings"
)

const (
	ProjectContext = "project"
)

type ClientCredentialsMiddleware interface {
	BasicAuth() fiber.Handler
}

type clientCredentialsMiddleware struct {
	projectsRepository Repository
	signatureManager   SignatureManager
}

type credentials struct {
	AccessKey, Signature string
}

func NewClientCredentialsMiddleware(
	projectsRepository Repository,
	signatureManager SignatureManager,
) ClientCredentialsMiddleware {
	return &clientCredentialsMiddleware{
		projectsRepository: projectsRepository,
		signatureManager:   signatureManager,
	}
}

func (m clientCredentialsMiddleware) BasicAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := ctx.Get(fiber.HeaderAuthorization)
		timestamp := ctx.Get("Timestamp")
		intTimestamp, err := strconv.ParseInt(timestamp, 10, 64)

		if err != nil || auth == "" || timestamp == "" {
			return ctx.SendStatus(http.StatusUnauthorized)
		}

		decodedCredentials, err := m.decodeHeader(auth)

		if err != nil {
			return ctx.SendStatus(http.StatusUnauthorized)
		}

		project, err := m.projectsRepository.FindOneByAccessKey(decodedCredentials.AccessKey)
		if err != nil {
			var requestErr *httpLib.RequestError
			if ok := nativeError.As(err, &requestErr); ok && requestErr.StatusCode/100 == 4 {
				return ctx.SendStatus(requestErr.StatusCode)
			}

			return ctx.SendStatus(http.StatusUnauthorized)
		}

		if err := m.signatureManager.Verify(project, decodedCredentials.Signature, intTimestamp); err != nil {
			return ctx.SendStatus(http.StatusUnauthorized)
		}

		ctx.Locals(ProjectContext, project.Title)

		return ctx.Next()
	}
}

func (m clientCredentialsMiddleware) decodeHeader(headerValue string) (*credentials, error) {
	pair := strings.SplitN(headerValue, ":", 2)
	if len(pair) != 2 {
		return nil, errors.New("credentials must be a pair")
	}

	return &credentials{
		AccessKey: pair[0],
		Signature: pair[1],
	}, nil
}

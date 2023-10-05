package auth

import (
	"bytes"

	"fincal/pkg/handler"
	"fincal/pkg/models"

	"golang.org/x/crypto/argon2"
)

type Client struct {
	Handler *handler.Query
}

func New(h *handler.Query) *Client {
	return &Client{
		Handler: h,
	}
}

func (c *Client) Login(salt string, user *models.User, clientPassword string) (*models.User, bool, error) {
	var err error

	key := argon2.Key([]byte(clientPassword), []byte(salt), 3, 32*1024, 4, 32)

	//secrets := &models.Secrets{
	//	UsersID:  user.ID,
	//	Password: key,
	//}
	//_, _ = c.Handler.SaveSecrets(secrets)

	secrets, err := c.Handler.GetUserSecrets(user.ID)
	if err != nil {
		return nil, false, err
	}

	if !bytes.Equal(secrets.Password, key) {
		return user, false, nil
	}

	return user, true, nil
}

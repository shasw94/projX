package utils

import (
	"github.com/shasw94/projX/logger"
	"github.com/shasw94/projX/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass []byte) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		logger.Error("Failed to genrate password: ", err)
		return "", errors.Wrap(err, "utils.HashPassword")
	}

	return string(hashed), nil
}

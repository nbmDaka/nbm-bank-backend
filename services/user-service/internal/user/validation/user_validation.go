package validation

import (
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/domain"
)


func ValidateUserID(
	id int64,
) error {

	if id <= 0 {
		return domain.ErrInvalidUserID
	}

	return nil
}
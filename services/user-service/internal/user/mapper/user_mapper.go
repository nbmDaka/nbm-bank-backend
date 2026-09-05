package mapper

import (
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/domain"
	pb "github.com/nbmDaka/nbm-bank-backend/services/user-service/proto/user"
)


func ToProto(
	user *domain.User,
) *pb.User {

	return &pb.User{
		Id: user.ID,
		Email: user.Email,
		FirstName: user.FirstName,
		LastName: user.LastName,
	}
}
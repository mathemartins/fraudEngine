package gapi

import (
	db "github.com/mathemartins/fraudEngine/db/sqlc"
	"github.com/mathemartins/fraudEngine/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUserWatch(userWatch db.UserWatch) *pb.UserWatch {
	return &pb.UserWatch{
		UserId:      userWatch.UserID.String(),
		WatchReason: userWatch.WatchReason,
		CreatedAt:   timestamppb.New(userWatch.CreatedAt.Time),
	}
}

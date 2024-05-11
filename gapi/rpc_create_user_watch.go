package gapi

import (
	"context"

	"github.com/google/uuid"
	db "github.com/mathemartins/fraudEngine/db/sqlc"
	"github.com/mathemartins/fraudEngine/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUserWatch(ctx context.Context, req *pb.CreateUserWatchRequest) (*pb.CreateUserWatchResponse, error) {
	user_uid, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert string to uuid: %s", err)
	}

	arg := db.CreateUserWatchParams{
		UserID:      user_uid,
		WatchReason: req.GetWatchReason(),
	}

	userWatch, err := server.store.CreateUserWatch(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "userwatch already exists: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &pb.CreateUserWatchResponse{
		UserWatch: convertUserWatch(userWatch),
	}
	return rsp, nil
}

package api

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
)

func (o *officeAPI) RemoveOfficeV1(
	ctx context.Context,
	req *pb.RemoveOfficeV1Request,
) (*pb.RemoveOfficeV1Response, error) {

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "RemoveOfficeV1 - invalid argument", "err", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	officeFound, err := o.service.RemoveOffice(ctx, req.GetOfficeId())
	if err != nil {
		logger.ErrorKV(ctx, "RemoveOfficeV1 -- failed", "err", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !officeFound {
		logger.DebugKV(ctx, "RemoveOfficeV1 - office not found", "officeId", req.GetOfficeId())
	} else {
		logger.DebugKV(ctx, "RemoveOfficeV1 - success", "err", err)
	}

	return &pb.RemoveOfficeV1Response{
		Found: officeFound,
	}, nil
}

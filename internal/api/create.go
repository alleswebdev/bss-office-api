package api

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/logger"
	"github.com/ozonmp/bss-office-api/internal/model"
	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *officeAPI) CreateOfficeV1(
	ctx context.Context,
	req *pb.CreateOfficeV1Request,
) (*pb.CreateOfficeV1Response, error) {
	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "CreateOfficeV1 - invalid argument", "err", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	officeID, err := o.service.CreateOffice(ctx, model.Office{
		Name:        req.GetName(),
		Description: req.Description,
	})

	if err != nil {
		logger.ErrorKV(ctx, "CreateOfficeV1 -- failed", "err", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.DebugKV(ctx, "CreateOfficeV1 - success")

	return &pb.CreateOfficeV1Response{
		OfficeId: officeID,
	}, nil
}

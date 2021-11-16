package api

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/logger"
	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *officeAPI) ListOfficesV1(
	ctx context.Context,
	req *pb.ListOfficesV1Request,
) (*pb.ListOfficesV1Response, error) {
	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "ListOfficesV1 - invalid argument", "err", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	items, err := o.service.ListOffices(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		logger.ErrorKV(ctx, "ListOfficesV1 -- failed", "err", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.DebugKV(ctx, "ListOfficesV1 - success")

	pbItems := make([]*pb.Office, 0, len(items))

	for _, item := range items {
		pbItems = append(pbItems, convertBssOfficeToPb(item))
	}

	return &pb.ListOfficesV1Response{
		Items: pbItems,
	}, nil
}

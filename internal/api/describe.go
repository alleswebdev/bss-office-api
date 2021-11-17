package api

import (
	"context"
	"errors"
	"github.com/ozonmp/bss-office-api/internal/logger"
	"github.com/ozonmp/bss-office-api/internal/metrics"
	"github.com/ozonmp/bss-office-api/internal/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
)

func (o *officeAPI) DescribeOfficeV1(
	ctx context.Context,
	req *pb.DescribeOfficeV1Request,
) (*pb.DescribeOfficeV1Response, error) {

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "DescribeOfficeV1 - invalid argument", "err", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	office, err := o.service.DescribeOffice(ctx, req.GetOfficeId())

	if err != nil {
		logger.ErrorKV(ctx, "DescribeOfficeV1 -- failed", "err", err)

		if errors.Is(err, repo.ErrOfficeNotFound) {
			logger.DebugKV(ctx, "DescribeOfficeV1 - office not found", "officeId", req.GetOfficeId())
			metrics.IncTotalNotFound()

			return nil, status.Error(codes.NotFound, "office not found")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	if office == nil {
		logger.DebugKV(ctx, "DescribeOfficeV1 - office is nil", "officeId", req.GetOfficeId())

		return nil, status.Error(codes.Internal, "office is nil")
	}

	logger.DebugKV(ctx, "DescribeOfficeV1 - success")

	return &pb.DescribeOfficeV1Response{
		Value: convertBssOfficeToPb(office),
	}, nil
}

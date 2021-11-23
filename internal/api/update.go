package api

import (
	"context"
	"errors"
	"github.com/ozonmp/bss-office-api/internal/logger"
	"github.com/ozonmp/bss-office-api/internal/metrics"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/ozonmp/bss-office-api/internal/repo"
	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *officeAPI) UpdateOfficeV1(
	ctx context.Context,
	req *pb.UpdateOfficeV1Request,
) (*pb.UpdateOfficeV1Response, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := o.service.UpdateOffice(ctx, req.GetOfficeId(),
		model.Office{
			ID:          req.GetOfficeId(),
			Name:        req.GetName(),
			Description: req.GetDescription(),
		})

	if err != nil {
		logger.ErrorKV(ctx, "UpdateOfficeV1 -- failed", "err", err)

		if errors.Is(err, repo.ErrOfficeNotFound) {
			metrics.IncTotalNotFound()

			return nil, status.Error(codes.NotFound, "office not found")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	metrics.IncTotalCud(model.Updated)

	return &pb.UpdateOfficeV1Response{
		Status: result,
	}, nil
}

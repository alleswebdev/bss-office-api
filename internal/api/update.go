package api

import (
	"context"
	"errors"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/ozonmp/bss-office-api/internal/repo"
	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *officeAPI) UpdateOfficeV1(
	ctx context.Context,
	req *pb.UpdateOfficeV1Request,
) (*pb.UpdateOfficeV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("UpdateOfficeV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := o.service.UpdateOffice(ctx, req.GetOfficeId(),
		model.Office{
			ID:          req.GetOfficeId(),
			Name:        req.GetName(),
			Description: req.GetDescription(),
		})

	if errors.Is(err, repo.ErrOfficeNotFound) {
		log.Debug().Uint64("officeId", req.GetOfficeId()).Msg("office not found")
		totalOfficeNotFound.Inc()

		return nil, status.Error(codes.NotFound, "office not found")
	}

	if err != nil {
		log.Error().Err(err).Msg("UpdateOfficeV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("UpdateOfficeV1 - success")

	return &pb.UpdateOfficeV1Response{
		Status: result,
	}, nil
}

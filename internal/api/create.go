package api

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
)

func (o *officeAPI) CreateOfficeV1(
	ctx context.Context,
	req *pb.CreateOfficeV1Request,
) (*pb.CreateOfficeV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateOfficeV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	officeID, err := o.repo.CreateOffice(ctx, model.Office{
		Name:        req.GetName(),
		Description: req.Description,
	})

	if err != nil {
		log.Error().Err(err).Msg("CreateOfficeV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("CreateOfficeV1 - success")

	return &pb.CreateOfficeV1Response{
		OfficeId: officeID,
	}, nil
}

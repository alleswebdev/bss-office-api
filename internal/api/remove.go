package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
)

func (o *officeAPI) RemoveOfficeV1(
	ctx context.Context,
	req *pb.RemoveOfficeV1Request,
) (*pb.RemoveOfficeV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveOfficeV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := o.repo.RemoveOffice(ctx, req.GetOfficeId())
	if err != nil {
		log.Error().Err(err).Msg("RemoveOfficeV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("RemoveOfficeV1 - success")

	return &pb.RemoveOfficeV1Response{
		Found: result,
	}, nil
}

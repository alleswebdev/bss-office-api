package api

import (
	"context"
	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *officeAPI) ListOfficesV1(
	ctx context.Context,
	req *pb.ListOfficesV1Request,
) (*pb.ListOfficesV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateOfficeV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	items, err := o.repo.ListOffices(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		log.Error().Err(err).Msg("ListOfficesV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("ListOfficesV1 - success")
	pbItems := make([]*pb.Office, 0, len(items))

	for _, item := range items {
		pbItems = append(pbItems, convertBssOfficeToPb(item))
	}

	return &pb.ListOfficesV1Response{
		Items: pbItems,
	}, nil
}

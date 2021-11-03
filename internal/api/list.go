package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
)

func (o *officeAPI) ListOfficesV1(
	ctx context.Context,
	req *pb.ListOfficesV1Request,
) (*pb.ListOfficesV1Response, error) {
	items, err := o.repo.ListOffices(ctx)
	if err != nil {
		log.Error().Err(err).Msg("ListOfficesV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("ListOfficesV1 - success")
	pbItems := make([]*pb.Office, 0, len(items))

	for _, item := range items {
		pbItems = append(pbItems, convertOfficeToPb(&item))
	}

	return &pb.ListOfficesV1Response{
		Items: pbItems,
	}, nil
}

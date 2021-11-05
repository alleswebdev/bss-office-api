package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
)

func (o *officeAPI) DescribeOfficeV1(
	ctx context.Context,
	req *pb.DescribeOfficeV1Request,
) (*pb.DescribeOfficeV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeOfficeV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	office, err := o.repo.DescribeOffice(ctx, req.OfficeId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeOfficeV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if office == nil {
		log.Debug().Uint64("officeId", req.OfficeId).Msg("office not found")
		totalOfficeNotFound.Inc()

		return nil, status.Error(codes.NotFound, "office not found")
	}

	log.Debug().Msg("DescribeOfficeV1 - success")

	return &pb.DescribeOfficeV1Response{
		Value: convertBssOfficeToPb(office),
	}, nil
}

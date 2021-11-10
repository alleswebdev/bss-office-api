package api

import (
	"github.com/ozonmp/bss-office-api/internal/model"
	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertBssOfficeToPb(office *model.Office) *pb.Office {
	return &pb.Office{
		Id:          office.ID,
		Name:        office.Name,
		Description: office.Description,
		Removed:     office.Removed,
		Created:     timestamppb.New(office.Created),
		Updated:     timestamppb.New(office.Updated),
	}
}

func convertPbToBssOffice(pb *pb.Office) *model.Office {
	return &model.Office{
		ID:          pb.GetId(),
		Name:        pb.GetName(),
		Description: pb.GetDescription(),
		Removed:     pb.GetRemoved(),
		Created:     pb.GetCreated().AsTime(),
		Updated:     pb.GetUpdated().AsTime(),
	}
}

package api

import (
	"github.com/ozonmp/bss-office-api/internal/model"
	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
)

func convertBssOfficeToPb(office *model.Office) *pb.Office {
	return &pb.Office{
		Id:          office.ID,
		Name:        office.Name,
		Description: office.Description,
	}
}

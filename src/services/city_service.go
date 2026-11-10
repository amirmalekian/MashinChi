package services

import (
	"context"

	"github.com/amirmalekian/Mashinchi/src/api/dto"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/amirmalekian/Mashinchi/src/data/db"
	"github.com/amirmalekian/Mashinchi/src/data/models"
	"github.com/amirmalekian/Mashinchi/src/pkg/logging"
)

type CityService struct {
	base *BaseService[models.City, dto.CreateCityRequest, dto.UpdateCityRequest, dto.CityResponse]
}

func NewCityService(cfg *config.Config) *CityService {
	return &CityService{
		base: &BaseService[models.City, dto.CreateCityRequest, dto.UpdateCityRequest, dto.CityResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Country"},
			},
		},
	}
}

// Create
func (s *CityService) Create(ctx context.Context, req *dto.CreateCityRequest) (*dto.CityResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CityService) Update(ctx context.Context, id int, req *dto.UpdateCityRequest) (*dto.CityResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CityService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CityService) GetById(ctx context.Context, id int) (*dto.CityResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CityService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CityResponse], error) {
	return s.base.GetByFilter(ctx, req)
}

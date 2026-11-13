package services

import (
	"context"

	"github.com/amirmalekian/Mashinchi/src/api/dto"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/amirmalekian/Mashinchi/src/data/db"
	"github.com/amirmalekian/Mashinchi/src/data/models"
	"github.com/amirmalekian/Mashinchi/src/pkg/logging"
)

type FileService struct {
	base *BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse]
}

func NewFileService(cfg *config.Config) *FileService {
	return &FileService{
		base: &BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *FileService) Create(ctx context.Context, req *dto.CreateFileRequest) (*dto.FileResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *FileService) Update(ctx context.Context, id int, req *dto.UpdateFileRequest) (*dto.FileResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *FileService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *FileService) GetById(ctx context.Context, id int) (*dto.FileResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *FileService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.FileResponse], error) {
	return s.base.GetByFilter(ctx, req)
}

package repository

import (
	"my_gql_server/app/models"
	"my_gql_server/graph/model"
	"time"

	"gorm.io/gorm"
)

type PanRepository interface {
	CreatePan(panInput *model.PanInput) (*models.Pan, error)
	GetOnePan(id string) (*models.Pan, error)
	GetAllPans() ([]*model.Pan, error)
}

type PanService struct {
	Db *gorm.DB
}

var _ PanRepository = &PanService{}

func NewPanService(db *gorm.DB) *PanService {
	return &PanService{
		Db: db,
	}
}

func (b *PanService) CreatePan(panInput *model.PanInput) (*models.Pan, error) {
	layout := "2023-07-06T08:32:22.090Z"
	parsedTime, _ := time.Parse(layout, panInput.CreatedAt)
	pan := &models.Pan{
		ID:        panInput.ID,
		Name:      panInput.Name,
		CreatedAt: parsedTime,
	}
	err := b.Db.Create(&pan).Error

	return pan, err
}

func (b *PanService) GetOnePan(id string) (*models.Pan, error) {
	pan := &models.Pan{}
	err := b.Db.Where("id = ?", id).First(pan).Error
	return pan, err
}

func (b *PanService) GetAllPans() ([]*model.Pan, error) {
	pans := []*model.Pan{}
	err := b.Db.Find(&pans).Error
	return pans, err

}

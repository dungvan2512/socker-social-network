package feed

import (
	"github.com/dungvan/soccer-social-network-api/shared/base"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	SampleUsecase()
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

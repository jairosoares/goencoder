package repositories

import (
	"encoder/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type IVideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

type VideoRepository struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{Db: db}
}

func (repo VideoRepository) Insert(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}
	err := repo.Db.Create(video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (repo VideoRepository) Find(id string) (*domain.Video, error) {
	var video domain.Video
	//Interesting, Preload load Jobs before
	repo.Db.Preload("Jobs").First(&video, "id = ?", id)
	if video.ID =="" {
		return nil, fmt.Errorf("Video does not exist.")
	}
	return &video, nil
}
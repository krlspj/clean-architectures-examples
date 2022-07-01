package repositories

import (
	"time"
	"websocket-sandbox/interfaces"
	"websocket-sandbox/models"

	"errors"

	"github.com/jinzhu/gorm"
)

type NotificationRepository struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) interfaces.NotificationRepository {
	return &NotificationRepository{
		DB: db,
	}
}

func (r *NotificationRepository) Add(userId, project_id int64) error {
	otification := models.Notifications{
		UserID:    userId,
		ProjectID: project_id,
	}
	err := r.DB.Create(&otification).Error
	if err != nil {
		return errors.New("")
	}
	return nil
}

func (r *NotificationRepository) GetNonSend(userId, project_id int64) ([]models.Notifications, error) {
	var notifications []models.Notifications
	if err := r.DB.Where("user_id = ?", userId).Where("project_id = ?", project_id).Where("send_at IS NULL").Find(&notifications).Error; err != nil {
		return []models.Notifications{}, errors.New("")
	} else {
		return notifications, nil
	}
}

func (r *NotificationRepository) GetNonRead(userId, project_id int64) ([]models.Notifications, error) {
	var notifications []models.Notifications
	if err := r.DB.Where("user_id = ?", userId).Where("project_id = ?", project_id).Where("read_at IS NULL").Find(&notifications).Error; err != nil {
		return []models.Notifications{}, errors.New("")
	} else {
		return notifications, nil
	}
}

func (r *NotificationRepository) UpdateSendAt(id int64) error {
	updateAt := time.Now()
	if err := r.DB.Model(&models.Notifications{}).Where("id = ?", id).Update(models.Notifications{SendAt: &updateAt}).Error; err != nil {
		return errors.New("")
	} else {
		return nil
	}
}

func (r *NotificationRepository) UpdateReadAt(id int64) error {
	updateAt := time.Now()
	if err := r.DB.Model(&models.Notifications{}).Where("id = ?", id).Updates(models.Notifications{ReadAt: &updateAt}).Error; err != nil {
		return errors.New("")
	} else {
		return nil
	}
}

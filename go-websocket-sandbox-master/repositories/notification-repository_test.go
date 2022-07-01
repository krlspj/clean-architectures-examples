package repositories_test

import (
	"testing"
	"websocket-sandbox/db"

	"websocket-sandbox/repositories"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

func TestGetNonReadRecord(t *testing.T) {

	var mockUserID int64
	var mockProjectID int64

	faker.FakeData(&mockUserID)
	faker.FakeData(&mockProjectID)

	notificationRepo := repositories.NewNotificationRepository(db.DB())
	err := notificationRepo.Add(mockUserID, mockProjectID)
	assert.NoError(t, err)

	notSendsRecords1, err := notificationRepo.GetNonSend(mockUserID, mockProjectID)
	assert.NoError(t, err)
	assert.Equal(t, len(notSendsRecords1), 1)

	err = notificationRepo.UpdateSendAt(notSendsRecords1[0].ID)
	assert.NoError(t, err)

	notSendsRecords2, err := notificationRepo.GetNonSend(mockUserID, mockProjectID)
	assert.NoError(t, err)
	assert.Equal(t, len(notSendsRecords2), 0)

	notReadsRecords1, err := notificationRepo.GetNonRead(mockUserID, mockProjectID)
	assert.NoError(t, err)
	assert.Equal(t, len(notReadsRecords1), 1)

	err = notificationRepo.UpdateReadAt(notReadsRecords1[0].ID)
	assert.NoError(t, err)

	notReadsRecords2, err := notificationRepo.GetNonRead(mockUserID, mockProjectID)
	assert.NoError(t, err)
	assert.Equal(t, len(notReadsRecords2), 0)
}

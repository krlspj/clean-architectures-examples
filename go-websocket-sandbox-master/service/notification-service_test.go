package service_test

import (
	"testing"

	"websocket-sandbox/service"

	types "websocket-sandbox/types"

	"websocket-sandbox/db"

	"websocket-sandbox/repositories"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

func TestNotification(t *testing.T) {

	notificationRepo := repositories.NewNotificationRepository(db.DB())

	var mockProjectID int64
	var mockClient1 types.Client
	var mockClient2 types.Client

	faker.FakeData(&mockProjectID)
	faker.FakeData(&mockClient1)
	faker.FakeData(&mockClient2)

	notificationService := service.NewNotificationService(notificationRepo)

	// Add mockClient1
	notificationService.AppendeClient(mockClient1)
	assert.Equal(t, len(service.Clients), 1)
	notificationService.AppendeClient(mockClient1)
	assert.Equal(t, len(service.Clients), 1)

	// Add mockClient1 when change projectID
	mockClient1.ProjectID = mockProjectID
	notificationService.AppendeClient(mockClient1)
	assert.Equal(t, len(service.Clients), 1)

	// Add mockClient2
	notificationService.AppendeClient(mockClient2)
	assert.Equal(t, len(service.Clients), 2)
	notificationService.AppendeClient(mockClient1)
	assert.Equal(t, len(service.Clients), 2)

	// Remove mockClient1
	notificationService.RemoveClient(mockClient1)
	assert.Equal(t, len(service.Clients), 1)
	notificationService.RemoveClient(mockClient1)
	assert.Equal(t, len(service.Clients), 1)

	// Remove mockClient2
	notificationService.RemoveClient(mockClient2)
	assert.Equal(t, len(service.Clients), 0)
	notificationService.RemoveClient(mockClient2)
	assert.Equal(t, len(service.Clients), 0)
}

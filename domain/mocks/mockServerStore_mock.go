package mocks

import (
	"bitbucket.org/libertywireless/circles-sandbox/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockServerStoreMock struct{}

func (o MockServerStoreMock) CreateRecord(mockServerId primitive.ObjectID,
	data *domain.MockServerStoreCreateData) (string, error) {
	return "", nil
}

func (o MockServerStoreMock) GetRecord(mockServerId primitive.ObjectID, namespace string) ([]string, error) {
	return nil, nil
}

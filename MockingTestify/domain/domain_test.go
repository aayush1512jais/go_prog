package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aayush1512jais/go_prog/MockingTestify/mocks"
)

//var Err = errors.New("False Error")

func TestDomain_GetUser(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockRepo.On("Get").Return(1, nil)
	testDomain := NewDomain(mockRepo)
	id, err := testDomain.GetUser(1)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, id)
	assert.Nil(t, err)

}

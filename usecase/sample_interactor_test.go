package usecase

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

const testString = "hoge"

func TestExec(t *testing.T) {
	tests := []struct {
		name  string
		setup func(m *MockSampleRepository)
	}{
		{
			name: "Test the argument",
			setup: func(m *MockSampleRepository) {
				m.EXPECT().Read(strings.NewReader(testString))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create gomock controller.
			ctrl := gomock.NewController(t)

			// Create mock repository
			mock := NewMockSampleRepository(ctrl)

			// Setup mock.
			tt.setup(mock)

			// Use mock as sample service.
			sampleService := NewSampleService(mock)

			// Execute target function.
			sampleService.Exec(testString)
		})
	}
}

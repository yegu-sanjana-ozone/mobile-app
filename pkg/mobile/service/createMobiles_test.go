package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	mockMobile "github.com/yegu-sanjana-ozone/mobile-app/mocks"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/model"
	"github.com/stretchr/testify/assert"
)

func TestMobileHandler_CreateMobile(t *testing.T) {

	type args struct {
		mob Model.Mobile
	}
	mockDB := new(mockMobile.Store)


	testCases := []struct {
		name string
		args args
		srvc  *mockMobile.Service 
		wantErr error
	}{
		{
			name: "success",
			srvc: func() *mockMobile.Service {
				mockStore := new(mockMobile.Store)
				mockStore.On("CreateMobile",mock.Anything)
				return  &mockMobile.Service {
					store: mockDB,
				}
			}(),
			args : args {
				mob : Model.Mobile{
					ID : 1,
					Brand: "oneplus",
					Model: "oneplus",
					Year:  "2002",
					Price: "1000000",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.srvc.CreateMobile(tt.args.mob)
			assert.Equal(t, err, tt.wantErr)
		})
	}	

}
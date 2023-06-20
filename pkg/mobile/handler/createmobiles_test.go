package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mockMobile "github.com/yegu-sanjana-ozone/mobile-app/mocks"
	handler "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/handler"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/model"
	service "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/service"
)

func TestMobileHandler_CreateMobile(t *testing.T) {
	testCases := []struct {
		desc          string
		service       service.Service
		expStatusCode int
		path          string
		body          string
		mw            gin.HandlerFunc
	}{
		{
			desc: "success",
			service: func() service.Service {
				mobileService := new(mockMobile.Service)
				mobileService.On("CreateMobile", mock.Anything).Return(nil)
				return mobileService
			}(),
			expStatusCode: 200,
			path:          "/mobile",
			body: `{
				"id": 4,
				"brand":"Google",
				"model":"something ",
				"price":"900000",
				"year":"2022"
			}`,
			mw: gin.HandlerFunc(func(c *gin.Context) {
				c.Set("User", Model.User{
					Email:    "test@test.com",
					Password: "hello",
				})
			}),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			engine := gin.New()
			mH := handler.New()
			engine.POST("/mobile", tC.mw, mH.CreateMobiles)
			testServer := httptest.NewServer(engine)
			defer testServer.Close()
			var body io.Reader
			if tC.body != "" {
				body = strings.NewReader(tC.body)
			}
			req, err := http.NewRequest(http.MethodPost, testServer.URL+tC.path, body)
			if !assert.NoError(t, err, "unexpected error") {
				return
			}
			res, err := http.DefaultClient.Do(req)
			if !assert.NoError(t, err, "unexpected error") {
				return
			}
			assert.Equal(t, tC.expStatusCode, res.StatusCode, "got unexpected statusCode")
		    // tC.service.(*mockMobile.Service).AssertExpectations(t)
		})
	}
}

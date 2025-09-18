package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/LucasKeley/CRUD_Go/src/model"
	"github.com/LucasKeley/CRUD_Go/src/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TesteUserControllerInterface_DeleteUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainInterface(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("userId_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "teste",
			},
		}

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().DeleteUser(id).Return(
			rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().DeleteUser(id).Return(
			model.NewUserDomain(
				"test@test.com",
				"test",
				"test",
				20), nil)

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})

}

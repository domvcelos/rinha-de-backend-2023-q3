package people

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExpectedResponse struct {
	Data  int
	Error error
}

type CountPeopleTestSuit struct {
	suite.Suite
	service *PeopleServiceMock
	sut     *PeopleHandler
}

func (suite *CountPeopleTestSuit) SetupTest() {
	suite.service = NewPeopleServiceMock()
	suite.sut = NewHandler(suite.service)

}
func (s *CountPeopleTestSuit) TestCountPeopleHandler_Should_Retrun_Ok() {
	//Arrange
	ctx := context.Background()
	expectedResponse := ExpectedResponse{Data: 1, Error: nil}
	s.service.On("Count", ctx).Return(expectedResponse.Data, expectedResponse.Error)
	r := httptest.NewRequest("GET", "/contagem-pessoas", nil)
	w := httptest.NewRecorder()
	//Act
	s.sut.Count(w, r)
	resp := w.Result()
	data, _ := io.ReadAll(resp.Body)
	i, _ := strconv.Atoi((strings.Replace(string(data), "\n", "", -1)))
	//Assert
	assert.Equal(s.T(), expectedResponse.Data, i)
	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
}

func (s *CountPeopleTestSuit) TestCountPeopleHandler_Should_Retrun_InternalError() {
	//Arrange
	ctx := context.Background()
	errorMessage := "any_error"
	expectedResponse := ExpectedResponse{Data: 0, Error: errors.New(errorMessage)}
	s.service.On("Count", ctx).Return(expectedResponse.Data, expectedResponse.Error)
	r := httptest.NewRequest("GET", "/contagem-pessoas", nil)
	w := httptest.NewRecorder()
	//Act
	s.sut.Count(w, r)
	resp := w.Result()
	data, _ := io.ReadAll(resp.Body)
	i, _ := strconv.Atoi((strings.Replace(string(data), "\n", "", -1)))
	//Assert
	assert.Equal(s.T(), expectedResponse.Data, i)
	assert.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CountPeopleTestSuit))
}

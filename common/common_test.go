package common

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CommonTestSuite struct {
	suite.Suite
}

func TestCommonTestSuite(t *testing.T) {
	suite.Run(t, new(CommonTestSuite))
}

func (s *CommonTestSuite) TestCredentials() {

}
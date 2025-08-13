package utils

import (
	"encoding/base64"
	"requirements/ent"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const pageSize = 10

func stringReturnStub(t *ent.Requirement) string { return "1234" }

func TestGenerateNextTokenNilReqsEmptyString(t *testing.T) {
	nextToken := GenerateNextToken(nil, stringReturnStub, pageSize)
	assert.Equal(t, "", nextToken)
}

func TestGenerateNextTokenEmptyReqsEmptyString(t *testing.T) {
	reqs := []*ent.Requirement{}
	nextToken := GenerateNextToken(reqs, stringReturnStub, pageSize)
	assert.Equal(t, "", nextToken)
}

func TestGenerateNextTokenPageSizeMoreThanReqLengthEmptyString(t *testing.T) {
	// req length does not meet the need for more pages
	path := "test/test/test"

	req := &ent.Requirement{
		Path: path,
	}
	reqs := []*ent.Requirement{req}
	nextToken := GenerateNextToken(reqs, stringReturnStub, pageSize)
	assert.Equal(t, "", nextToken)
}

func TestGenerateNextTokenPageSizeLessThanReqLengthEmptyString(t *testing.T) {
	req := &ent.Requirement{
		Path: "test/test/test",
	}
	path2 := "test/test/test2"
	req2 := &ent.Requirement{
		Path: path2,
	}

	reqs := []*ent.Requirement{req, req2}

	pageSize := len(reqs) - 1
	nextToken := GenerateNextToken(reqs, func(t *ent.Requirement) string { return path2 }, pageSize)
	decodedNextToken, err := base64.URLEncoding.DecodeString(nextToken)
	require.NoError(t, err)
	assert.Equal(t, path2, string(decodedNextToken))
}

func TestGenerateNextTokenPageSizeEqualToReqLengthEmptyString(t *testing.T) {
	req := &ent.Requirement{
		Path: "test/test/test",
	}

	reqs := []*ent.Requirement{req}
	pageSize := len(reqs)
	nextToken := GenerateNextToken(reqs, stringReturnStub, pageSize)
	assert.Equal(t, "", nextToken)
}

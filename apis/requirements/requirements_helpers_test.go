package requirements

import (
	"encoding/base64"
	"requirements/ent"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateNextTokenNilReqsEmptyString(t *testing.T) {
	nextToken := GenerateNextToken(nil, 10)
	assert.Equal(t, "", nextToken)
}

func TestGenerateNextTokenEmptyReqsEmptyString(t *testing.T) {
	reqs := []*ent.Requirement{}
	nextToken := GenerateNextToken(reqs, 10)
	assert.Equal(t, "", nextToken)
}

func TestGenerateNextTokenPageSizeMoreThanReqLengthEmptyString(t *testing.T) {
	// req length does not meet the need for more pages
	req := &ent.Requirement{
		Path: "test/test/test",
	}
	reqs := []*ent.Requirement{req}
	nextToken := GenerateNextToken(reqs, 10)
	assert.Equal(t, "", nextToken)
}

func TestGenerateNextTokenPageSizeLessThanReqLengthEmptyString(t *testing.T) {
	req := &ent.Requirement{
		Path: "test/test/test",
	}
	req2 := &ent.Requirement{
		Path: "test/test/test2",
	}

	reqs := []*ent.Requirement{req, req2}
	pageSize := len(reqs) - 1
	nextToken := GenerateNextToken(reqs, pageSize)
	decodedNextToken, err := base64.URLEncoding.DecodeString(nextToken)
	require.NoError(t, err)
	assert.Equal(t, req2.Path, string(decodedNextToken))
}

func TestGenerateNextTokenPageSizeEqualToReqLengthEmptyString(t *testing.T) {
	req := &ent.Requirement{
		Path: "test/test/test",
	}

	reqs := []*ent.Requirement{req}
	pageSize := len(reqs)
	nextToken := GenerateNextToken(reqs, pageSize)
	assert.Equal(t, "", nextToken)
}

package requirements

import (
	"encoding/base64"
	"requirements/ent"
)

func GenerateNextToken(reqs []*ent.Requirement, pageSize int) string {
	if len(reqs) <= 0 || len(reqs) <= pageSize {
		return ""
	}

	lastPath := []byte(reqs[len(reqs)-1].Path)
	return base64.URLEncoding.EncodeToString(lastPath)
}

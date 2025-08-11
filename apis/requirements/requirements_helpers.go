package requirements

import (
	"encoding/base64"
)

// func GenerateNextToken(reqs []*ent.Requirement, pageSize int) string {
// 	if len(reqs) <= 0 || len(reqs) <= pageSize {
// 		return ""
// 	}

// 	lastPath := []byte(reqs[len(reqs)-1].Path)
// 	return base64.URLEncoding.EncodeToString(lastPath)
// }

// ideas
// generic items
// somehow communicate the field to the method
// spec the field
// some kind of interface and wrap the struct
// map the item to the required field

func GenerateNextToken(reqs []string, pageSize int) string {
	if len(reqs) <= 0 || len(reqs) <= pageSize {
		return ""
	}

	lastPath := []byte(reqs[len(reqs)-1])
	return base64.URLEncoding.EncodeToString(lastPath)
}

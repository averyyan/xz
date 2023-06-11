package xzutil

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"sort"

	xzcommon "github.com/averyyan/xz/core/common"
	"github.com/google/uuid"
)

// 生成Http请求ID
func GenerateRequestID(req *http.Request) string {
	var requestID string
	if req != nil {
		requestID = req.Header.Get(xzcommon.HeaderRequestId)
	}
	if len(requestID) == 0 {
		requestID = GenerateEventID()
	}

	return requestID
}

// 生成事件ID
func GenerateEventID() string {
	var eventID string
	if randId, err := uuid.NewRandom(); err != nil {
		eventID = ""
	} else {
		eventID = randId.String()
	}
	return eventID
}

// 判断字符串是否在数组内部
func InArrayString(target string, array []string) bool {
	sort.Strings(array)
	index := sort.SearchStrings(array, target)
	if index < len(array) && array[index] == target {
		return true
	}
	return false
}

// 32位md5加密
func MD5(signStr string) string {
	hash := md5.New()
	hash.Write([]byte(signStr))
	return hex.EncodeToString(hash.Sum(nil))
}

package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"crypto/md5"
    "encoding/hex"
)

func PaginationInf(ctx *gin.Context) (int, int, int, string) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	count, _ := strconv.Atoi(ctx.Query("count"))
	sort := ctx.Query("sort")
	if page == 0 {
		page = 1
	}
	switch {
	case size > 100:
		size = 100
	case size < 1:
		size = 20
	}
	if sort == "" {
		sort = "ID"
	}
	offset := (page - 1) * size
	return size, offset, count, sort
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
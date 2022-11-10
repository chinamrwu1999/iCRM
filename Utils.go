package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationInf(ctx *gin.Context) (int, int, int, string) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	count, _ := strconv.Atoi(ctx.Query("count")) //count==0 时需求查询记录总数
	sort := ctx.Query("sort")
	fmt.Printf("page=%d,size=%d\n", page, size)
	//fmt.Printf("size=%d\n", size)
	switch {
	case size > 100:
		size = 100
	case size < 1:
		size = 20
	}
	if sort == "" {
		sort = "ID"
	}

	offset := page * size

	//fmt.Printf("page=%d offsize=%d ", page, offset)
	return size, offset, count, sort
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

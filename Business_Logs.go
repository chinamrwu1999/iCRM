package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
)

type BusinessLog struct {
	ID         uint   `gorm:"column:ID"`
	EmployeeId string `gorm:"column:employeeId"`
	HospitalId int    `gorm:"column:hospitalId"`
	ProxyId    int    `gorm:"column:proxyId"`
	CustomerId int    `gorm:"column:customerId"`
	Content    string `gorm:"column:content"`
}

func AddLogs(c *gin.Context) {
	var obj BusinessLog
	if err := c.BindJSON(&obj); err != nil {
		fmt.Println("添工作日志时，解析参数发生错误")
		fmt.Println(obj)
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}

	if err := db.Create(&obj).Error; err != nil {
		fmt.Println("添加工作日志到数据库失败：", err)
		return
	}

	fmt.Println(obj)
	c.JSON(http.StatusOK, obj)
}

/***************  单个客户信息  ***********************/
func fetchLog(c *gin.Context) {

	logId, err := strconv.Atoi(c.Param("logId"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}

	var obj BusinessLog
	err = db.First(&obj, logId).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, obj)
}

/***************************** 查询  *****************************************/
func QueryEmployeeLogs(c *gin.Context) {
	var objs []map[string]interface{}
	var paras map[string]string

	if err := c.BindJSON(&paras); err != nil {
		fmt.Println("查询销售工作日志：解析参数发生错误")
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	EmployeeId, _ := strconv.Atoi(paras["employeeId"])
	HospitalId, _ := strconv.Atoi(paras["hospitalId"])
	CustomerId, _ := strconv.Atoi(paras["customerId"])
	ProxyId, _ := strconv.Atoi(paras["proxyId"])
	strBeginDate := paras["beginDate"]
	strEndDate := paras["endDate"]

	var sql = `select A.ID,E.Name,H.Name as hospitalName,C.Name as customerName,P.Name as proxyName,content,DATE_FORMAT(workingDate,"%Y-%m-%d") as wdate
	from businesslog A
	left join Hospital H ON A.hospitalId=H.ID
	left join customer C ON A.customerId=C.ID
	left join proxy P ON A.proxyID=P.ID
	left join employee E ON A.employeeId=E.ID
	WHERE A.employeeID=? 
	AND workingDate >= ? AND workingDate <= ?
	ORDER BY ? limit ?,?
	`

	var countSql = `SELECT count(*)  
			FROM businesslog A
			LEFT JOIN Hospital H ON A.hospitalId=H.ID
			LEFT JOIN customer C ON A.customerId=C.ID
			LEFT JOIN proxy P ON A.proxyID=P.ID
			LEFT JOIN employee E ON A.employeeId=E.ID	
			WHERE A.employeeId= ? AND workingDate >= ? AND workingDate <= ?
		`

	frmTime := time.Now().AddDate(0, 0, -30)
	endTime := time.Now()

	if strBeginDate != "" {
		frmTime, err = time.Parse("YYYY-MM-dd", strBeginDate)

		//countSql += "AND workingDate >= ? "
	}

	if strEndDate != "" {
		endTime, err = time.Parse("YYYY-MM-dd", strBeginDate)
		//countSql += "AND workingDate <= ? "
	}

	var pagination Pagination
	var ct int64
	size, offset, count, sort := PaginationInf(c)
	//fmt.Printf("offset=%d\n", offset)
	pagination.PageSize = size
	pagination.StartIndex = offset

	DB := db.Raw(sql, EmployeeId, frmTime, endTime, sort, offset, size) //.Find(&objs).Error
	DB1 := db.Raw(countSql, EmployeeId, frmTime, endTime)

	if HospitalId != 0 {
		//sql += " AND A.hospitalId=? "
		DB.Where("A.hospitalId=?", HospitalId)
		DB1.Where("A.hospitalId=?", HospitalId)
	}
	if CustomerId != 0 {
		//sql += " AND A.customerId=? "
		DB.Where("A.customerId=?", CustomerId)
		DB1.Where("A.customerId=?", CustomerId)
	}
	if ProxyId != 0 {
		sql += " AND A.proxyId=? "
		DB.Where("A.proxyId=?", ProxyId)
		DB1.Where("A.proxyId=?", ProxyId)
	}

	err = DB.Find(&objs).Error

	pagination.Rows = objs

	if count == 0 {
		DB1.Count((&ct))
		fmt.Print(ct)
	}
	pagination.StartIndex = 0
	pagination.TotalRows = ct
	pagination.TotalPages = int(math.Ceil(float64(ct) / float64(pagination.PageSize)))

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, pagination)
}

func QueryLogs(c *gin.Context) {
	var objs []map[string]interface{}
	var paras map[string]string

	if err := c.BindJSON(&paras); err != nil {
		fmt.Println("查询销售工作日志：解析参数发生错误")
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	id := paras["id"]
	intId, _ := strconv.Atoi(id)
	targetId, _ := paras["target"]
	strBeginDate := paras["date1"]
	strEndDate := paras["date2"]

	// fmt.Print("===>id=" + id + "target=" + targetId + ",date1=" + strBeginDate + ",date2=" + strEndDate + "   <<<\n")

	var sql = `select A.ID,E.Name,H.Name as hospitalName,C.Name as customerName,P.Name as proxyName,content,DATE_FORMAT(workingDate,"%Y-%m-%d") as wdate
	from businesslog A
	left join Hospital H ON A.hospitalId=H.ID
	left join customer C ON A.customerId=C.ID
	left join proxy P ON A.proxyID=P.ID
	left join employee E ON A.employeeId=E.ID
	WHERE  workingDate >= ? AND workingDate <= ?
	AND A.employeeId in 
	`
	var countSql = `SELECT count(*)  
			FROM businesslog A
			LEFT JOIN Hospital H ON A.hospitalId=H.ID
			LEFT JOIN customer C ON A.customerId=C.ID
			LEFT JOIN proxy P ON A.proxyID=P.ID
			LEFT JOIN employee E ON A.employeeId=E.ID	
			WHERE  workingDate >= ? AND workingDate <= ?
		`

	frmTime := time.Now().AddDate(0, 0, -10)
	endTime := time.Now()

	if strBeginDate != "" {
		// fmt.Println(strBeginDate)
		frmTime, err = time.Parse("2006-01-02 15:04", strBeginDate+" 00:01")

		//countSql += "AND workingDate >= ? "
	}

	if strEndDate != "" {
		// fmt.Println(strEndDate)
		endTime, err = time.Parse("2006-01-02 15:04", strEndDate+" 11:59")

		//countSql += "AND workingDate <= ? "
	}

	var pagination Pagination
	var ct int64
	size, offset, count, _ := PaginationInf(c)
	//fmt.Printf("offset=%d\n", offset)
	pagination.PageSize = size
	pagination.StartIndex = offset

	// DB := db.Raw(sql, frmTime, endTime, offset, size) //.Find(&objs).Error
	// DB1 := db.Raw(countSql, frmTime, endTime)

	if targetId == "employeeId" {
		fmt.Println("target=" + targetId + " id=" + id)
		sql += " AND A.employeeId=?  ORDER By workingDate desc limit ?,?"
		countSql += fmt.Sprintf(" AND A.employeeId='%s'", id)
		err = db.Raw(sql, frmTime, endTime, id, offset, size).Find(&objs).Error

	} else if targetId == "hospitalId" {
		fmt.Println("target=" + targetId + " id=" + id)
		sql += " AND A.hospitalId=?  ORDER By workingDate desc limit ?,?"
		countSql += fmt.Sprintf(" AND A.hospitalId=%d", intId)
		err = db.Raw(sql, frmTime, endTime, intId, offset, size).Find(&objs).Error

	} else if targetId == "proxyId" {
		fmt.Println("target=" + targetId + " id=" + id)
		sql += " AND A.proxyId=? ORDER By workingDate desc limit ?,?"
		countSql += fmt.Sprintf(" AND A.proxyId=%d", intId)
		err = db.Raw(sql, frmTime, endTime, intId, offset, size).Find(&objs).Error
	} else {

	}

	pagination.Rows = objs

	if count == 0 {
		db.Raw(countSql, frmTime, endTime).Count((&ct))
		fmt.Print(ct)

		pagination.StartIndex = 0
		pagination.TotalRows = ct
		pagination.TotalPages = int(math.Ceil(float64(ct) / float64(pagination.PageSize)))
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, pagination)
}

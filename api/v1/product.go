package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gomaster/common"
	"gomaster/model"
	"net/http"
	"strconv"
)

var (
	sql1 = "select p.* from product p where p.id =? "
)

// @Tags Product
// @Summary 产品信息
// @Produce  application/json
// @Param id path int true "产品信息id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /product/{id} [get]
func Product(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.Log().Error("参数不合法:",err)
		return
	}
	product := &model.Product{}
	err = common.Db().Raw(sql1, id).Scan(&product).Error
	if err != nil {
		common.Log().Error("查询出错 :",err)
		return
	}
	bs , err := json.Marshal(product)
	if err != nil {
		common.Log().Error("解析出错 :",err)
		return
	}
	c.String(http.StatusOK,string(bs))
}

package v1

import "github.com/gogf/gf/net/ghttp"

// @Tags DictionaryDetail
// @Summary 创建DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "创建DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DictionaryDetail/createDictionaryDetail [post]
func CreateDictionaryDetail(r *ghttp.Request) {
}

// @Tags DictionaryDetail
// @Summary 删除DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "删除DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DictionaryDetail/deleteDictionaryDetail [delete]
func DeleteDictionaryDetail(r *ghttp.Request) {
}

// @Tags DictionaryDetail
// @Summary 更新DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "更新DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DictionaryDetail/updateDictionaryDetail [put]
func UpdateDictionaryDetail(r *ghttp.Request) {
}

// @Tags DictionaryDetail
// @Summary 用id查询DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "用id查询DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DictionaryDetail/findDictionaryDetail [get]
func FindDictionaryDetail(r *ghttp.Request) {
}

// @Tags DictionaryDetail
// @Summary 分页获取DictionaryDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DictionaryDetailSearch true "分页获取DictionaryDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DictionaryDetail/getDictionaryDetailList [get]
func GetDictionaryDetailList(r *ghttp.Request) {
}

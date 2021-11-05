package management

import (
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
)

// ListAuditLogs
// 审计日志列表查询
func (c *Client) ListAuditLogs(req *model.ListAuditLogsRequest) (*struct {
	List       []interface{} `json:"list"`
	TotalCount int64         `json:"totalCount"`
}, error) {

	if req.UserIds != nil {

		var formatUserIds = make([]string, 0)
		for _, d := range *req.UserIds {
			formatUserId := "arn:cn:authing:user:" + d
			formatUserIds = append(formatUserIds, formatUserId)
		}
		req.UserIds = &formatUserIds
	}
	vars := make(map[string]interface{})
	url := fmt.Sprintf("%s/api/v2/analysis/audit", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
		Data    struct {
			List       []interface{} `json:"list"`
			TotalCount int64         `json:"totalCount"`
		} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// ListUserAction
// 查看用户操作日志
func (c *Client) ListUserAction(req *model.ListUserActionRequest) (*struct {
	List       []interface{} `json:"list"`
	TotalCount int64         `json:"totalCount"`
}, error) {

	if req.UserIds != nil {

		var formatUserIds = make([]string, 0)
		for _, d := range *req.UserIds {
			formatUserId := "arn:cn:authing:user:" + d
			formatUserIds = append(formatUserIds, formatUserId)
		}
		req.UserIds = &formatUserIds
	}
	vars := make(map[string]interface{})
	url := fmt.Sprintf("%s/api/v2/analysis/user-action", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
		Data    struct {
			List       []interface{} `json:"list"`
			TotalCount int64         `json:"totalCount"`
		} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

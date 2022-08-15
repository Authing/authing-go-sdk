package management

import (
	"encoding/json"
	"errors"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

// ListUdf
// 获取自定义字段定义
func (c *Client) ListUdf(targetType model.EnumUDFTargetType) (*[]model.UserDefinedField, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListUdfDocument,
		map[string]interface{}{"targetType": targetType})
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			Udf []model.UserDefinedField `json:"udf"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Udf, nil
}

// SetUdf
// 设置自定义字段元数据
func (c *Client) SetUdf(req *model.SetUdfInput) (*model.UserDefinedField, error) {
	data, _ := jsoniter.Marshal(req)
	vars := make(map[string]interface{})
	jsoniter.Unmarshal(data, &vars)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.SetUdfDocument, vars)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			SetUdf model.UserDefinedField `json:"setUdf"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.SetUdf, nil
}

// RemoveUdf
// 删除自定义字段
func (c *Client) RemoveUdf(targetType model.EnumUDFTargetType, key string) (*model.CommonMessageAndCode, error) {
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RemoveUdfDocument, map[string]interface{}{
		"targetType": targetType,
		"key":        key,
	})
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			RemoveUdf model.CommonMessageAndCode `json:"removeUdf"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.RemoveUdf, nil
}

// ListUdfValue
// 获取某一实体的自定义字段数据列表
func (c *Client) ListUdfValue(targetType model.EnumUDFTargetType, targetId string) (*[]model.UserDefinedData, error) {
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UdvDocument, map[string]interface{}{
		"targetType": targetType,
		"targetId":   targetId,
	})
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			Udv []model.UserDefinedData `json:"udv"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Udv, nil
}

// SetUdvBatch
// 批量添加自定义数据
func (c *Client) SetUdvBatch(id string, targetType model.EnumUDFTargetType, udv *[]model.KeyValuePair) (*[]model.UserDefinedData, error) {
	variables := make(map[string]interface{})

	variables["targetType"] = targetType
	variables["targetId"] = id
	var reqUdv []model.KeyValuePair
	for _, v := range *udv {
		v1, _ := json.Marshal(&v.Value)
		v.Value = string(v1)
		reqUdv = append(reqUdv, v)
	}
	variables["udvList"] = reqUdv

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.SetRoleUdfValueDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			SetUdvBatch []model.UserDefinedData `json:"setUdvBatch"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.SetUdvBatch, nil
}

package management

import (
	"errors"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
)

//GetWhileList
//获取白名单记录
func (c *Client) GetWhileList(whileListType model.EnumWhitelistType) (*[]model.WhiteList, error) {
	//var req model.UpdateUserpoolInput
	//if whileListType == model.EnumWhitelistTypeUSERNAME {
	//	req.Whitelist.EmailEnabled = true
	//}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.WhileListDocument, map[string]interface{}{
		"type": whileListType,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			WhileList []model.WhiteList `json:"whitelist"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.WhileList, nil
}

//AddWhileList
//添加白名单记录
func (c *Client) AddWhileList(whileListType model.EnumWhitelistType, ids []string) (*[]model.WhiteList, error) {
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.AddWhileListDocument, map[string]interface{}{
		"type": whileListType,
		"list": ids,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			WhileList []model.WhiteList `json:"addWhitelist"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.WhileList, nil
}

//RemoveWhileList
//移除白名单记录
func (c *Client) RemoveWhileList(whileListType model.EnumWhitelistType, ids []string) (*[]model.WhiteList, error) {
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RemoveWhileListDocument, map[string]interface{}{
		"type": whileListType,
		"list": ids,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			WhileList []model.WhiteList `json:"removeWhitelist"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.WhileList, nil
}

//EnableWhileList
//开启白名单
func (c *Client) EnableWhileList(whileListType model.EnumWhitelistType) (*model.UserPool, error) {
	var req model.UpdateUserpoolInput
	enable := true
	if whileListType == model.EnumWhitelistTypeUSERNAME {
		req = model.UpdateUserpoolInput{
			Whitelist: &model.RegisterWhiteListConfigInput{
				UsernameEnabled: &enable,
			},
		}
	}

	if whileListType == model.EnumWhitelistTypeEMAIL {
		req = model.UpdateUserpoolInput{
			Whitelist: &model.RegisterWhiteListConfigInput{
				EmailEnabled: &enable,
			},
		}
	}

	if whileListType == model.EnumWhitelistTypePHONE {
		req = model.UpdateUserpoolInput{
			Whitelist: &model.RegisterWhiteListConfigInput{
				PhoneEnabled: &enable,
			},
		}
	}
	rep, err := c.UpdateUserPool(req)
	return rep, err
}

//DisableWhileList
//关闭白名单
func (c *Client) DisableWhileList(whileListType model.EnumWhitelistType) (*model.UserPool, error) {
	var req model.UpdateUserpoolInput
	flag := false
	if whileListType == model.EnumWhitelistTypeUSERNAME {
		req = model.UpdateUserpoolInput{
			Whitelist: &model.RegisterWhiteListConfigInput{
				UsernameEnabled: &flag,
			},
		}
	}

	if whileListType == model.EnumWhitelistTypeEMAIL {
		req = model.UpdateUserpoolInput{
			Whitelist: &model.RegisterWhiteListConfigInput{
				EmailEnabled: &flag,
			},
		}
	}

	if whileListType == model.EnumWhitelistTypePHONE {
		req = model.UpdateUserpoolInput{
			Whitelist: &model.RegisterWhiteListConfigInput{
				PhoneEnabled: &flag,
			},
		}
	}
	rep, err := c.UpdateUserPool(req)
	return rep, err
}

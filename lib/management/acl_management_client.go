package management

import (
	"encoding/json"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	"github.com/bitly/go-simplejson"
	"log"
)

func (c *Client) IsAllowed(request model.IsAllowedRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.IsActionAllowedDocument, variables)
	if err != nil {
		return false, err
	}
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("data").Get("isActionAllowed").Bool()
	if err != nil {
		return false, err
	}
	return result, nil
}

func (c *Client) Allow(request model.AllowRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.AllowDocument, variables)
	if err != nil {
		return false, err
	}
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("data").Get("isActionAllowed").Bool()
	if err != nil {
		return false, err
	}
	return result, nil

}

func (c *Client) AuthorizeResource(request model.AuthorizeResourceRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.AuthorizeResourceDocument, variables)
	if err != nil {
		return false, err
	}
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("data").Get("authorizeResource").Get("code").Int64()
	if err != nil {
		return false, err
	}
	return result == 200, nil
}

func (c *Client) RevokeResource(request model.RevokeResourceRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+"/api/v2/acl/revoke-resource", constant.HttpMethodPost, constant.StringEmpty, variables)
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("code").Int64()
	if err != nil {
		return false, err
	}
	return result == 200, nil
}

func (c *Client) CheckResourcePermissionBatch(request model.CheckResourcePermissionBatchRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+"/api/v2/acl/check-resource-permission-batch", constant.HttpMethodPost, constant.StringEmpty, variables)
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("code").Int64()
	if err != nil {
		return false, err
	}
	return result == 200, nil
}

func (c *Client) GetAuthorizedResourcesOfResourceKind(request model.GetAuthorizedResourcesOfResourceKindRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+"/api/v2/acl/get-authorized-resources-of-resource-kind", constant.HttpMethodPost, constant.StringEmpty, variables)
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("code").Int64()
	if err != nil {
		return false, err
	}
	return result == 200, nil
}

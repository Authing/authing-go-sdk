package management

import (
	"encoding/json"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
)

func (c *Client) Detail(userId string) (*model.User, error) {
	b, err := c.SendHttpRequest(c.Host+"/api/v2/users/"+userId, "GET", "", nil)
	if err != nil {
		return nil, err
	}
	var userDetail model.UserDetailResponse
	jsoniter.Unmarshal(b, &userDetail)
	return &userDetail.Data, nil
}

func (c *Client) GetUserList(request model.QueryListRequest) (*model.PaginatedUsers, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.UsersDocument, variables)
	if err != nil {
		return nil, err
	}
	result := model.ListUserResponse{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data.Users, nil
}

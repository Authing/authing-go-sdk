package management

import (
	"authing-golang-sdk/lib/model"
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

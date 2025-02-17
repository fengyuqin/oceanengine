package keyword

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/keyword"
)

// Get 获取关键词列表
// 根据过滤条件获取符合条件的所有关键词。
// 目前仅支持根据ad_id获取该计划下的关键词。
func Get(clt *core.SDKClient, accessToken string, req *keyword.GetRequest) ([]keyword.Keyword, error) {
	var resp keyword.GetResponse
	err := clt.Get("2/keyword/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.List, nil
}

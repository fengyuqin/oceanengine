package advertiser

import (
	"net/url"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// FundDailyStatRequest 查询账户日流水 API Request
type FundDailyStatRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 开始时间，格式YYYY-MM-DD，默认当前年份的1月1日
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate 结束时间，格式YYYY-MM-DD，默认为今天
	EndDate time.Time `json:"end_date,omitempty"`
	// Page 页码. 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量. 默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r FundDailyStatRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if !r.StartDate.IsZero() {
		values.Set("start_date", r.StartDate.Format("2006-01-02"))
	}
	if !r.EndDate.IsZero() {
		values.Set("end_date", r.EndDate.Format("2006-01-02"))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// FundDailyStatResponse 查询账户日流水 API Response
type FundDailyStatResponse struct {
	model.BaseResponse
	// Date json返回值
	Data *FundDailyStatResponseData `json:"data,omitempty"`
}

// FundDailyStatResponseData json返回值
type FundDailyStatResponseData struct {
	// List list
	List []FundDailyStatResponseList `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// FundDailyStatResponseList 广告主流水
type FundDailyStatResponseList struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Date 日期
	Date string `json:"date,omitempty"`
	// Balance 日终结余(单位元）
	Balance float64 `json:"balance,omitempty"`
	// CashCost 现金支出(单位元)
	CashCost float64 `json:"cash_cost,omitempty"`
	// Cost 总支出(单位元)
	Cost float64 `json:"cost,omitempty"`
	// Frozen 冻结(单位元)
	Frozen float64 `json:"frozen,omitempty"`
	// Income 总存入(单位元)
	Income float64 `json:"income,omitempty"`
	// RewardCost 赠款支出(单位元)
	RewardCost float64 `json:"reward_cost,omitempty"`
	// ReturnGoodsCost 返货支出(单位元)
	ReturnGoodsCost float64 `json:"return_goods_cost,omitempty"`
	// TransferIn 总转入(单位元)
	TransferIn float64 `json:"transfer_in,omitempty"`
	// TransferOut 总转出(单位元)
	TransferOut float64 `json:"transfer_out,omitempty"`
}

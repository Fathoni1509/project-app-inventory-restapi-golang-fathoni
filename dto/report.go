package dto

type ReportResponse struct {
	Name      string  `json:"name"`
	Remain    int     `json:"remain"`
	Sold      int     `json:"sold"`
	SellPrice float32 `json:"sell_price"`
	Sales     float32 `json:"sales"`
	Income    float32 `json:"income"`
}

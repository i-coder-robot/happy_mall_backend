package query

type ListQuery struct {
	PageSize int `json:"pageSize"`
	Page     int `json:"page"`
}

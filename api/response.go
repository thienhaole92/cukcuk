package api

type Response[T any] struct {
	Code    int  `json:"Code"`
	Data    T    `json:"Data"`
	Total   int  `json:"Total"`
	Success bool `json:"Success"`
}

type responseCommon[T any] struct {
	Code         int    `json:"Code"`
	Data         T      `json:"Data"`
	Total        int    `json:"Total"`
	Success      bool   `json:"Success"`
	ErrorType    int    `json:"ErrorType"`
	ErrorMessage string `json:"ErrorMessage"`
}

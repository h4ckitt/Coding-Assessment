package helpers

type success struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

type failure struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Error  string `json:"error"`
}

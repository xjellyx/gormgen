package example

type fieldData struct {
	Value  interface{} `json:"value" form:"value"`
	Symbol string      `json:"symbol" form:"symbol"`
}

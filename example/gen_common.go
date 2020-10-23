package example

type FieldData struct {
	Value  interface{} `json:"value" form:"value"`
	Symbol string      `json:"symbol" form:"symbol"`
}

package response

type Error struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Fields  []Field `json:"fields"`
}

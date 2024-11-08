package response

type ContentBlock struct {
	Topic      string        `json:"topic"`
	Content    LocalizedText `json:"content"`
	References References    `json:"references"`
}

package models

type GetCardRequest struct {
	Index string `json:"index"`
	ID    int    `json:"id"`
}

var GetCardRequestLayout = `{
    "index": "%s",
    "query": {
        "equals": {
            "id": %d
        }
    }
}`

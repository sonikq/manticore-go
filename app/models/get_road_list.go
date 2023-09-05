package models

type GetListRequest struct {
	Index    string `json:"index"`
	FullName string `json:"full_name" binding:"required"`
}

var GetListRequestLayout = `{
    "index": "%s",
    "query": {
        "match": {
            "full_name": "%s"
        }
    }
}`

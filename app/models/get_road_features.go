package models

type GetFeaturesRequest struct {
	Index       string    `json:"index"`
	EditMode    int       `json:"edit_mode"`
	Zoom        int       `json:"p_zoom"`
	ScaleFactor float64   `json:"p_scale_factor"`
	Points      []float64 `json:"p_box"`
}

/*const (
	IdxWFSRaw  = "idx_wfs_road_raw"
	IdxWFS10   = "idx_wfs_road_10"
	IdxWFS100  = "idx_wfs_road_100"
	IdxWFS1000 = "idx_wfs_road_1000"
)*/

var GetFeaturesRequestLayout = `{
	"index": "%s",
	"_source": "geom*",
	"limit": %d,
	"max_matches": %d,
    "query": {
        "bool": {
            "should": [
                {
                    "bool": {
                        "must": [
                          {"range":{"zoom":{"lte":%d}}},
                            {
                                "range": {
                                    "xmin": {
                                        "gte": %f,
                                        "lte": %f
                                    }
                                }
                            },
                            {
                                "range": {
                                    "ymin": {
                                        "gte": %f,
                                        "lte": %f
                                    }
                                }
                            }
                        ]
                    }
                },
                {
                    "bool": {
                        "must": [
                          {"range":{"zoom":{"lte":%d}}},
                            {
                                "range": {
                                    "xmax": {
                                        "gte": %f,
                                        "lte": %f
                                    }
                                }
                            },
                            {
                                "range": {
                                    "ymin": {
                                        "gte": %f,
                                        "lte": %f
                                    }
                                }
                            }
                        ]
                    }
                },
                {
                    "bool": {
                        "must": [
                          {"range":{"zoom":{"lte":%d}}},
                            {
                                "range": {
                                    "xmin": {
                                        "gte": %f,
                                        "lte": %f
                                    }
                                }
                            },
                            {
                                "range": {
                                    "ymax": {
                                        "gte": %f,
                                        "lte": %f
                                    }
                                }
                            }
                        ]
                    }
                },            
                {
                    "bool": {
                        "must": [
                          {"range":{"zoom":{"lte":%d}}},
                            {
                                "range": {
                                    "xmax": {
                                        "gte": %f,
                                        "lte": %f
                                    }
                                }
                            },
                            {
                                "range": {
                                    "ymax": {
                                        "gte": %f,
                                        "lte": %f
                                    }
                                }
                            }
                        ]
                    }
                }                                                                                
            ]
        }
    }
}`

package models

import "github.com/crawlab-team/crawlab/core/models/models/v2"

type PermissionV2 struct {
	any                              `collection:"permissions"`
	models.BaseModelV2[PermissionV2] `bson:",inline"`
	Key                              string   `json:"key" bson:"key"`
	Name                             string   `json:"name" bson:"name"`
	Description                      string   `json:"description" bson:"description"`
	Type                             string   `json:"type" bson:"type"`
	Target                           []string `json:"target" bson:"target"`
	Allow                            []string `json:"allow" bson:"allow"`
	Deny                             []string `json:"deny" bson:"deny"`
}

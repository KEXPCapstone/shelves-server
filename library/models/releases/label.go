package releases

// Label represents a label in the KEXP library, with additional
// information about associated releases
type Label struct {
	LabelMBID      string        `json:"labelsMBID" bson:"_id"`
	LabelName      string        `json:"name" bson:"labelName"`
	LabelSortName  string        `json:"labelSortName" bson:"labelSortName"`
	Disambiguation string        `json:"disambiguation" bson:"disambiguation"`
	Releases       []interface{} `json:"releases" bson:"releases"`
}

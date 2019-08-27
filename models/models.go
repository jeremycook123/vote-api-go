package models

type Codedetail struct {
	Usecase  string `json:"usecase,omitempty" bson:"usecase"`
	Rank     int    `json:"rank,omitempty" bson:"rank"`
	Compiled bool   `json:"compiled" bson:"compiled"`
	Homepage string `json:"homepage,omitempty" bson:"homepage"`
	Download string `json:"download,omitempty" bson:"download"`
	Votes    int    `json:"votes" bson:"votes"`
}

type Language struct {
	Name   string     `json:"name,omitempty" bson:"name"`
	Detail Codedetail `json:"codedetail,omitempty" bson:"codedetail"`
}

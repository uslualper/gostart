package log

type Param struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
}

type Context struct {
	Trace   string  `bson:"trace"`
	Message string  `bson:"message"`
	Params  []Param `bson:"params,omitempty"`
}

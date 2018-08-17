package bitbu

//DataBit is the lowest level data exchange from and to the application
//used for db operations like create,read, update and delete
type DataBit interface {
	//returns coloumn family or table name
	BitName() string
	Fields() []string
	FieldValue(fieldName string) (interface{}, error)
	IsUpdated() bool
	SetForUpdate(bool) error
	SetValue(fieldName string, value interface{}) error
}

//DataBucket is a collection of DataBits
type DataBucket interface {
	DataBucketDef
	DataBits() map[string]DataBit
	Fields() []string
	FieldValue(fieldName string) (interface{}, error)
}

type DataBucketDef interface {
	AddDataBit(bitUsage string, dataBit DataBit) error
	AddField(fieldName string, bitFieldName string, dataBitUsage string)
	SetFieldValue(fieldName string, value interface{}) error
}

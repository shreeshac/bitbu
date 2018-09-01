package bitbu

//DataBit is the lowest level data exchange from and to the application
//used for db operations like create,read, update and delete
type DataBit interface {
	//returns coloumn family or table name
	BitName() string
}

type BitReader interface {
	DataBit
	Fields() []string
	FieldValue(fieldName string) (interface{}, error)
	FieldValues() []interface{}
}
type BitUpdater interface {
	DataBit
	IsUpdated() bool
	SetForUpdate(bool) error
	SetValue(fieldName string, value interface{}) error
}
type BitReadUpdater interface {
	BitReader
	IsUpdated() bool
	SetForUpdate(bool) error
	SetValue(fieldName string, value interface{}) error
}

//DataBucket is a collection of DataBits
type DataBucket interface {
	DataBucketDef
	Filters() map[string]DataBucketField
	FilterFieldNames() []string
	DataBits() map[string]DataBit
	DataBitUsages() []string
	Fields(changed BucketFieldListOptions) []string
	FieldValue(fieldName string) (interface{}, error)
}

type DataBucketDef interface {
	AddDataBit(bitUsage string, dataBit DataBit) error
	AddField(fieldName string, bitFieldName string, dataBitUsage string, value interface{})
	SetFieldValue(fieldName string, value interface{}) error
}

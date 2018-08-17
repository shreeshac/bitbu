package bitbu

import "errors"

type DataBucketField struct {
	Name                  string
	BitFieldNameWithUsage []string
	Value                 interface{}
}

type DefaultDataBucket struct {
	//to store sql table name or NoSQL column family name
	Name       string
	fields     map[string]DataBucketField
	fieldNames []string
	dataBits   map[string]DataBit
	//for update
	isUpdated bool
}

func (b DefaultDataBucket) DataBits() map[string]DataBit {

	return b.dataBits
}

func (b DefaultDataBucket) Fields() []string {
	return b.fieldNames
}

func (b DefaultDataBucket) FieldValue(fieldName string) (interface{}, error) {
	switch fieldName {
	case "Name":
		return b.Name, nil
	}
	return nil, errors.New(ErrNoSuchField)
}
func (b DefaultDataBucket) AddDataBit(bitUsage string, dataBit DataBit) error {
	b.dataBits[bitUsage] = dataBit
	return nil
}
func (b DefaultDataBucket) AddField(fieldName string, bitFieldName string, dataBitUsage string) {
	b.fieldNames = append(b.fieldNames, fieldName)
	b.fields[fieldName] = DataBucketField{
		Name: fieldName,
		BitFieldNameWithUsage: []string{bitFieldName + "." + dataBitUsage},
	}

}
func (b DefaultDataBucket) SetFieldValue(fieldName string, value interface{}) error {
	if field, ok := b.fields[fieldName]; !ok {
		return errors.New(ErrNoSuchField)
	} else {
		field.Value = value
	}
	return nil
}

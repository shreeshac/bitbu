package bitbu

import (
	"errors"
	"strings"
)

var (
	ErrInvalidFieldNameWithUsage = "Error: Invalid FieldNameWithUsage"
)

type FieldNameWithUsage string

func (f FieldNameWithUsage) BitUsage() (string, error) {
	if ok, err := f.Valid(); ok {
		return strings.Split(string(f), ".")[1], nil
	} else {
		return "", err
	}

}
func (f FieldNameWithUsage) Name() (string, error) {
	if ok, err := f.Valid(); ok {
		return strings.Split(string(f), ".")[0], nil
	} else {
		return "", err
	}
}
func (f FieldNameWithUsage) Valid() (bool, error) {
	if strings.Count(string(f), ".") == 1 {
		return true, nil
	}

	return false, errors.New(ErrInvalidFieldNameWithUsage)
}

func NewFieldNameWithUsage(bitFieldName, bitUsage string) FieldNameWithUsage {
	return FieldNameWithUsage(strings.Join([]string{bitFieldName, bitUsage}, "."))
}

type DataBucketField struct {
	Name                  string
	BitFieldNameWithUsage []string
	OldValue              interface{}
	Value                 interface{}
}

type DefaultDataBucket struct {
	//to store sql table name or NoSQL column family name
	Name string
	BaseBucket
}

func (b DefaultDataBucket) FieldValue(fieldName string) (interface{}, error) {
	switch fieldName {
	case "Name":
		return b.Name, nil
	}
	return nil, errors.New(ErrNoSuchField)
}

func (b *DefaultDataBucket) SetFieldValue(fieldName string, value interface{}) error {
	_, ok := b.fields[fieldName]
	if !ok {
		return errors.New(ErrNoSuchField)
	}

	switch fieldName {
	case "Name":
		b.Name = value.(string)
	}

	b.changedFieldNames = append(b.changedFieldNames, fieldName)
	b.isUpdated = true
	return nil
}

func NewDefaultDataBucket() DefaultDataBucket {
	defaultDataBucket := DefaultDataBucket{
		BaseBucket: BaseBucket{dataBits: make(map[string]DataBit),
			fields:       make(map[string]DataBucketField),
			filterFields: make(map[string]DataBucketField)},
	}
	ddbit := DefaultDataBit{}
	defaultDataBucket.AddDataBit("Users", ddbit)

	defaultDataBucket.AddField("Name", "Name", "Users", &defaultDataBucket.Name)
	//defaultDataBucket.fields

	return defaultDataBucket
}

package bitbu

import "errors"

type ValidationValue int8

const (
	ValidationForUpdate ValidationValue = iota
	ValidationForCreate
	ValidationForDelete
	ValidationForRead
)

var (
	ErrNoSuchField    = "Error: No such field"
	ErrBitNameIsEmpty = "Error: Bit name can't be empty"
)

//DefaultDataBit is a default implementation of DataBit
type DefaultDataBit struct {
	//to store sql table name or NoSQL column family name
	_tableName string

	FieldNames  []string
	fieldValues []interface{}
	//for update
	isUpdated bool
}

func (b DefaultDataBit) BitType() string {
	return b._tableName
}

func (b DefaultDataBit) Fields() []string {
	return b.FieldNames
}

func (b DefaultDataBit) FieldValue(fieldName string) (interface{}, error) {

	return nil, errors.New(ErrNoSuchField)
}
func (b DefaultDataBit) FieldValues() []interface{} {

	return b.fieldValues
}

func (b DefaultDataBit) IsUpdated() bool {
	return b.isUpdated
}
func (b *DefaultDataBit) SetForUpdate(t bool) error {
	b.isUpdated = t
	return nil
}
func (b *DefaultDataBit) SetValue(fieldName string, value interface{}) error {
	b.FieldNames = append(b.FieldNames, fieldName)
	b.fieldValues = append(b.fieldValues, value)

	return nil

}
func (b DefaultDataBit) Valdiate() bool {
	return true
}

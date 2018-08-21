package bitbu

const (
	CQLTokenUpdate = " UPDATE "
	CQLTokenSet    = " SET "
	CQLTokenWhere  = " WHERE "
	CQLTokenAnd    = " AND "
)

type BucketFieldListOptions bool

const (
	ListBucketFieldsUpdated BucketFieldListOptions = true
	ListBucketFieldsAll     BucketFieldListOptions = false
)

func genUpdateCQLstatement(bitB DataBucket) (string, []interface{}) {
	updateCQL := CQLTokenUpdate
	//fields := strings.Join(bitB.Fields(false), ",")
	assignments := ""
	var values []interface{}
	for _, updatedFieldName := range bitB.Fields(ListBucketFieldsUpdated) {

		assignments = assignments + " " + updatedFieldName + " = ?"
		v, _ := bitB.FieldValue(updatedFieldName)
		values = append(values, v)
	}

	filters := ""
	for _, field := range bitB.Filters() {
		if len(filters) == 0 {
			filters = filters + CQLTokenWhere + field.Name + " = ?"
			v := field.Value
			values = append(values, v)
		} else {
			filters = filters + CQLTokenAnd + field.Name + " = ?"
		}
	}
	return updateCQL + CQLTokenSet + assignments + filters, values
}

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
	for _, field := range bitB.FilterFieldNames() {
		if len(filters) == 0 {
			filters = filters + CQLTokenWhere + field + " = ?"
			v := bitB.Filters()[field].Value
			values = append(values, v)
		} else {
			filters = filters + CQLTokenAnd + field + " = ?"
		}
	}
	return updateCQL + bitB.DataBitUsages()[0] + CQLTokenSet + assignments + filters, values
}

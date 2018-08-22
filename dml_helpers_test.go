package bitbu

import (
	"strings"
	"testing"
)

// type DataBucketTest struct {
// 	Name string
// 	DefaultDataBucket
// }

func TestGenUpdateCQLstatement(t *testing.T) {
	defaultDataBucket := NewDefaultDataBucket()
	updateCQL, _ := genUpdateCQLstatement(&defaultDataBucket)
	if len(updateCQL) == 0 {
		t.Error("Update CQL string can't be empty")
	}
	if !strings.HasPrefix(updateCQL, CQLTokenUpdate) {
		t.Errorf("Update query is not starting with 'Update '. updateCQL-'%s'", updateCQL)
	}

	if strings.Count(updateCQL, CQLTokenSet) < 1 {
		t.Errorf("Update CQL should contain SET keyword. CQL='%s'", updateCQL)
	}
	err := defaultDataBucket.SetFieldValue("Name", "shreesha")
	if err != nil {
		t.Fatal(err)
	}
	updateCQL, _ = genUpdateCQLstatement(&defaultDataBucket)
	if strings.Count(updateCQL, "Name = ?") != 1 {
		t.Errorf("Update CQL should contain update coloumn 'Name'. CQL='%s'", updateCQL)

	}
	defaultDataBucket.AddFilter("Name", "shreesha2")
	updateCQL, _ = genUpdateCQLstatement(&defaultDataBucket)
	t.Log(updateCQL)
	if !strings.Contains(updateCQL, "WHERE Name = ?") {
		t.Errorf("AddFilter is not setting the filter. Substr expected='WHERE Name = ?' CQL='%s'", updateCQL)
	}
	defaultDataBucket.AddFilter("email", "shreesha2")
	updateCQL, values := genUpdateCQLstatement(&defaultDataBucket)
	t.Log(updateCQL)
	if strings.Count(updateCQL, "AND email = ?") != 1 {
		t.Errorf("AddFilter is not setting the filter. Substr expected='AND email = ?'. CQL='%s'", updateCQL)
	}
	if len(values) != 2 {
		t.Errorf("Bucket_AddFilter() values length should be 2. CQL='%s' len of values='%d'", updateCQL, len(values))

	}
	want := " UPDATE Users SET  Name = ? WHERE Name = ? AND email = ?"
	result, values := genUpdateCQLstatement(&defaultDataBucket)

	if want != result {
		t.Errorf("Expected '%s', found '%s'", want, result)
	}

}

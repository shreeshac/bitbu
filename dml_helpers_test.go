package bitbu

import (
	"strings"
	"testing"
)

// type DataBucketTest struct {
// 	Name string
// 	DefaultDataBucket
// }

func Test_genUpdateCQLstatement(t *testing.T) {
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
func Test_genInsertCQL(t *testing.T) {
	b := NewDefaultDataBucket()

	resQueries, resValues := genInsertCQL(&b)
	for usage, resQuery := range resQueries {
		if len(resQuery) == 0 {
			t.Error("insert CQL cannot be empty")
		}
		if len(resValues[usage]) == 0 {
			t.Error("values cannot be empty")
		}
		if !strings.HasPrefix(resQuery, CQLTokenInsert) {
			t.Errorf("Insert CQL should start with CQLTokenInsert keyword. found:'%s'", resQuery)
		}
		if !strings.Contains(resQuery, "?") {
			t.Errorf("Insert CQL should contain atleast one placeholder. found:'%s'", resQuery)
		}
	}
}

func Test_genValuesPlaceHolderString(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Two values",
			args: args{
				args: []interface{}{"val1", "val2"},
			},
			want: "?,?",
		},
		{
			name: "No values",
			args: args{
				args: []interface{}{},
			},
			want: "",
		},
		{
			name: "One value",
			args: args{
				args: []interface{}{"val1"},
			},
			want: "?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genValuesPlaceHolderString(tt.args.args...); got != tt.want {
				t.Errorf("genValuesPlaceHolderString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genSelectCQL(t *testing.T) {
	b := NewDefaultDataBucket()
	query, value := genSelectCQL(&b)
	if query == "" {
		t.Error("SELECT Query can not be empty")
	}
	if len(value) == 0 {
		t.Error("SELECT CQL- Value can't be empty")
	}

	if !strings.HasPrefix(query, CQLTokenSelect) {
		t.Errorf("Select CQL should start with '%s'.found '%s'", CQLTokenSelect, query)
	}

	if !strings.Contains(query, CQLTokenFrom) {
		t.Errorf("Select CQL should contain 'From' keyword. found '%s'", query)
	}
	if !strings.Contains(query, b.DataBits()["Users"].BitType()) {
		t.Errorf("genSelectCQL: Select CQL should contain 'Users' table. found '%s'", query)
	}
	if strings.Contains(query, strings.Join(b.Fields(false), ",")) {
		t.Errorf("genSelectCQL: Select CQL should contain all fields in the table. found '%s'", query)
	}
}

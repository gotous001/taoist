package conf

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func Test_conf_Get(t *testing.T) {
	r, err := os.Open("conf1_test.json")
	if err != nil {
		panic(err)
	}
	Conf = newConf(r)

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "simpleKey1",
			args: args{key: "name"},
			want: "Alexander",
		},
		{
			name: "simpleKey2",
			args: args{key: "address"},
			want: "No 178, abc street, salt lake city",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Conf.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
			fmt.Println("key: ", tt.args.key, " val: ", got)
		})
	}
}

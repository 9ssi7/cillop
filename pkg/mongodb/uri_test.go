package mongodb

import "testing"

func TestCalcMongoUri(t *testing.T) {
	type args struct {
		params UriParams
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestCalcMongoUri with all params",
			args: args{
				params: UriParams{
					Host: "localhost",
					Port: "27017",
					User: "user",
					Pass: "pass",
					Db:   "db",
				},
			},
			want: "mongodb://user:pass@localhost:27017/db",
		},
		{
			name: "TestCalcMongoUri without user and pass",
			args: args{
				params: UriParams{
					Host: "localhost",
					Port: "27017",
					User: "",
					Pass: "",
					Db:   "db",
				},
			},
			want: "mongodb://localhost:27017/db",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcMongoUri(tt.args.params); got != tt.want {
				t.Errorf("CalcMongoUri() = %v, want %v", got, tt.want)
			}
		})
	}
}

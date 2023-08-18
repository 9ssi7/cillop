package mongodb

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDB_TransformId(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    primitive.ObjectID
		wantErr bool
	}{
		{
			name: "Test TransformId with valid id",
			args: args{
				id: "000000000000000000000000",
			},
			want:    primitive.NilObjectID,
			wantErr: false,
		},
		{
			name: "Test TransformId with invalid id",
			args: args{
				id: "x",
			},
			want:    primitive.NilObjectID,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.TransformId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.TransformId() = %v, want %v", got, tt.want)
			}
		})
	}
}

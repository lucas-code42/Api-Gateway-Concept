package requests

import (
	"Api-Gateway-lcs42/models"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPrepareRequest(t *testing.T) {
	type args struct {
		c      *gin.Context
		method string
	}
	tests := []struct {
		name    string
		args    args
		want    models.RequestHost
		wantErr bool
	}{}
	// TODO: Add test cases.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrepareRequest(tt.args.c, tt.args.method)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrepareRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrepareRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

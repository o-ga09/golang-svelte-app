package handlers_test

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGetAllNotes(t *testing.T) {
	cases := []struct{
		name string
		want events.APIGatewayProxyResponse
		err error
		wantErr bool
	}{
		{name:"ケース1",want: events.APIGatewayProxyResponse{StatusCode: 200,Body: "OK"},err: nil,wantErr: false},
	}

	for _, tt := range cases { 
		t.Run(tt.name, func(t *testing.T) {
			
		})
	}
}
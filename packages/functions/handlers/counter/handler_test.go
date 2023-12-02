package handlers_test

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/o-ga09/golang-svelte-app/counter"
	handlers "github.com/o-ga09/golang-svelte-app/handlers/counter"
)

func TestGetCounter(t *testing.T) {
	cases := []struct {
		name    string
		want    events.APIGatewayProxyResponse
		err     error
		wantErr bool
	}{
		{name: "ケース1", want: events.APIGatewayProxyResponse{StatusCode: 200, Body: "{\"Count\":1}"}, err: nil, wantErr: false},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			repo := &counter.RepositoryMock{
				GetFunc: func() (interface{}, error) {
					return counter.Clicks{Counter: "clicks", Count: 1}, tt.err
				},
			}
			count := counter.New(repo)
			handler := handlers.New(count)
			got, err := handler.GetCouter(context.TODO())
			if (err == nil) == tt.wantErr {
				t.Errorf("want err : %v, but got err : %v", tt.want, err)
			}

			if got.Body != tt.want.Body {
				t.Errorf("want : %v, but got : %v", tt.want.Body, got.Body)
			}

			if got.StatusCode != tt.want.StatusCode {
				t.Errorf("want : %v, but got : %v", tt.want.StatusCode, got.StatusCode)
			}
		})
	}
}

func TestUpdateCount(t *testing.T) {
	cases := []struct {
		name    string
		want    events.APIGatewayProxyResponse
		err     error
		wantErr bool
	}{
		{name: "ケース1", want: events.APIGatewayProxyResponse{StatusCode: 200, Body: "{\"Count\":1}"}, err: nil, wantErr: false},
		// {name: "ケース2", want: events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error"}, wantErr: true},
		// {name: "ケース3", want: events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error"}, wantErr: true},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			repo := &counter.RepositoryMock{
				GetFunc: func() (interface{}, error) {
					return counter.Clicks{Counter: "clicks", Count: 1}, tt.err
				},
				PutFunc: func(interface{}) error {
					return tt.err
				},
			}
			count := counter.New(repo)
			handler := handlers.New(count)
			got, err := handler.UpdateCounter(context.TODO())
			if (err == nil) == tt.wantErr {
				t.Errorf("want err : %v, but got err : %v", tt.want, err)
			}

			if got.Body != tt.want.Body {
				t.Errorf("want : %v, but got : %v", tt.want.Body, got.Body)
			}

			if got.StatusCode != tt.want.StatusCode {
				t.Errorf("want : %v, but got : %v", tt.want.StatusCode, got.StatusCode)
			}
		})
	}
}

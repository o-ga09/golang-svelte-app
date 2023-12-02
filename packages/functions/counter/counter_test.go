package counter_test

import (
	"testing"

	"github.com/guregu/dynamo"
	"github.com/o-ga09/golang-svelte-app/counter"
)

func TestGetCount(t *testing.T) {
	cases := []struct {
		name    string
		want    int
		err     error
		wantErr bool
	}{
		{name: "ケース1", want: 1, err: nil, wantErr: false},
		{name: "ケース2", want: 0, err: dynamo.ErrNotFound, wantErr: true},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			repo := &counter.RepositoryMock{
				GetFunc: func() (interface{}, error) {
					return counter.Clicks{Counter: "clicks", Count: tt.want}, tt.err
				},
			}
			count := counter.New(repo)
			got, err := count.GetCount()
			if (err == nil) == tt.wantErr {
				t.Errorf("want err : %v, but got err : %v", tt.err, err)
			}

			if got != tt.want {
				t.Errorf("want : %v, but got : %v", tt.want, got)
			}
		})
	}
}

func TestUpdateCount(t *testing.T) {
	cases := []struct {
		name    string
		want    int
		err     error
		wantErr bool
	}{
		{name: "ケース1", err: nil, wantErr: false},
		{name: "ケース2", err: dynamo.ErrNotFound, wantErr: true},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			repo := &counter.RepositoryMock{
				PutFunc: func(interface{}) error {
					return tt.err
				},
				GetFunc: func() (interface{}, error) {
					return counter.Clicks{Counter: "clicks",Count: 1}, tt.err
				},
			}
			count := counter.New(repo)
			err := count.UpdateCount()
			if (err == nil) == tt.wantErr {
				t.Errorf("want err : %v, but got err : %v", tt.err, err)
			}
		})
	}
}

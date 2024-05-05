package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
	searchapi "github.com/Leonargo404-code/find-my-brand/pkg/searchAPI"
	"github.com/stretchr/testify/mock"
)

func Test_service_SearchTerms(t *testing.T) {
	type fields struct {
		SearchAPI func() searchapi.SearchAPI
	}

	type args struct {
		findBrandReq *brand.FindBrandRequest
		location     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *brand.Result
		wantErr bool
	}{
		{
			name: "should return a error empty terms",
			fields: fields{
				SearchAPI: func() searchapi.SearchAPI {
					m := &searchapi.Mock{}

					return m
				},
			},
			args: args{
				findBrandReq: &brand.FindBrandRequest{
					Terms: "",
				},
				location: "Brazil",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return a error empty terms 2",
			fields: fields{
				SearchAPI: func() searchapi.SearchAPI {
					m := &searchapi.Mock{}

					return m
				},
			},
			args: args{
				findBrandReq: &brand.FindBrandRequest{
					Terms: "  ",
				},
				location: "Brazil",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return a google search error",
			fields: fields{
				SearchAPI: func() searchapi.SearchAPI {
					m := &searchapi.Mock{}

					m.On("GoogleSearch", mock.Anything, mock.Anything, mock.Anything).Return(
						nil, errors.New(""),
					)

					return m
				},
			},
			args: args{
				findBrandReq: &brand.FindBrandRequest{
					Terms: "buy",
				},
				location: "Brazil",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return a google search results",
			fields: fields{
				SearchAPI: func() searchapi.SearchAPI {
					m := &searchapi.Mock{}

					m.On("GoogleSearch", mock.Anything, mock.Anything, mock.Anything).Return(
						&brand.Result{
							Ads: []brand.Ads{
								{
									Position: 1,
								},
							},
						}, nil,
					)

					return m
				},
			},
			args: args{
				findBrandReq: &brand.FindBrandRequest{
					Terms: "buy",
				},
				location: "Brazil",
			},
			want: &brand.Result{
				Ads: []brand.Ads{
					{
						Position: 1,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				SearchAPI: tt.fields.SearchAPI(),
			}

			got, err := s.SearchTerms(tt.args.findBrandReq, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.SearchTerms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.SearchTerms() = %v, want %v", got, tt.want)
			}
		})
	}
}

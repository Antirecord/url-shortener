package service

import (
	"testing"

	"github.com/antirecord/url-shortener/internal/app/entity"
	"github.com/stretchr/testify/assert"
)

func TestSimpleUrlShortener_Shorten(t *testing.T) {
	us := SimpleUrlShortener{storage: map[string]entity.StorageEntity{}}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		us      SimpleUrlShortener
		args    args
		val     string
		wantErr bool
	}{
		{
			name:    "first url shorten test",
			us:      us,
			args:    args{"https://testurl.ru"},
			val:     "newUrl",
			wantErr: false,
		},
		{
			name:    "second url shorten test",
			us:      us,
			args:    args{"testurl.ru"},
			val:     "newUrl",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.us.Shorten(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimpleUrlShortener.Shorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil {
				assert.Equal(t, err.Error(), "url должен начинаться на http:// или https://")
				return
			}
			assert.NotEmpty(t, got)
			assert.Contains(t, got, "localhost:8080")
			assert.NotEqual(t, got, tt.val)
		})
	}
}

func TestSimpleUrlShortener_GetBaseUrl(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		us      SimpleUrlShortener
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.us.GetBaseUrl(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimpleUrlShortener.GetBaseUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SimpleUrlShortener.GetBaseUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeHash(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first mergeHash test",
			args: args{
				hash: "123asd45",
			},
			want: "http://localhost:8080/123asd45",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeHash(tt.args.hash); got != tt.want {
				t.Errorf("mergeHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

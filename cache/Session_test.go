package cache

import (
	"reflect"
	"testing"
)

func TestSession_Put(t *testing.T) {
	type fields struct {
		Name             string
		TTL              int64
		SessionInterface SessionInterface
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				Name:             tt.fields.Name,
				TTL:              tt.fields.TTL,
				SessionInterface: tt.fields.SessionInterface,
			}
			if err := s.Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Session.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSession_Get(t *testing.T) {
	type fields struct {
		Name             string
		TTL              int64
		SessionInterface SessionInterface
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				Name:             tt.fields.Name,
				TTL:              tt.fields.TTL,
				SessionInterface: tt.fields.SessionInterface,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setSliceMap(t *testing.T) {
	type args struct {
		m     map[string]interface{}
		keys  []string
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setSliceMap(tt.args.m, tt.args.keys, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setSliceMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSliceMap(t *testing.T) {
	type args struct {
		m    map[string]interface{}
		keys []string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSliceMap(tt.args.m, tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSliceMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_delSliceMap(t *testing.T) {
	type args struct {
		m    map[string]interface{}
		keys []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delSliceMap(tt.args.m, tt.args.keys)
		})
	}
}

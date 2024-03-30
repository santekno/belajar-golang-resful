package middleware

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/julienschmidt/httprouter"
)

func TestNew(t *testing.T) {
	type args struct {
		constructors []Constructor
	}
	tests := []struct {
		name string
		args args
		want Chain
	}{
		{
			name: "case 1 -> success",
			args: args{
				constructors: []Constructor{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.constructors...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChain_Then(t *testing.T) {
	type fields struct {
		constructors []Constructor
	}
	type args struct {
		h httprouter.Handle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   httprouter.Handle
	}{
		{
			name: "Test Middleware 1 success",
			args: args{
				h: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				},
			},
			want: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

			},
		},
		{
			name: "Test Middleware 2 more than one",
			args: args{
				h: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				},
			},
			fields: fields{
				constructors: []Constructor{
					func(h httprouter.Handle) httprouter.Handle {
						return h
					},
					func(h httprouter.Handle) httprouter.Handle {
						return h
					},
				},
			},
			want: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

			},
		},
		{
			name: "Test Middleware 3 is nil",
			args: args{
				h: nil,
			},
			want: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Chain{
				constructors: tt.fields.constructors,
			}
			if got := c.Then(tt.args.h); assert.ObjectsAreEqual(got, tt.want) {
				t.Errorf("Chain.Then() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChain_Append(t *testing.T) {
	type fields struct {
		constructors []Constructor
	}
	type args struct {
		constructors []Constructor
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Chain
	}{
		{
			name: "Test append middleware",
			args: args{
				constructors: []Constructor{
					func(httprouter.Handle) (h httprouter.Handle) {
						return
					},
				},
			},
			want: Chain{
				constructors: []Constructor{
					func(httprouter.Handle) (h httprouter.Handle) {
						return
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Chain{
				constructors: tt.fields.constructors,
			}
			if got := c.Append(tt.args.constructors...); assert.ObjectsAreEqual(got, tt.want) {
				t.Errorf("Chain.Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

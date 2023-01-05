package models

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateSpecDoc(t *testing.T) {
	validationRuleset = "testdata/spectral-validation.yaml"
	type args struct {
		ctx context.Context
		sd  SpecDoc
	}
	tests := []struct {
		name      string
		args      args
		wantValid bool
	}{
		{
			name: "valid - pestore v2 json",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/petstore-v2.json"),
			},
			wantValid: true,
		},
		{
			name: "valid - pestore v2 yaml",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/petstore-v2.yaml"),
			},
			wantValid: true,
		},
		{
			name: "valid - pestore v3 json",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/petstore-v3.json"),
			},
			wantValid: true,
		},
		{
			name: "valid - pestore v3 yaml",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/petstore-v3.yaml"),
			},
			wantValid: true,
		},
		{
			name: "invalid",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/sample-invalid-api.yaml"),
			},
			wantValid: false,
		},
		{
			name: "invalid - invalid ref in yaml spec",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/petstore-v2-invalid-ref.yaml"),
			},
			wantValid: false,
		},
		{
			name: "invalid - invalid ref in json spec",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/petstore-v2-invalid-ref.json"),
			},
			wantValid: false,
		},
		{
			name: "invalid - missing version in yaml spec",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/petstore-v2-missing-version.yaml"),
			},
			wantValid: false,
		},
		{
			name: "invalid - missing version in json spec",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/petstore-v2-missing-version.json"),
			},
			wantValid: false,
		},
		{
			name: "invalid - unknown key in yaml spec",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/petstore-v2-unknown-key.yaml"),
			},
			wantValid: false,
		},
		{
			name: "invalid - unknown key in json spec",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/petstore-v2-unknown-key.json"),
			},
			wantValid: false,
		},
		{
			name: "invalid - invalid json",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/petstore-v2-invalid-json.json"),
			},
			wantValid: false,
		},
		{
			name: "invalid - invalid json",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/sample-invalid-json.json"),
			},
			wantValid: false,
		},
		{
			name: "invalid - invalid yaml",
			args: args{
				ctx: context.TODO(),
				sd:  loadSpec("testdata/sample-invalid-yaml.yaml"),
			},
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateSpecDoc(tt.args.ctx, tt.args.sd)
			if err != nil {
				t.Errorf("ValidateSpecDoc error: %s", err)
				return
			}
			assert.Equal(t, tt.wantValid, got.Valid, "ValidateSpecDoc(%v, %v)", tt.args.ctx, tt.args.sd)
		})
	}
}

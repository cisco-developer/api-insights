// Copyright 2022 Cisco Systems, Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestNewSpecDocFromBytes(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want SpecDoc
	}{
		{
			name: "normal",
			args: args{data: loadSpecData("testdata/petstore-v2.json")},
			want: SpecDoc(loadSpec("testdata/petstore-v2.json")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpecDocFromBytes(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpecDocFromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpec_BeforeCreate(t *testing.T) {
	type fields struct {
		ID        string
		Doc       SpecDoc
		DocType   string
		Revision  string
		Score     *int
		ServiceID string
		State     string
		Valid     string
		Version   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DocOAS    *openapi3.T
	}
	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "normal - init id",
			fields:  fields{ID: ""},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Spec{
				ID:        tt.fields.ID,
				Doc:       tt.fields.Doc,
				Revision:  tt.fields.Revision,
				Score:     tt.fields.Score,
				ServiceID: tt.fields.ServiceID,
				State:     tt.fields.State,
				Valid:     tt.fields.Valid,
				Version:   tt.fields.Version,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DocOAS:    tt.fields.DocOAS,
			}
			if err := m.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NotEmpty(t, m.ID)
		})
	}
}

func TestSpec_GetID(t *testing.T) {
	type fields struct {
		ID        string
		Doc       SpecDoc
		DocType   string
		Revision  string
		Score     *int
		ServiceID string
		State     string
		Valid     string
		Version   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DocOAS    *openapi3.T
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{ID: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Spec{
				ID:        tt.fields.ID,
				Doc:       tt.fields.Doc,
				Revision:  tt.fields.Revision,
				Score:     tt.fields.Score,
				ServiceID: tt.fields.ServiceID,
				State:     tt.fields.State,
				Valid:     tt.fields.Valid,
				Version:   tt.fields.Version,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DocOAS:    tt.fields.DocOAS,
			}
			if got := m.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpec_GetIndex(t *testing.T) {
	type fields struct {
		ID        string
		Doc       SpecDoc
		DocType   string
		Revision  string
		Score     *int
		ServiceID string
		State     string
		Valid     string
		Version   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DocOAS    *openapi3.T
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "index service_id",
			fields: fields{},
			args:   args{field: "service_id"},
			want:   "idx_service_id",
		},
		{
			name:   "index version",
			fields: fields{},
			args:   args{field: "version"},
			want:   "idx_version",
		},
		{
			name:   "index revision",
			fields: fields{},
			args:   args{field: "revision"},
			want:   "idx_revision",
		},
		{
			name:   "index state",
			fields: fields{},
			args:   args{field: "state"},
			want:   "idx_state",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Spec{
				ID:        tt.fields.ID,
				Doc:       tt.fields.Doc,
				Revision:  tt.fields.Revision,
				Score:     tt.fields.Score,
				ServiceID: tt.fields.ServiceID,
				State:     tt.fields.State,
				Valid:     tt.fields.Valid,
				Version:   tt.fields.Version,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DocOAS:    tt.fields.DocOAS,
			}
			if got := m.GetIndex(tt.args.field); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpec_GetIndexValue(t *testing.T) {
	type fields struct {
		ID        string
		Doc       SpecDoc
		DocType   string
		Revision  string
		Score     *int
		ServiceID string
		State     string
		Valid     string
		Version   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DocOAS    *openapi3.T
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "index service_id",
			fields: fields{ServiceID: "test"},
			args:   args{field: "service_id"},
			want:   "test",
		},
		{
			name:   "index version",
			fields: fields{Version: "v1.0"},
			args:   args{field: "version"},
			want:   "v1.0",
		},
		{
			name:   "index revision",
			fields: fields{Revision: "1"},
			args:   args{field: "revision"},
			want:   "1",
		},
		{
			name:   "index state",
			fields: fields{State: "Development"},
			args:   args{field: "state"},
			want:   "Development",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Spec{
				ID:        tt.fields.ID,
				Doc:       tt.fields.Doc,
				Revision:  tt.fields.Revision,
				Score:     tt.fields.Score,
				ServiceID: tt.fields.ServiceID,
				State:     tt.fields.State,
				Valid:     tt.fields.Valid,
				Version:   tt.fields.Version,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DocOAS:    tt.fields.DocOAS,
			}
			if got := m.GetIndexValue(tt.args.field); got != tt.want {
				t.Errorf("GetIndexValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpec_GetTags(t *testing.T) {
	type fields struct {
		ID        string
		Doc       SpecDoc
		DocType   string
		Revision  string
		Score     *int
		ServiceID string
		State     string
		Valid     string
		Version   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DocOAS    *openapi3.T
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "normal",
			fields: fields{
				ServiceID: "test",
				Version:   "v1.0",
				Revision:  "1",
				State:     "Development",
			},
			want: []string{"test", "v1.0", "1", "Development"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Spec{
				ID:        tt.fields.ID,
				Doc:       tt.fields.Doc,
				Revision:  tt.fields.Revision,
				Score:     tt.fields.Score,
				ServiceID: tt.fields.ServiceID,
				State:     tt.fields.State,
				Valid:     tt.fields.Valid,
				Version:   tt.fields.Version,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DocOAS:    tt.fields.DocOAS,
			}
			if got := m.GetTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpec_LoadDocAsOAS(t *testing.T) {
	type fields struct {
		ID        string
		Doc       SpecDoc
		DocType   string
		Revision  string
		Score     *int
		ServiceID string
		State     string
		Valid     string
		Version   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DocOAS    *openapi3.T
	}
	type args struct {
		ctx        context.Context
		validate   bool
		setVersion bool
	}
	emptyDoc := ""

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *openapi3.T
		wantErr bool
	}{
		{
			name: "nil doc",
			fields: fields{
				ID:        "",
				Doc:       nil,
				DocType:   "",
				Revision:  "",
				Score:     nil,
				ServiceID: "",
				State:     "",
				Valid:     "",
				Version:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DocOAS:    nil,
			},
			args: args{
				ctx:        nil,
				validate:   false,
				setVersion: false,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty doc",
			fields: fields{
				ID:        "",
				Doc:       &emptyDoc,
				DocType:   "",
				Revision:  "",
				Score:     nil,
				ServiceID: "",
				State:     "",
				Valid:     "",
				Version:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DocOAS:    nil,
			},
			args: args{
				ctx:        nil,
				validate:   false,
				setVersion: false,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "with existing DocOAS",
			fields: fields{
				ID:        "",
				Doc:       nil,
				DocType:   "",
				Revision:  "",
				Score:     nil,
				ServiceID: "",
				State:     "",
				Valid:     "",
				Version:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DocOAS:    &openapi3.T{},
			},
			args: args{
				ctx:        nil,
				validate:   false,
				setVersion: false,
			},
			want:    &openapi3.T{},
			wantErr: false,
		},
		{
			name: "normal",
			fields: fields{
				ID:        "",
				Doc:       SpecDoc(loadSpec("testdata/sample-api.yaml")),
				DocType:   "",
				Revision:  "",
				Score:     nil,
				ServiceID: "",
				State:     "",
				Valid:     "",
				Version:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DocOAS:    nil,
			},
			args: args{
				ctx:        nil,
				validate:   true,
				setVersion: false,
			},
			want:    &openapi3.T{},
			wantErr: false,
		},
		{
			name: "normal - set version",
			fields: fields{
				ID:        "",
				Doc:       SpecDoc(loadSpec("testdata/sample-api.yaml")),
				DocType:   "",
				Revision:  "",
				Score:     nil,
				ServiceID: "",
				State:     "",
				Valid:     "",
				Version:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DocOAS:    nil,
			},
			args: args{
				ctx:        nil,
				validate:   true,
				setVersion: true,
			},
			want:    &openapi3.T{},
			wantErr: false,
		},
		{
			name: "normal - set docType",
			fields: fields{
				ID:        "",
				Doc:       SpecDoc(loadSpec("testdata/sample-api.yaml")),
				DocType:   "",
				Revision:  "",
				Score:     nil,
				ServiceID: "",
				State:     "",
				Valid:     "",
				Version:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DocOAS:    nil,
			},
			args: args{
				ctx:        nil,
				validate:   true,
				setVersion: true,
			},
			want:    &openapi3.T{},
			wantErr: false,
		},
		{
			name: "normal - set docType oas2",
			fields: fields{
				ID:        "",
				Doc:       SpecDoc(loadSpec("testdata/petstore-v2.json")),
				DocType:   "",
				Revision:  "",
				Score:     nil,
				ServiceID: "",
				State:     "",
				Valid:     "",
				Version:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DocOAS:    nil,
			},
			args: args{
				ctx:        nil,
				validate:   false,
				setVersion: true,
			},
			want:    &openapi3.T{},
			wantErr: false,
		},
		{
			name: "invalidate spec",
			fields: fields{
				ID:        "",
				Doc:       SpecDoc(loadSpec("testdata/sample-invalid-api.yaml")),
				DocType:   "",
				Revision:  "",
				Score:     nil,
				ServiceID: "",
				State:     "",
				Valid:     "",
				Version:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DocOAS:    nil,
			},
			args: args{
				ctx:        nil,
				validate:   true,
				setVersion: false,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Spec{
				ID:        tt.fields.ID,
				Doc:       tt.fields.Doc,
				Revision:  tt.fields.Revision,
				Score:     tt.fields.Score,
				ServiceID: tt.fields.ServiceID,
				State:     tt.fields.State,
				Valid:     tt.fields.Valid,
				Version:   tt.fields.Version,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DocOAS:    tt.fields.DocOAS,
			}
			got, err := m.LoadDocAsOAS(tt.args.ctx, tt.args.validate, tt.args.setVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadDocAsOAS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotNil(t, got)
				if tt.args.setVersion {
					assert.NotEmpty(t, m.Version)
				}
			}

			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("LoadDocAsOAS() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestSpec_Sortable(t *testing.T) {
	type fields struct {
		ID        string
		Doc       SpecDoc
		DocType   string
		Revision  string
		Score     *int
		ServiceID string
		State     string
		Valid     string
		Version   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DocOAS    *openapi3.T
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "sort by created_at",
			fields: fields{},
			args:   args{field: "created_at"},
			want:   true,
		},
		{
			name:   "sort by updated_at",
			fields: fields{},
			args:   args{field: "created_at"},
			want:   true,
		},
		{
			name:   "sort by unsupported field",
			fields: fields{},
			args:   args{field: "unsupported"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Spec{
				ID:        tt.fields.ID,
				Doc:       tt.fields.Doc,
				Revision:  tt.fields.Revision,
				Score:     tt.fields.Score,
				ServiceID: tt.fields.ServiceID,
				State:     tt.fields.State,
				Valid:     tt.fields.Valid,
				Version:   tt.fields.Version,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DocOAS:    tt.fields.DocOAS,
			}
			if got := m.Sortable(tt.args.field); got != tt.want {
				t.Errorf("Sortable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpec_String(t *testing.T) {
	type fields struct {
		ID        string
		Doc       SpecDoc
		DocType   string
		Revision  string
		Score     *int
		ServiceID string
		State     string
		Valid     string
		Version   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DocOAS    *openapi3.T
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{ServiceID: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Spec{
				ID:        tt.fields.ID,
				Doc:       tt.fields.Doc,
				Revision:  tt.fields.Revision,
				Score:     tt.fields.Score,
				ServiceID: tt.fields.ServiceID,
				State:     tt.fields.State,
				Valid:     tt.fields.Valid,
				Version:   tt.fields.Version,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DocOAS:    tt.fields.DocOAS,
			}
			got := m.String()
			assert.NotEmpty(t, got)
		})
	}
}

func TestSpec_TableName(t *testing.T) {
	type fields struct {
		ID        string
		Doc       SpecDoc
		DocType   string
		Revision  string
		Score     *int
		ServiceID string
		State     string
		Valid     string
		Version   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DocOAS    *openapi3.T
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{},
			want:   SpecTableName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Spec{
				ID:        tt.fields.ID,
				Doc:       tt.fields.Doc,
				Revision:  tt.fields.Revision,
				Score:     tt.fields.Score,
				ServiceID: tt.fields.ServiceID,
				State:     tt.fields.State,
				Valid:     tt.fields.Valid,
				Version:   tt.fields.Version,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DocOAS:    tt.fields.DocOAS,
			}
			if got := m.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func loadSpec(file string) *string {
	content, _ := os.ReadFile(file)
	s := string(content)
	return &s
}

func loadSpecData(file string) []byte {
	content, _ := os.ReadFile(file)
	return content
}

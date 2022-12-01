package persistence

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/iDevoid/cptx"
	"github.com/iDevoid/stygis/internal/constant/model"
	"github.com/iDevoid/stygis/internal/constant/query"
)

func TestUserInit(t *testing.T) {
	type args struct {
		db cptx.Database
	}
	tests := []struct {
		name string
		args args
		want UserPersistence
	}{
		{
			name: "success",
			args: args{
				db: nil,
			},
			want: &userPersistence{
				db: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserInit(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userPersistence_InsertUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	type fields struct {
		db cptx.Database
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				db: func() cptx.Database {
					_mockedMain := cptx.NewMockMainDB(ctrl)
					_mockedMain.EXPECT().QueryRowMustTx(gomock.Any(), query.UserInsert,
						map[string]interface{}{
							"username":     "clyf",
							"email":        "clyf@example.com",
							"hashed_email": "ajsdhkjasdns",
							"password":     "password",
							"create_time":  time.Date(2020, 12, 21, 12, 12, 12, 0, time.UTC),
							"status":       0,
						},
						gomock.Any(),
					).Return(errors.New("ERROR"))
					mocked := cptx.NewMockDatabase(ctrl)
					mocked.EXPECT().Main().Return(_mockedMain)
					return mocked
				}(),
			},
			args: args{
				ctx: context.TODO(),
				user: &model.User{
					Username:    "clyf",
					Email:       "clyf@example.com",
					HashedEmail: "ajsdhkjasdns",
					Password:    "password",
					CreateTime:  time.Date(2020, 12, 21, 12, 12, 12, 0, time.UTC),
				},
			},
			wantErr: true,
		}, {
			name: "success",
			fields: fields{
				db: func() cptx.Database {
					_mockedMain := cptx.NewMockMainDB(ctrl)
					_mockedMain.EXPECT().QueryRowMustTx(gomock.Any(), query.UserInsert,
						map[string]interface{}{
							"username":     "clyf",
							"email":        "clyf@example.com",
							"hashed_email": "ajsdhkjasdns",
							"password":     "password",
							"create_time":  time.Date(2020, 12, 21, 12, 12, 12, 0, time.UTC),
							"status":       0,
						},
						gomock.Any(),
					).Return(nil)
					mocked := cptx.NewMockDatabase(ctrl)
					mocked.EXPECT().Main().Return(_mockedMain)
					return mocked
				}(),
			},
			args: args{
				ctx: context.TODO(),
				user: &model.User{
					Username:    "clyf",
					Email:       "clyf@example.com",
					HashedEmail: "ajsdhkjasdns",
					Password:    "password",
					CreateTime:  time.Date(2020, 12, 21, 12, 12, 12, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := &userPersistence{
				db: tt.fields.db,
			}
			if err := up.InsertUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userPersistence.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

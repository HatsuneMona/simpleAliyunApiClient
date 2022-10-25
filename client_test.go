package simpleAliyunApiClient

import (
	"testing"
)

func Test_New(t *testing.T) {
	type Args struct {
		ak string
		as string
	}
	cases := []struct {
		name    string
		args    Args
		wantErr bool
	}{
		{
			name: "success",
			args: Args{
				ak: "test ak",
				as: "test as",
			},
			wantErr: false,
		},
		{
			name: "as empty",
			args: Args{
				ak: "test ak",
				as: "",
			},
			wantErr: true,
		},
		{
			name: "ak empty",
			args: Args{
				ak: "",
				as: "test as",
			},
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.ak, tt.args.as)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("new client result: %+v", got)
		})
	}
}

func Test_do(t *testing.T) {
	client, _ := New("akakakakak", "asasasasas")
	smsAction := &SendSms{
		PhoneNumbers:  "13866665555",
		SignName:      "HatsuneMona",
		TemplateCode:  "CODE_1234563",
		TemplateParam: "{param:sth}",
	}
	err := client.Do(smsAction)
	if err != nil {
		t.Error(err)
	}

}

package argsParser

import "testing"

func TestParseIssueNFTArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid arguments",
			args:    []string{"file.json", "password123", "1.0", "50000", "CoolCollection", "CCC"},
			wantErr: false,
		},
		{
			name:    "Invalid number of arguments",
			args:    []string{"file.json", "password123", "1.0", "50000", "CoolCollection"},
			wantErr: true,
			errMsg:  "invalid number of arguments",
		},
		{
			name:    "Short JSON file name",
			args:    []string{"json", "password123", "1.0", "50000", "CoolCollection", "CCC"},
			wantErr: true,
			errMsg:  "invalid length of jsonFile argument",
		},
		{
			name:    "Short password",
			args:    []string{"file.json", "pass", "1.0", "50000", "CoolCollection", "CCC"},
			wantErr: true,
			errMsg:  "password too short",
		},
		{
			name:    "Invalid value",
			args:    []string{"file.json", "password123", "abc", "50000", "CoolCollection", "CCC"},
			wantErr: true,
			errMsg:  "invalid value provided for NFT token",
		},
		{
			name:    "Invalid gas limit",
			args:    []string{"file.json", "password123", "1.0", "abc", "CoolCollection", "CCC"},
			wantErr: true,
			errMsg:  "invalid gasLimit provided for NFT token",
		},
		{
			name:    "Short collection name",
			args:    []string{"file.json", "password123", "1.0", "50000", "CC", "CC"},
			wantErr: true,
			errMsg:  "collection name too short",
		},
		{
			name:    "Short collection ticker",
			args:    []string{"file.json", "password123", "1.0", "50000", "CoolCollection", "C"},
			wantErr: true,
			errMsg:  "collection ticker too short, must be at least 3 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseIssueNFTArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIssueNFTArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("ParseIssueNFTArgs() error message = %v, wantErrMsg %v", err.Error(), tt.errMsg)
			}
		})
	}
}

package server

import "testing"

func Test_isOnDNSanctionsList(t *testing.T) {
	tests := map[string]struct {
		address string
		want    bool
	}{
		"Check different cases": {
			address: "0X0203298C575Bf7eFcEfB58c0dDd0aE7417439C19",
			want:    true,
		},
		"Check upper cases": {
			address: "0X0203298C575BF7EFCEFB58C0DDD0AE7417439C19",
			want:    true,
		},
		"Check unknown": {
			address: "0X5",
			want:    false,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := isOnDNSanctionsList(testCase.address); got != testCase.want {
				t.Errorf("isOnDNSanctionsList() = %v, want %v", got, testCase.want)
			}
		})
	}
}

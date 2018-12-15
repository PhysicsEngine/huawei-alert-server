package config

import (
	"os"
	"testing"
)

const (
	testGCPProjectID = "test_project"
)

func TestReadFromEnv(t *testing.T) {
	reset := setenvs(t, map[string]string{
		"ENV":            envDevelopment,
	})
	defer reset()

	env, err := ReadFromEnv()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

}

func TestReadFromEnvValidationFailed(t *testing.T) {
	reset := setenvs(t, map[string]string{
		"ENV":            "production",
	})
	defer reset()

	_, err := ReadFromEnv()
	if err == nil {
		t.Fatalf("expect to be faield")
	}
}

func TestReadFromEnvProcessFailed(t *testing.T) {
	reset := unsetenv(t, "GCP_PROJECT_ID")
	defer reset()

	_, err := ReadFromEnv()
	if err == nil {
		t.Fatalf("expect to be faield")
	}
}

func TestIsProduction(t *testing.T) {
	cases := []struct {
		env  *Env
		want bool
	}{
		{
			&Env{
				Env: envDevelopment,
			},
			false,
		},

		{
			&Env{
				Env: envProduction,
			},
			true,
		},

		{
			&Env{
				Env: "staging",
			},
			false,
		},
	}

	for _, tc := range cases {
		if got := tc.env.IsProduction(); got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestValidate(t *testing.T) {
	cases := map[string]struct {
		env     *Env
		success bool
	}{
		"Valid1": {
			&Env{
				Env: envDevelopment,
			},
			true,
		},

		"Valid2": {
			&Env{
				Env: envProduction,
			},
			true,
		},

		"InvalidEnv": {
			&Env{
				Env: "staging",
			},
			false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := tc.env.validate()
			if err != nil {
				if tc.success {
					t.Fatalf("expect not to be failed: %s", err)
				}
				return
			}

			if !tc.success {
				t.Fatalf("expect to be failed")
			}
		})
	}
}

func setenv(t *testing.T, k, v string) func() {
	t.Helper()

	prev := os.Getenv(k)
	if err := os.Setenv(k, v); err != nil {
		t.Fatal(err)
	}

	return func() {
		if prev == "" {
			os.Unsetenv(k)
		} else {
			if err := os.Setenv(k, prev); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func unsetenv(t *testing.T, k string) func() {
	t.Helper()

	prev := os.Getenv(k)
	if err := os.Unsetenv(k); err != nil {
		t.Fatal(err)
	}

	return func() {
		if prev == "" {
			return
		}
		if err := os.Setenv(k, prev); err != nil {
			t.Fatal(err)
		}
		// default
		return
	}
}

func setenvs(t *testing.T, kv map[string]string) func() {
	t.Helper()

	resetFs := make([]func(), 0, len(kv))
	for k, v := range kv {
		resetF := setenv(t, k, v)
		resetFs = append(resetFs, resetF)
	}

	return func() {
		for _, resetF := range resetFs {
			resetF()
		}
	}
}

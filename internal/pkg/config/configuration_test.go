package config

import (
	"flag"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func TestConfigurationPrecedence(t *testing.T) {
	for _, override := range []bool{false, true} {
		name := "file"
		if override {
			name = "environment"
		}
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.json")
			if err := os.WriteFile(path, []byte(`{"Server":{"Port":"4321"},"Database":{"Driver":"sqlite","MaxOpenConns":5},"Data":{"BaseURL":"https://file.test","WebhookID":"file-id","WebhookToken":"file-token","KVAPIKeys":["file-key"],"Sites":[{"Name":"file-site","Url":"https://file.test"}],"SpecialLinks":[{"Key":"reserved","Link":"https://special.test"}]}}`), 0600); err != nil {
				t.Fatal(err)
			}
			prepareConfigurationTest(t, path)
			if override {
				t.Setenv("BASE_URL", "https://env.test")
				t.Setenv("WEBHOOK_ID", "env-id")
				t.Setenv("WEBHOOK_TOKEN", "env-token")
				t.Setenv("ENABLECORS", "on")
				t.Setenv("CONFIG_DATA_SITES", `[{"Name":"env-site","Url":"https://env.test"}]`)
			}
			Setup()
			got := GetConfig()
			if got.Server.Port != "4321" || got.Server.Mode != "release" || got.Database.Driver != "sqlite" || got.Database.MaxOpenConns != 5 {
				t.Fatalf("file/default settings changed: server=%+v database=%+v", got.Server, got.Database)
			}
			wantURL, wantID, wantToken, wantSite := "https://file.test", "file-id", "file-token", "file-site"
			if override {
				wantURL = "https://env.test"
			}
			if got.Data.EnableCORS != "off" {
				t.Fatal("unsupported CORS environment variable changed the default")
			}
			if got.Data.BaseURL != wantURL || got.Data.WebhookID != wantID || got.Data.WebhookToken != wantToken {
				t.Fatal("configuration precedence or webhook fields changed")
			}
			if len(got.Data.Sites) != 1 || got.Data.Sites[0].Name != wantSite || len(got.Data.KVAPIKeys) != 1 || got.Data.KVAPIKeys[0] != "file-key" || len(got.Data.SpecialLinks) != 1 || got.Data.SpecialLinks[0].Key != "reserved" {
				t.Fatal("configuration collections changed")
			}
		})
	}
}

func TestConfigurationMissingFileDefaults(t *testing.T) {
	prepareConfigurationTest(t, filepath.Join(t.TempDir(), "missing.json"))
	Setup()
	if len(Config.Data.CommonAPIKeys) != 0 {
		t.Fatal("common keys must default to empty")
	}
	if Config.Server.Port != "3000" || Config.Server.Mode != "release" || Config.Database.Driver != "" {
		t.Fatalf("fallback settings changed: %+v", Config)
	}
	assertDataDefaults(t, Config.Data)
}

func TestConfigurationDataDefaults(t *testing.T) {
	for _, content := range []string{`{}`, `{"Data":{}}`} {
		t.Run(content, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.json")
			if err := os.WriteFile(path, []byte(content), 0600); err != nil {
				t.Fatal(err)
			}
			prepareConfigurationTest(t, path)
			Setup()
			assertDataDefaults(t, Config.Data)
		})
	}
}

func assertDataDefaults(t *testing.T, data DataConfiguration) {
	t.Helper()
	if data.EnableCORS != "off" || data.EnableWebhookAPI != "off" {
		t.Fatal("optional switches must default to off")
	}
	if data.WebhookID != "" || data.WebhookToken != "" || data.BaseURL != "" || data.AgentRecordsPath != "" || data.LinkRedirectPageURL != "" {
		t.Fatal("optional strings must default to empty")
	}
	if len(data.WebhookAPIKeys) != 0 || len(data.CommonAPIKeys) != 0 || len(data.KVAPIKeys) != 0 || len(data.Sites) != 0 || len(data.SpecialLinks) != 0 {
		t.Fatal("optional lists must default to empty")
	}
}

func TestConfigurationExplicitDataValues(t *testing.T) {
	for _, setting := range []string{"on", "off", ""} {
		t.Run("switches="+setting, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.json")
			content := `{"Data":{"EnableCORS":"` + setting + `","EnableWebhookAPI":"` + setting + `","WebhookID":"","WebhookToken":"","BaseURL":"","AgentRecordsPath":"","LinkRedirectPageURL":"","WebhookAPIKeys":[],"CommonAPIKeys":[],"KVAPIKeys":[],"Sites":[],"SpecialLinks":[]}}`
			if err := os.WriteFile(path, []byte(content), 0600); err != nil {
				t.Fatal(err)
			}
			prepareConfigurationTest(t, path)
			Setup()
			if Config.Data.EnableCORS != setting || Config.Data.EnableWebhookAPI != setting {
				t.Fatal("explicit switch values were overridden by defaults")
			}
			data := Config.Data
			data.EnableCORS, data.EnableWebhookAPI = "off", "off"
			assertDataDefaults(t, data)
		})
	}
}

func TestConfigurationEnvironmentOverridesDefaults(t *testing.T) {
	prepareConfigurationTest(t, filepath.Join(t.TempDir(), "missing.json"))
	t.Setenv("ENABLECORS", "on")
	t.Setenv("WEBHOOK_ID", "env-id")
	t.Setenv("WEBHOOK_TOKEN", "env-token")
	t.Setenv("BASE_URL", "https://env.test")
	Setup()
	if Config.Data.EnableCORS != "off" || Config.Data.EnableWebhookAPI != "off" || Config.Data.WebhookID != "" || Config.Data.WebhookToken != "" || Config.Data.BaseURL != "https://env.test" {
		t.Fatal("only BASE_URL may override defaults when the file is missing")
	}
}

func TestConfigurationIgnoresUnsupportedEnvironment(t *testing.T) {
	for _, content := range []string{`{}`, `{"Data":{"EnableCORS":"on","EnableWebhookAPI":"on","WebhookID":"file-id","WebhookToken":"file-token","Sites":[{"Name":"file-site"}]}}`} {
		t.Run(content, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.json")
			if err := os.WriteFile(path, []byte(content), 0600); err != nil {
				t.Fatal(err)
			}
			prepareConfigurationTest(t, path)
			Setup()
			want := Config.Data
			prepareConfigurationTest(t, path)
			for key, value := range map[string]string{
				"ENABLECORS":            "off",
				"WEBHOOK_ID":            "env-id",
				"WEBHOOK_TOKEN":         "env-token",
				"CONFIG_DATA_SITES":     `[{"Name":"env-site"}]`,
				"CONFIG_TYPE":           "yaml",
				"CONFIG-PATH":           "ignored.json",
				"DATA.ENABLECORS":       "off",
				"DATA.WEBHOOKID":        "env-id",
				"DATA.WEBHOOKTOKEN":     "env-token",
				"DATA.ENABLEWEBHOOKAPI": "off",
			} {
				t.Setenv(key, value)
			}
			Setup()
			if !reflect.DeepEqual(Config.Data, want) {
				t.Fatal("unsupported environment variables changed JSON settings or defaults")
			}
		})
	}
}

func TestConfigurationAbsentBaseURLEnvironment(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(path, []byte(`{"Data":{"BaseURL":"https://file.test"}}`), 0600); err != nil {
		t.Fatal(err)
	}
	prepareConfigurationTest(t, path)
	if err := os.Unsetenv("BASE_URL"); err != nil {
		t.Fatal(err)
	}
	Setup()
	if Config.Data.BaseURL != "https://file.test" {
		t.Fatal("absent BASE_URL must preserve the JSON value")
	}
}

func TestConfigurationDefaultPathIgnoresEnvironment(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "data", "config.json")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(`{"Server":{"Port":"4321"}}`), 0600); err != nil {
		t.Fatal(err)
	}
	prepareConfigurationTest(t, path)
	t.Chdir(root)
	os.Args = []string{"configuration-test"}
	t.Setenv("CONFIG-PATH", "ignored.json")
	t.Setenv("CONFIG_TYPE", "yaml")
	Setup()
	if Config.Server.Port != "4321" {
		t.Fatal("unsupported environment variables changed the default JSON config path")
	}
}

func TestCommonAPIKeysConfiguration(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(path, []byte(`{"Data":{"CommonAPIKeys":["first","second"],"WebhookAPIKeys":["webhook"],"KVAPIKeys":["kv"]}}`), 0600); err != nil {
		t.Fatal(err)
	}
	prepareConfigurationTest(t, path)
	Setup()
	if len(Config.Data.CommonAPIKeys) != 2 || Config.Data.CommonAPIKeys[0] != "first" || Config.Data.CommonAPIKeys[1] != "second" || Config.Data.WebhookAPIKeys[0] != "webhook" || Config.Data.KVAPIKeys[0] != "kv" {
		t.Fatal("independent API key lists were not loaded")
	}
}

func TestWarningConfigurationIndependentOfCORS(t *testing.T) {
	for _, cors := range []string{"on", "off"} {
		t.Run(cors, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.json")
			content := `{"Data":{"LinkRedirectPageURL":"https://warning.test/","EnableCORS":"` + cors + `"}}`
			if err := os.WriteFile(path, []byte(content), 0600); err != nil {
				t.Fatal(err)
			}
			prepareConfigurationTest(t, path)
			Setup()
			if Config.Data.LinkRedirectPageURL != "https://warning.test/" || Config.Data.EnableCORS != cors {
				t.Fatal("warning configuration changed with CORS")
			}
		})
	}
}

func prepareConfigurationTest(t *testing.T, path string) {
	t.Helper()
	previousArgs, previousFlags, previousPFlags, previousConfig := os.Args, flag.CommandLine, pflag.CommandLine, Config
	os.Args = []string{"configuration-test", "--config-path", path}
	flag.CommandLine = flag.NewFlagSet("configuration-test", flag.ContinueOnError)
	pflag.CommandLine = pflag.NewFlagSet("configuration-test", pflag.ContinueOnError)
	viper.Reset()
	for _, key := range []string{"BASE_URL", "WEBHOOK_ID", "WEBHOOK_TOKEN", "ENABLECORS", "CONFIG_DATA_SITES"} {
		t.Setenv(key, "")
	}
	t.Setenv("CONFIG_TYPE", "json")
	// This unsupported environment variable must not select a config file.
	t.Setenv("CONFIG-PATH", "ignored-by-explicit-flag.json")
	t.Cleanup(func() {
		os.Args, flag.CommandLine, pflag.CommandLine, Config = previousArgs, previousFlags, previousPFlags, previousConfig
		viper.Reset()
	})
}

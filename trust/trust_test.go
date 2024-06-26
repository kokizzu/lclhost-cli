package trust

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/anchordotdev/cli"
	"github.com/anchordotdev/cli/api/apitest"
	"github.com/anchordotdev/cli/cmdtest"
	"github.com/anchordotdev/cli/ui/uitest"
	"github.com/stretchr/testify/require"
)

var srv = &apitest.Server{
	Host:    "api.anchor.lcl.host",
	RootDir: "../..",
}

func TestMain(m *testing.M) {
	flag.Parse()

	if err := srv.Start(context.Background()); err != nil {
		panic(err)
	}

	defer os.Exit(m.Run())

	srv.Close()
}

func TestCmdTrust(t *testing.T) {
	cmd := CmdTrust
	cfg := cli.ConfigFromCmd(cmd)
	cfg.Test.SkipRunE = true

	t.Run("--help", func(t *testing.T) {
		cmdtest.TestOutput(t, cmd, "trust", "--help")
	})

	t.Run("--no-sudo", func(t *testing.T) {
		t.Cleanup(func() {
			cfg.Trust.Org = ""
		})

		cmdtest.TestExecute(t, cmd, "trust", "--no-sudo")

		require.Equal(t, true, cfg.Trust.NoSudo)
	})

	t.Run("--organization test", func(t *testing.T) {
		t.Cleanup(func() {
			cfg.Trust.Org = ""
		})

		cmdtest.TestExecute(t, cmd, "trust", "--organization", "test")

		require.Equal(t, "test", cfg.Trust.Org)
	})

	t.Run("-o test", func(t *testing.T) {
		t.Cleanup(func() {
			cfg.Trust.Org = ""
		})

		cmdtest.TestExecute(t, cmd, "trust", "-o", "test")

		require.Equal(t, "test", cfg.Trust.Org)
	})

	t.Run("--realm test", func(t *testing.T) {
		t.Cleanup(func() {
			cfg.Trust.Realm = ""
		})

		cmdtest.TestExecute(t, cmd, "trust", "--realm", "test")

		require.Equal(t, "test", cfg.Trust.Realm)
	})

	t.Run("-r test", func(t *testing.T) {
		t.Cleanup(func() {
			cfg.Trust.Realm = ""
		})

		cmdtest.TestExecute(t, cmd, "trust", "-r", "test")

		require.Equal(t, "test", cfg.Trust.Realm)
	})

	t.Run("--trust-stores nss,system", func(t *testing.T) {
		t.Cleanup(func() {
			cfg.Trust.Stores = []string{"homebrew", "nss", "system"}
		})

		cmdtest.TestExecute(t, cmd, "trust", "--trust-stores", "nss,system")

		require.Equal(t, []string{"nss", "system"}, cfg.Trust.Stores)
	})
}

func TestTrust(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := new(cli.Config)
	cfg.API.URL = srv.URL
	cfg.NonInteractive = true
	cfg.Trust.MockMode = true
	cfg.Trust.NoSudo = true
	cfg.Trust.Stores = []string{"mock"}
	var err error
	if cfg.API.Token, err = srv.GeneratePAT("anky@anchor.dev"); err != nil {
		t.Fatal(err)
	}
	ctx = cli.ContextWithConfig(ctx, cfg)

	t.Run("basics", func(t *testing.T) {
		if !srv.IsProxy() {
			t.Skip("trust unsupported in mock mode")
		}

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		cmd := Command{}

		uitest.TestTUIOutput(ctx, t, cmd.UI())
	})
}

package component_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/anchordotdev/cli"
	"github.com/anchordotdev/cli/api"
	"github.com/anchordotdev/cli/api/apitest"
	"github.com/anchordotdev/cli/component"
	"github.com/anchordotdev/cli/ui/uitest"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"
)

var srv = &apitest.Server{
	Host:    "api.anchor.lcl.host",
	RootDir: "../..",
}

func TestMain(m *testing.M) {
	if err := srv.Start(context.Background()); err != nil {
		panic(err)
	}

	defer os.Exit(m.Run())

	srv.Close()
}

func TestSelector(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error
	cfg := new(cli.Config)
	if cfg.API.Token, err = srv.GeneratePAT("anky@anchor.dev"); err != nil {
		t.Fatal(err)
	}
	cfg.API.URL = srv.URL

	anc, err := api.NewClient(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("orgs solo", func(t *testing.T) {
		if srv.IsProxy() {
			t.Skip("selector tests unsupported in proxy mode")
		}

		cfg.Test.Prefer = map[string]cli.ConfigTestPrefer{
			"/v0/orgs": {
				Example: "solo",
			},
		}
		ctx = cli.ContextWithConfig(ctx, cfg)

		drv, tm := uitest.TestTUI(ctx, t)

		errc := make(chan error, 1)
		go func() {
			choicec, err := component.OrgSelector(ctx, drv, anc, "Which organization do you want for this test?")
			if err != nil {
				errc <- err
			}

			org := <-choicec

			if want, got := "solo-org-slug", org.Slug; want != got {
				errc <- fmt.Errorf("Want org choice: %q, Got: %q", want, got)
			}

			errc <- tm.Quit()
		}()

		if err := <-errc; err != nil {
			t.Fatal(err)
		}

		tm.WaitFinished(t, teatest.WithFinalTimeout(time.Second*3))
		uitest.TestGolden(t, drv.Golden())
	})

	t.Run("orgs double", func(t *testing.T) {
		if srv.IsProxy() {
			t.Skip("selector tests unsupported in proxy mode")
		}

		cfg.Test.Prefer = map[string]cli.ConfigTestPrefer{
			"/v0/orgs": {
				Example: "double",
			},
		}
		ctx = cli.ContextWithConfig(ctx, cfg)

		drv, tm := uitest.TestTUI(ctx, t)

		choicec := make(chan api.Organization, 1)
		errc := make(chan error, 1)
		go func() {
			selectorChoicec, err := component.OrgSelector(ctx, drv, anc, "Which organization do you want for this test?")
			if err != nil {
				errc <- err
			}

			choicec <- <-selectorChoicec

			errc <- tm.Quit()
		}()

		uitest.WaitForGoldenContains(t, drv, errc,
			"? Which organization do you want for this test?",
		)

		tm.Send(tea.KeyMsg{
			Type: tea.KeyDown,
		})
		tm.Send(tea.KeyMsg{
			Type: tea.KeyEnter,
		})

		org := <-choicec

		if want, got := "second-org-slug", org.Slug; want != got {
			errc <- fmt.Errorf("Want org choice: %q, Got: %q", want, got)
		}

		tm.WaitFinished(t, teatest.WithFinalTimeout(time.Second*3))
		uitest.TestGolden(t, drv.Golden())
	})

}

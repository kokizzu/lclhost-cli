package models

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/anchordotdev/cli/truststore"
	"github.com/anchordotdev/cli/ui"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type TrustHeader struct{}

func (m *TrustHeader) Init() tea.Cmd { return nil }

func (m *TrustHeader) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

func (m *TrustHeader) View() string {
	var b strings.Builder
	fmt.Fprintln(&b, ui.Header(fmt.Sprintf("Manage CA Certificates in your Local Trust Store(s) %s", ui.Whisper("`anchor trust`"))))
	return b.String()
}

type TrustHint struct{}

func (m *TrustHint) Init() tea.Cmd { return nil }

func (m *TrustHint) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

func (m *TrustHint) View() string {
	var b strings.Builder
	fmt.Fprintln(&b, ui.StepHint(fmt.Sprintf("%s %s",
		ui.Accentuate("This may require sudo privileges, learn why here: "),
		ui.URL("https://lcl.host/why-sudo"),
	)))
	return b.String()
}

type TrustUpdateConfirm struct {
	ConfirmCh chan<- struct{}
}

func (m *TrustUpdateConfirm) Init() tea.Cmd { return nil }

func (m *TrustUpdateConfirm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.ConfirmCh != nil {
				close(m.ConfirmCh)
				m.ConfirmCh = nil
			}
		}
	}

	return m, nil
}

func (m *TrustUpdateConfirm) View() string {
	var b strings.Builder

	if m.ConfirmCh != nil {
		fmt.Fprintln(&b, ui.StepAlert(fmt.Sprintf("%s to install missing certificates. (%s)", ui.Action("Press Enter"), ui.Accentuate("requires sudo"))))

		return b.String()
	}

	return b.String()
}

type TrustUpdateStore struct {
	Store truststore.Store

	installing *truststore.CA
	installed  map[string][]string

	spinner spinner.Model
}

func (m *TrustUpdateStore) Init() tea.Cmd {
	m.installed = make(map[string][]string)
	m.spinner = ui.WaitingSpinner()

	return m.spinner.Tick
}

type (
	TrustStoreInstallingCAMsg struct {
		truststore.CA
	}

	TrustStoreInstalledCAMsg struct {
		truststore.CA
	}
)

func (m *TrustUpdateStore) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TrustStoreInstallingCAMsg:
		m.installing = &msg.CA
		return m, nil
	case TrustStoreInstalledCAMsg:
		m.installing = nil
		m.installed[msg.CA.Subject.CommonName] = append(m.installed[msg.CA.Subject.CommonName], msg.CA.PublicKeyAlgorithm.String())
		return m, nil
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m *TrustUpdateStore) View() string {
	var b strings.Builder

	if len(m.installed) > 0 {
		var styledCAs []string

		for subjectCommonName, algorithms := range m.installed {
			styledCAs = append(styledCAs, fmt.Sprintf("%s [%s]",
				ui.Underline(subjectCommonName),
				ui.Whisper(strings.Join(algorithms, ", ")),
			))
		}

		fmt.Fprintln(&b, ui.StepDone(fmt.Sprintf("Updated %s: installed %s",
			ui.Emphasize(m.Store.Description()),
			strings.Join(styledCAs, ", "),
		)))
	}

	if m.installing != nil {
		// present thumbprint for comparison with pop up prompt
		if runtime.GOOS == "windows" {
			fmt.Fprintln(&b, ui.StepHint(fmt.Sprintf("\"%s\" Thumbprint (sha1): %s",
				m.installing.Subject.CommonName,
				m.installing.WindowsThumbprint(),
			)))
		}

		fmt.Fprintln(&b, ui.StepInProgress(fmt.Sprintf("Updating %s: installing %s %s.",
			ui.Emphasize(m.Store.Description()),
			ui.Underline(m.installing.Subject.CommonName),
			ui.Whisper(m.installing.PublicKeyAlgorithm.String()),
		)))
	}

	return b.String()
}

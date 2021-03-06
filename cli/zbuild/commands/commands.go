package commands

import (
	"github.com/dimes/zbuild/buildlog"

	"github.com/manifoldco/promptui"
)

var (
	// Build is the command that executes a build
	Build Command = &build{}

	// InitWorkspace is the command that initializes a workspace on the local file system
	InitWorkspace Command = &initWorkspace{}

	// Publish is the command that uploads an artifact
	Publish Command = &publish{}

	// Refresh refreshes the workspace metadata
	Refresh Command = &refresh{}
)

// Command is an interface for commands
type Command interface {
	Describe() string
	Exec(workingDir string, args ...string) error
}

func readLineWithPrompt(label string, validate promptui.ValidateFunc, defaultVal string) string {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Default:  defaultVal,
	}

	result, err := prompt.Run()
	if err != nil {
		buildlog.Fatalf("Error getting option for label %s: %+v", label, err)
	}

	return result
}

func getYnConfirmation(label string) (bool, error) {
	prompt := promptui.Select{
		Label: label,
		Items: []string{"Yes", "No"},
	}

	selectedIndex, _, err := prompt.Run()
	if err != nil {
		return false, err
	}

	return selectedIndex == 0, nil
}

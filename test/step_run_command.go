package test

import (
	"log"

	"github.com/jen20/riviera/azure"
)

type StepRunCommand struct {
	RunCommand     azure.ApiCall
	CleanupCommand azure.ApiCall
	StateCommand   azure.ApiCall
	StateBagKey    string
}

func (s *StepRunCommand) Run(state AzureStateBag) StepAction {
	if s.RunCommand == nil && s.StateCommand == nil {
		return Continue
	}

	azureClient := state.Client()
	if s.RunCommand != nil {
		log.Printf("[INFO] Running %T command...", s.RunCommand)

		r := azureClient.NewRequest()
		r.Command = s.RunCommand
		response, err := r.Execute()
		if err != nil {
			state.AppendError(err)
			return Halt
		}

		if response.IsSuccessful() {
			state.Put(s.StateBagKey, response.Parsed)
		} else {
			state.AppendError(response.Error)
			return Halt
		}
	}

	if s.StateCommand != nil {
		log.Printf("[INFO] Refreshing state with %T command...", s.StateCommand)

		r := azureClient.NewRequest()

		r.Command = s.StateCommand
		response, err := r.Execute()
		if err != nil {
			state.AppendError(err)
			return Halt
		}

		if response.IsSuccessful() {
			state.Remove(s.StateBagKey)
			state.Put(s.StateBagKey, response.Parsed)
			return Continue
		} else {
			state.Remove(s.StateBagKey)
			state.AppendError(response.Error)
			return Halt
		}
	}

	return Continue
}

func (s *StepRunCommand) Cleanup(state AzureStateBag) {
	if s.CleanupCommand == nil {
		return
	}
	azureClient := state.Client()

	log.Printf("[INFO] Cleaning up with %T command...", s.CleanupCommand)

	request := azureClient.NewRequest()
	request.Command = s.CleanupCommand
	response, err := request.Execute()
	if err != nil {
		state.AppendError(err)
		return
	}

	if !response.IsSuccessful() {
		log.Printf("[INFO] Error running clean up %T command", s.CleanupCommand)
		state.AppendError(response.Error)
	}
}

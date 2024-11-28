package state_machine

import (
	"fmt"
	"os"
)

type Context interface{}

type State interface {
	Execute() error
	Transition() (State, error)
}

type StateMachine struct {
	State State
}

func (sm *StateMachine) HandleError(err error) {
	if err == nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run(sm *StateMachine) {
	var err error
	for sm.State != nil {
		err = sm.State.Execute()
		sm.HandleError(err)
		sm.State, err = sm.State.Transition()
		sm.HandleError(err)
	}
}

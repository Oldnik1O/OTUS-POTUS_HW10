// Тесты выявления типа текущего состояния и для выполнения команд

package main

import (
	"testing"
)

// MockCommand для тестирования
type MockCommand struct {
	Executed bool
}

func (c *MockCommand) Execute() {
	c.Executed = true
}

func TestStateTransitions(t *testing.T) {
	// Переключение с NormalState на MoveToState
	normal := &NormalState{}
	moveTo := normal.Handle(&MoveToCommand{})
	_, ok := moveTo.(*MoveToState)
	if !ok {
		t.Errorf("Expected state to be MoveToState, got: %T", moveTo)
	}

	// Переключение с MoveToState на NormalState
	normalAgain := moveTo.Handle(&RunCommand{})
	_, ok = normalAgain.(*NormalState)
	if !ok {
		t.Errorf("Expected state to be NormalState, got: %T", normalAgain)
	}

	// Переключение на nil состояние с NormalState
	nilState := normalAgain.Handle(&HardStopCommand{})
	if nilState != nil {
		t.Errorf("Expected state to be nil, got: %T", nilState)
	}
}

func TestCommandExecution(t *testing.T) {
	mockCommand := &MockCommand{}
	normal := &NormalState{}
	normal.Handle(mockCommand)

	if !mockCommand.Executed {
		t.Errorf("Expected MockCommand to be executed")
	}
}

func TestHardStopCommandExecution(t *testing.T) {
	hardStop := &HardStopCommand{}
	state := normal.Handle(hardStop)

	if state != nil {
		t.Errorf("Expected state to be nil after HardStopCommand")
	}
}

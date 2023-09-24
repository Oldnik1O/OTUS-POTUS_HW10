// Смена режимов обработки команд
// Начинаем в "обычном" состоянии и обрабатываем команды последовательно
// При выполнении HardStopCommand состояние становится nil и обработка прекращается

package main
import "fmt"
// State интерфейс определяет метод Handle для обработки команд в зависимости от текущего состояния
type State interface {
  Handle(command Command) State
}
// Command интерфейс для различных команд
type Command interface {
  Execute()
}
// HardStopCommand останавливает обработку
type HardStopCommand struct{}

func (c *HardStopCommand) Execute() {
  fmt.Println("Hard stop command executed!")
}
// MoveToCommand передает команды в другую очередь  
type MoveToCommand struct{}

func (c *MoveToCommand) Execute() {
  fmt.Println("Move to command executed!")
}
// RunCommand выполняет команду.
type RunCommand struct{}

func (c *RunCommand) Execute() {
  fmt.Println("Run command executed!")
}
Теперь реализуем два состояния: "Обычное" и "MoveTo"
go
// NormalState обычное состояние.
type NormalState struct{}
func (s *NormalState) Handle(command Command) State {
  command.Execute()
  switch c := command.(type) {
  case *HardStopCommand:
    return nil
  case *MoveToCommand:
    return &MoveToState{}
  default:
    return s
  }
}
// MoveToState состояние перенаправления команд
type MoveToState struct{}
func (s *MoveToState) Handle(command Command) State {
  command.Execute()
  switch c := command.(type) {
  case *HardStopCommand:
    return nil
  case *RunCommand:
    return &NormalState{}
  default:
    return s
  }
}
// Демонстрационный код:

func main() {
  commands := []Command{&RunCommand{}, &MoveToCommand{}, &RunCommand{}, &HardStopCommand{}}
  currentState := &NormalState{}
  for _, command := range commands {
    if currentState == nil {
      break
    }
    currentState = currentState.Handle(command)
  }
}

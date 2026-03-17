package commands

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	// check for number of arguments
	if len(cmd.arguments) == 0 {
		return errors.New("no arguments given")
	} else if len(cmd.arguments) > 1 {
		return errors.New("more than one arguments given")
	}
	// set new user
	userName := cmd.arguments[0]
	s.File.SetUser(userName)
	fmt.Println("User has been set to", userName)
	return nil
}

// run a function with given state and name
func (comm *Commands) run(s *State, cmd Command) error {
	// check if function exists
	funcName := cmd.name
	fun, ok := comm.RegisteredComms[funcName]
	if !ok {
		return errors.New("function does not exist")
	}
	// call the function
	fun(s, cmd)
	return nil
}

// register a command with a given name
func (comm *Commands) Register(name string, f func(*State, Command) error) {
	comm.RegisteredComms[name] = f
}

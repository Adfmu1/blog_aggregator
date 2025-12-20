package commands

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	// check for number of arguments
	if len(cmd.arguments) == 0 {
		return errors.New("no arguments given")
	} else if len(cmd.arguments) > 1 {
		return errors.New("more than one arguments given")
	}
	// set new user
	userName := cmd.arguments[0]
	s.file.SetUser(userName)
	fmt.Println("User has been set to", userName)
	return nil
}

// run a function with given state and name
func (comm *commands) run(s *state, cmd command) error {
	// check if function exists
	funcName := cmd.name
	fun, ok := comm.nameToHandr[funcName]
	if !ok {
		return errors.New("function does not exist")
	}
	// call the function
	fun(s, cmd)
	return nil
}

// register a command with a given name
func (comm *commands) register(name string, f func(*state, command) error) {
	comm.nameToHandr[name] = f
}

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
	s.file.CurrentUser = userName
	fmt.Println("User has been set to", userName)
	return nil
}

func (comm commands) run(s *state, cmd command) error {
	funcName := cmd.name
	fun, ok := comm.nameToHandr[funcName]
	if !ok {
		return errors.New("function does not exist")
	}
	fun(s, cmd)
	return nil
}

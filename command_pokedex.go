package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemon")
	}

	fmt.Println("Your Pokemon:")
	for _, poke := range cfg.caughtPokemon {
		fmt.Println(" - ", poke.Name)
	}
	return nil
}

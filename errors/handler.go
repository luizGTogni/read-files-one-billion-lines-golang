package errors

import (
	"errors"
	"fmt"
	"project-advanced-concepts/errors/types"
)

func HandlerError(err error) {
	if err != nil {
		switch {
			case errors.Is(err, &types.FileNotFoundError{}):
				fmt.Printf("%s", err)
			default:
				panic(err)
		}
	}
}
package canonical

import (
	"errors"
	"fmt"
)

var (
	ErrorNotFound = fmt.Errorf("entity not found")
)

func HandleError(err error) error {
	if errors.Is(err, ErrorNotFound) {
		return err
	}
	return fmt.Errorf("unexpected error occurred %w", err)

}

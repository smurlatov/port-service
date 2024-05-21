package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

func fetchPortsFromJSON(ctx context.Context, r io.Reader, portChan chan Port) error {
	decoder := json.NewDecoder(r)

	// Get first delimiter
	token, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("failed to read delimiter: %w", err)
	}

	// check first delimiter is `{`
	if token != json.Delim('{') {
		return fmt.Errorf("expected {, got %v", token)
	}

	for decoder.More() {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		// Read port ID.
		token, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("failed to read port ID: %w", err)
		}

		portId, ok := token.(string)
		if !ok {
			return fmt.Errorf("expected string, got %v", token)
		}

		// Read port structure
		var port Port
		if err := decoder.Decode(&port); err != nil {
			return fmt.Errorf("failed to decode port: %w", err)
		}

		port.Id = portId
		portChan <- port
	}

	return nil
}

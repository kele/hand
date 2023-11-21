package hand

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
)

// Prompt is called after the function supplied to HelpWith returns an error.
var Prompt = func(fErr error, args ...any) error {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("hand: %v:%v -- f(%v) = (_, %v).\n", file, line, args, fErr)
	fmt.Println("hand: Fix?")
	return nil
}

// GetAnswer should fill the object with the needed value.
var GetAnswer = func(object any) error {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	return json.Unmarshal(input, object)
}

// IsAnAnswer returns true if the answer supplied via GetAnswer should be
// considered as one (cf. treated as "I don't know the answer").
var IsAnAnswer = func(object any, err error) bool {
	return err == nil
}

// HelpWith helps recover from errors.
func HelpWith[Arg1 any, Result any](f func(Arg1) (Result, error)) func(Arg1) (Result, error) {
	return func(arg1 Arg1) (Result, error) {
		v, fErr := f(arg1)
		if fErr == nil {
			return v, nil
		}
		var ret Result
		if err := Prompt(fErr, arg1); err != nil {
			return ret, fmt.Errorf("hand.Prompt() = %v; original error: %w", err, fErr)
		}
		if err := GetAnswer(&ret); IsAnAnswer(ret, err) {
			return ret, nil
		} else {
			return ret, fmt.Errorf("hand.IsAnAnswer() = false, hand.GetAnswer() = %v; original error: %w", err, fErr)
		}
	}
}

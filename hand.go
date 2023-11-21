package hand

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
)

func HelpWith[Arg1 any, Result any](f func(Arg1) (Result, error)) func(Arg1) (Result, error) {
	return func(arg1 Arg1) (Result, error) {
		v, err := f(arg1)
		if err == nil {
			return v, nil
		}

		var ret Result
		_, file, line, _ := runtime.Caller(1)

		fmt.Printf("%v:%v -- f(%v) = (_, %v).\nFix?\n", file, line, arg1, err)
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			return ret, err
		}
		err = json.Unmarshal(input, &ret)
		return ret, err
	}
}

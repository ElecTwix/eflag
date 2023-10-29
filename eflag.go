package eflag

import (
	"errors"
	"os"

	"github.com/ElecTwix/eflag/flag"
	"github.com/ElecTwix/eflag/pkg/convert"
)

type FlagHandler struct {
	FlagMap     map[string]flag.Flag
	CurrentFlag *flag.Flag
}

func New() FlagHandler {
	flagMap := make(map[string]flag.Flag)
	return FlagHandler{FlagMap: flagMap, CurrentFlag: nil}
}

func (handler FlagHandler) AddFlag(flag *flag.Flag) error {
	_, exists := handler.FlagMap[flag.Key]
	if exists {
		return errors.New("cannot overwrite flag with normal method")
	} else {
		handler.FlagMap[flag.Key] = *flag
	}

	return nil
}

func (handler *FlagHandler) ParseOSArgs() ([]flag.Flag, error) {
	return handler.ParseRaw(os.Args[1:])
}

func (handler *FlagHandler) ParseRaw(rawArgs []string) ([]flag.Flag, error) {
	flags := make([]flag.Flag, 0)
	argsLen := len(rawArgs)

	for i := 0; i < argsLen; i++ {
		flag, ok := handler.FlagMap[rawArgs[i][1:]]
		if !ok {
			return nil, errors.New("flag not exists")
		}
		if flag.Used {
			return nil, errors.New("cannot use same flag multible times")
		}
		if flag.TakesInput {
			if i+1 < argsLen {
				val, err := convert.ConvertStringToType(rawArgs[i+1], flag.InputType)
				if err != nil {
					return nil, err
				}
				flag.Data = val
				i++
			} else {
				return nil, errors.New("out bound")
			}
		}
		flag.Used = true
		handler.FlagMap[rawArgs[i]] = flag
		flags = append(flags, flag)
	}
	return flags, nil
}

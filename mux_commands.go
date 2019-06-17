package cli

import (
	"sort"
	"strings"
)

func (receiver *Mux) Commands(filter ...string) ([][]string, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	var prefix string
	{
		var err error

		prefix, err = serialize(filter...)
		if nil != err {
			return nil, err
		}
	}

	var serializedCommands []string
	{
		receiver.mutex.RLock()
		serializedCommands = make([]string, 0, len(receiver.handlers))
		for key, _ := range receiver.handlers {
			serializedCommands = append(serializedCommands, key)
		}
		receiver.mutex.RUnlock()
	}

	var filteredSerializedCommands []string = serializedCommands[:0]
	for _, serializedCommand := range serializedCommands {
		if strings.HasPrefix(serializedCommand, prefix) {
			filteredSerializedCommands = append(filteredSerializedCommands, serializedCommand)
		}
	}

	sort.Strings(filteredSerializedCommands)

	var commands [][]string = make([][]string, 0, len(filteredSerializedCommands))
	for _, serialized := range filteredSerializedCommands {
		command := unserialize(serialized)

		commands = append(commands, command)
	}

	return commands, nil
}

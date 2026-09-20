package commands

func CreateCommand(name string, args []string) Command {
	/*Creates and returns command struct*/

	return Command{Name: name, Args: args}
}

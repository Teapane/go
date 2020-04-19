package twofer

// ShareWith returns the string "One for <name>, one for me", where name is the argument; if the argument is empty, returns "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}

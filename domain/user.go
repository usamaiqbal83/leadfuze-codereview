package domain


type IUser interface {
	// returns initials as usamaiqbal
	FullName() string
	// returns initials from first name and last name
	Initials() string
	// return combination in format like uiqbal
	Combination1() string
	// return combination in format like usamai
	Combination2() string
	// return combination in format like
	Combination3(infix string) (string, error)
}
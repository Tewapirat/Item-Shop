package exception

type Nopermission struct {}

func (e *Nopermission)Error() string {
	return "no permission"
}
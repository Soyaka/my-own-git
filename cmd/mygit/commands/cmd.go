package commands



type Command struct{
	Name string
	Usage string
	Opts []string
	// Realazor func()error
}
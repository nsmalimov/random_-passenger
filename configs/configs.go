package configs

type Config struct {
	CentralLatitude   float64
	CentralLongitude  float64
	PathToNamesData   string
	Host              string
	Radius            int
	Port              int
	MinSecSleepDriver int
	MaxSecSleepDriver int
	MinSecSleepOrder  int
	MaxSecSleepOrder  int
}

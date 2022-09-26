package lib

type Config struct {
	StartOfDay *HourOfDay
	EndOfDay   *HourOfDay
}

func GetConfig() *Config {
	startOfDay, _ := NewHourOfDay(9, 0)
	endOfDay, _ := NewHourOfDay(9, 0)
	return &Config{
		StartOfDay: startOfDay,
		EndOfDay:   endOfDay,
	}
}

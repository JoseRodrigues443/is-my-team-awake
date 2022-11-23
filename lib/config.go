package lib

type Config struct {
	StartOfDay *HourOfDay
	EndOfDay   *HourOfDay
	Clock      Clock
}

func GetConfig() *Config {
	startOfDay, _ := NewHourOfDay(9, 0)
	endOfDay, _ := NewHourOfDay(18, 0)
	realClock := NewClock()
	return &Config{
		StartOfDay: startOfDay,
		EndOfDay:   endOfDay,
		Clock:      realClock,
	}
}

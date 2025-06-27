package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Flags struct {
	Port                int
	Concurrency         int
	MaxIdleConnsPerHost int
	MaxIdleConns        int
	ReadBody            bool
	CloseBody           bool
}

func ParseFlags() (Flags, error) {

	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	var flags Flags

	flagSet.IntVar(&flags.Port, "port", getIntEnv("CS_PORT", 8080), "server port")
	flagSet.IntVar(&flags.Concurrency, "concurrency", getIntEnv("CS_CONCURRENCY", 2), "concurrent connections")
	flagSet.IntVar(&flags.MaxIdleConnsPerHost, "max-idle-conn-host", getIntEnv("CS_MAX_IDLE_CONN_HOST", 2), "max idle connections per host")
	flagSet.IntVar(&flags.MaxIdleConns, "max-idle-conn", getIntEnv("CS_MAX_IDLE_CONN", 100), "max idle connections")
	flagSet.BoolVar(&flags.ReadBody, "read-body", getBoolEnv("CS_READ_BODY", false), "read response body")
	flagSet.BoolVar(&flags.CloseBody, "close-body", getBoolEnv("CS_CLOSE_BODY", false), "close response body")

	flagSet.Usage = func() {
		fmt.Fprint(flagSet.Output(), "Usage: flag [flags] [args]\n")
		flagSet.PrintDefaults()
	}

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return Flags{}, err
	}

	err := flags.validate()
	return flags, err
}

func (f Flags) validate() error {
	if f.Port < 1 || f.Port > 65535 {
		return fmt.Errorf("invalid port %d", f.Port)
	}
	return nil
}

func getIntEnv(envName string, defaultValue int) int {
	env, ok := os.LookupEnv(envName)
	if !ok {
		return defaultValue
	}

	if intValue, err := strconv.Atoi(env); err == nil {
		return intValue
	}
	return defaultValue
}

func getBoolEnv(envName string, defaultValue bool) bool {
	env, ok := os.LookupEnv(envName)
	if !ok {
		return defaultValue
	}

	if boolValue, err := strconv.ParseBool(env); err == nil {
		return boolValue
	}
	return defaultValue
}

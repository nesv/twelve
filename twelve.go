// Package twelve provides utilities to help make your applications more
// Twelve-Factor friendly.
//
// It exposes an API similar in spirit to the "flag" package.
package twelve

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	prefix string = ""

	ErrNotSet = errors.New("environment variable not set")
)

// Function SetPrefix allows you to set the environment variable prefix.
//
// For example, if your application's name is "myapp", setting the prefix to
// "myapp" will prefix all environment variable lookups with "MYAPP_".
func SetPrefix(p string) {
	prefix = strings.ToUpper(p)
}

func getenv(name string) (string, error) {
	envar := fmt.Sprintf("%s_%s", prefix, strings.ToUpper(name))
	if v := os.Getenv(envar); v == "" {
		return v, ErrNotSet
	} else {
		return v, nil
	}
}

// Function Bool returns the parsed boolean value of the environment variable
// "name".
func Bool(name string) (bool, error) {
	v, err := getenv(name)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(v)
}

// Function Duration parses the specified environment variable, and attempts
// to coerce it into a value.
func Duration(name string) (d time.Duration, err error) {
	var v string
	if v, err = getenv(name); err != nil {
		return
	}
	d, err = time.ParseDuration(v)
	return
}

func Float64(name string) (f float64, err error) {
	var v string
	if v, err = getenv(name); err != nil {
		return
	}
	f, err = strconv.ParseFloat(v, 64)
	return
}

func Int(name string) (i int, err error) {
	var v string
	if v, err = getenv(name); err != nil {
		return
	}
	i, err = strconv.Atoi(v)
	return
}

func Int64(name string) (i int64, err error) {
	var v string
	if v, err = getenv(name); err != nil {
		return
	}
	i, err = strconv.ParseInt(v, 0, 64)
	return
}

// Function String returns the value associated with the environment variable
// e, as a string.
func String(name string) (s string, err error) {
	return getenv(name)
}

func UInt(name string) (u uint, err error) {
	var v string
	if v, err = getenv(name); err != nil {
		return
	}
	var i uint64
	i, err = strconv.ParseUint(v, 0, 0)
	u = uint(i)
	return
}

func UInt64(name string) (u uint64, err error) {
	var v string
	if v, err = getenv(name); err != nil {
		return
	}
	u, err = strconv.ParseUint(v, 0, 64)
	return
}

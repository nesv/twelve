// Package twelve provides utilities to help make your applications more
// Twelve-Factor friendly.
//
// It exposes an API similar to that of the "flag" package.
package twelve

import (
        "fmt"
        "strings"
        "os"
)

var (
        prefix string = ""
)

// Function SetPrefix allows you to set the environment variable prefix.
//
// For example, if your application's name is "myapp", setting the prefix to
// "myapp" will prefix all environment variable lookups with "MYAPP_".
func SetPrefix(p string) {
        prefix = strings.ToUpper(p)
}

func getenv(name string) string {
        return os.Getenv(fmt.Sprintf("%s_%s", prefix, strings.ToUpper(name)))
}

// Function String returns the value associated with the environment variable
// e, as a string.
//
// This function will always return a nil error, and maintains this return
// signature simply for consistency.
func String(name string) *string {
        v := getenv(name)
        return &v
}

func StringVar(p *string, name string) {
        v := getenv(name)
        p = &v
}

func Bool(name string) *bool {
        v := false
        return &v
}

func BoolVar(p *bool, name string) {
        v := false
        p = &v
        return
}
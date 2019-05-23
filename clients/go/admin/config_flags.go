// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-04-12 16:49:03.214864204 -0700 PDT m=+7.007166926

package admin

import (
	"fmt"

	"github.com/spf13/pflag"
)

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "endpoint"), "", "For admin types,  specify where the uri of the service is located.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "insecure"), *new(bool), "Use insecure connection.")
	return cmdFlags
}
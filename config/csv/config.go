package csv

import (
	qfio "github.com/tobgu/qframe/internal/io"
	"github.com/tobgu/qframe/types"
)

// Config holds configuration for reading CSV files into QFrames.
// It should be considered a private implementation detail and should never be
// referenced or used directly outside of the QFrame code. To manipulate it
// use the functions returning ConfigFunc below.
type Config qfio.CsvConfig

// ConfigFunc is a function that operates on a Config object.
type ConfigFunc func(*Config)

// NewConfig creates a new Config object.
// This function should never be called from outside QFrame.
func NewConfig(ff []ConfigFunc) Config {
	conf := Config{Delimiter: ','}
	for _, f := range ff {
		f(&conf)
	}
	return conf
}

// EmptyNull configures if empty strings should be considered as empty strings (default) or null.
//
// emptyNull - If set to true empty string will be translated to null.
func EmptyNull(emptyNull bool) ConfigFunc {
	return func(c *Config) {
		c.EmptyNull = emptyNull
	}
}

// IgnoreEmptyLines configures if a line without any characters should be ignored or interpreted
// as a zero length string.
//
// ignoreEmptyLines - If set to true empty lines will not produce any data.
func IgnoreEmptyLines(ignoreEmptyLines bool) ConfigFunc {
	return func(c *Config) {
		c.IgnoreEmptyLines = ignoreEmptyLines
	}
}

// Delimiter configures the delimiter/separator between columns.
// Only byte representable delimiters are supported. Default is ','.
//
// delimiter - The delimiter to use.
func Delimiter(delimiter byte) ConfigFunc {
	return func(c *Config) {
		c.Delimiter = delimiter
	}
}

// Types is used set types for certain columns.
// If types are not given a best effort attempt will be done to auto detected the type.
//
// typs - map column name -> type name. For a list of type names see package qframe/types.
func Types(typs map[string]string) ConfigFunc {
	return func(c *Config) {
		c.Types = make(map[string]types.DataType, len(typs))
		for k, v := range typs {
			c.Types[k] = types.DataType(v)
		}
	}
}

// EnumValues is used to list the possible values and internal order of these values for an enum column.
//
// values - map column name -> list of valid values.
//
// Enum columns that do not specify the values are automatically assigned values based on the content
// of the column. The ordering between these values is undefined. It hence doesn't make much sense to
// sort a QFrame on an enum column unless the ordering has been specified.
//
// Note that the column must be listed as having an enum type (using Types above) for this option to take effect.
func EnumValues(values map[string][]string) ConfigFunc {
	return func(c *Config) {
		c.EnumVals = make(map[string][]string)
		for k, v := range values {
			c.EnumVals[k] = v
		}
	}
}

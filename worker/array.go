package worker

import (
	"encoding/json"
	"fmt"

	"github.com/opensvc/oc3/feeder"
)

type (
	// arrayRows are the rows an array report makes, in the units of the
	// tables: sizes in MB.
	arrayRows struct {
		// Name is the name of the array in stor_array, and the
		// disk_arrayid of its disks.
		Name     string
		Model    string
		Cache    int64
		Firmware string

		DGs    []arrayDG
		TgtIDs []string
		Disks  []arrayDisk
	}

	arrayDG struct {
		Name string
		Free int64
		Used int64
		Size int64
	}

	arrayDisk struct {
		ID    string
		Name  string
		DevID string
		Size  int64
		Alloc int64
		Raid  string
		Group string
	}

	// arrayParser makes the rows of the array name from the sections of its
	// report.
	arrayParser func(name string, data map[string]json.RawMessage) (*arrayRows, error)
)

// arrayParsers are the parsers by array type. A type absent has its reports
// refused by the feeder: see feeder.ArrayRequiredSections.
var arrayParsers = map[feeder.ArrayType]arrayParser{
	feeder.Pure: parsePure,
}

// arraySection decodes the section of a report into v.
func arraySection(data map[string]json.RawMessage, section string, v any) error {
	b, ok := data[section]
	if !ok {
		return fmt.Errorf("missing section %s", section)
	}
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("section %s: %w", section, err)
	}
	return nil
}

// mb returns a size in bytes in MB, rounded down, as the old collector
// computed it.
func mb(b int64) int64 {
	return b / 1024 / 1024
}

package GraphBlender

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rocketlaunchr/dataframe-go/imports"
	"os"
	"path/filepath"
)

type AutoLoader struct {
	DataFrames DataFrameCollection
}

func (al *AutoLoader) Run(dir string) error {
	csvFiles, err := filepath.Glob(dir + "/*.csv")
	if err != nil {
		return err
	}

	for _, filePath := range csvFiles {
		file, err := os.Open(filePath)
		if err != nil {
			return errors.Wrapf(err, "failed to open file: %s", filePath)
		}

		df, err := imports.LoadFromCSV(context.Background(), file, imports.CSVLoadOptions{})
		if err != nil {
			return errors.Wrapf(err, "failed to parse csv: %s", filePath)
		}

		al.DataFrames = append(al.DataFrames, DataFrameEntry{
			Name:      file.Name(),
			DataFrame: df,
		})
	}

	return nil
}

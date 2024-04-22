package reader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/kidnation-pd/hw02_fix_app/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bt, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal(bt, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

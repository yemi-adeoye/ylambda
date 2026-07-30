package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"ylambda/internal/model"
)

type FileName string

func Read(filename FileName) (*model.FunctionJsonDef, error) {
	buffer := make([]byte, 1024)
	var out strings.Builder
	var errResponse error

	reader, err := os.Open(string(filename))

	if err != nil {
		return nil, err
	}

	defer func(reader *os.File) {
		err := reader.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(reader)

	for {
		bytesRead, err := reader.Read(buffer)
		out.WriteString(string(buffer[:bytesRead]))

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}
	}

	functionJsonDef := model.FunctionJsonDef{}
	err = json.Unmarshal([]byte(out.String()), &functionJsonDef)

	if err != nil {
		return nil, err
	}

	return &functionJsonDef, errResponse
}

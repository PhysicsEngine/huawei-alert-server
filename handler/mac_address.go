package handler

import (
	"strings"
	"os"
	"bufio"
	"io"
	"go.uber.org/zap"
)

type MacAddrHandler struct {
	name      string
	addresses []string
}

func match(handler *MacAddrHandler, target string) bool {
	for _, addr := range handler.addresses {
		if strings.Contains(target, addr) {
			return true
		}
	}
	return false
}

func createHandler(logger *zap.SugaredLogger, name string, fileName string) (*MacAddrHandler, error) {
	logger.Infow("read file::%s", fileName)
	fp, err := os.Open(fileName) // For read access.
	defer fp.Close()

	if err != nil {
		logger.Errorf("read file::%s", fileName)
		return nil, err
	}
	reader := bufio.NewReaderSize(fp, 4096)
	var addresses []string
	for {
		line, _, err := reader.ReadLine()
		logger.Debugf("read line:: %s", line)
		addresses = append(addresses, string(line))
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return &MacAddrHandler{name, addresses}, nil
}

func createHuawaiHandler(logger *zap.SugaredLogger) (*MacAddrHandler, error){
	return createHandler(logger, "hoawai", "./hoawai.txt")
}

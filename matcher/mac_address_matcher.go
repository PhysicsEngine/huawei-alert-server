package matcher

import (
	"bufio"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type MacAddrHandler struct {
	name     string
	prefixes []string
	logger   *zap.SugaredLogger
}

func (handler *MacAddrHandler) Match(target string) bool {
	for _, addr := range handler.prefixes {
		if strings.Contains(target, addr) {
			handler.logger.Infof("Match target::%s in prefix::%s", target, addr)
			return true
		}
	}
	handler.logger.Infof("Match target::%s in any prefix", target)
	return false
}

func createMatcher(logger *zap.SugaredLogger, name string, fileName string) (*MacAddrHandler, error) {
	logger.Infof("read file::%s", fileName)
	fp, err := os.Open(fileName) // For read access.
	defer fp.Close()

	if err != nil {
		logger.Errorf("read file::%s", fileName)
		return nil, err
	}
	var addresses []string
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		addresses = append(addresses, line)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return &MacAddrHandler{name, addresses, logger}, nil
}

func CreateHuaweiMatcher(logger *zap.SugaredLogger) (*MacAddrHandler, error) {
	fileNmae, err := filepath.Abs("./huawei.txt")
	if err != nil {
		return nil, err
	}
	return createMatcher(logger, "huawai", fileNmae)
}

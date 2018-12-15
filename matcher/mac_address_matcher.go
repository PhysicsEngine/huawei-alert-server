package matcher

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type MacAddrMatcher struct {
	name     string
	prefixes []string
	logger   *zap.SugaredLogger
}

func (matcher *MacAddrMatcher) Match(target string) bool {
	for _, addr := range matcher.prefixes {
		if strings.Contains(target, addr) {
			matcher.logger.Infof("Match target::%s in prefix::%s", target, addr)
			return true
		}
	}
	matcher.logger.Infof("Does not match target::%s in any prefix", target)
	return false
}

func createMatcher(logger *zap.SugaredLogger, name string, fileName string) (*MacAddrMatcher, error) {
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

	return &MacAddrMatcher{name, addresses, logger}, nil
}

func CreateHuaweiMatcher(logger *zap.SugaredLogger, basedir string) (*MacAddrMatcher, error) {
	fileName, err := filepath.Abs(fmt.Sprintf("./%s/huawei.txt", basedir))
	if err != nil {
		return nil, err
	}
	return createMatcher(logger, "huawai", fileName)
}

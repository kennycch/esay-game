package log

import (
	"fmt"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

const (
	fileMaxage   = 7 * 24 * time.Hour
	rotationTime = 24 * time.Hour
)

func newFile(directory string) *rotatelogs.RotateLogs {
	d, err := time.ParseDuration("720h")
	if err != nil {
		panic(err)
	} else if d < fileMaxage {
		panic(fmt.Errorf("日志最少要保存7天，当前配置为：%d", d))
	}

	if err = os.MkdirAll(directory, os.ModeDir|0755); err != nil {
		panic(err)
	}

	filename := "%Y%m%d" + ".log"

	w, err := rotatelogs.New(
		path.Join(directory, filename),
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithMaxAge(d),
	)
	if err != nil {
		panic(err)
	}

	return w
}

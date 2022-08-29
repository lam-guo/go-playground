package playground

import (
	"testing"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func TestXxx(t *testing.T) {
	err := errors.New("123")
	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true, // 如果不为true,无法换行
	})
	log.Errorf("%+v\n", err)
}

package cli

import (
	"errors"
	"flag"

	"github.com/sirupsen/logrus"
)

var addMessage string
var listLast bool
var listAll bool

const (
	addUsage  = "Add a block to the blockchain"
	listUsage = "List blocks in the blockchain"
)

var errEmptyAddMessage = errors.New("add command should have a non-empty message")

func Parse() (*string, *bool, *bool, error) {
	flag.StringVar(&addMessage, "add", "", addUsage)
	flag.StringVar(&addMessage, "a", "", addUsage)

	flag.BoolVar(&listLast, "list-last", true, listUsage)
	flag.BoolVar(&listLast, "ll", true, listUsage)

	flag.BoolVar(&listAll, "list-all", false, listUsage)
	flag.BoolVar(&listAll, "la", false, listUsage)

	flag.Parse()

	return &addMessage, &listLast, &listAll, nil
}

func PrintLine() {
	logrus.Info("=== === === ===")
}

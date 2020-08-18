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
	addUsage      = "Add a block to the blockchain"
	fullListUsage = "List blocks in the blockchain"
	lastListUsage = "Shows last block in the blockchain"
)

var errEmptyAddMessage = errors.New("add command should have a non-empty message")

func Parse() (*string, *bool, *bool, error) {
	flag.StringVar(&addMessage, "add", "", addUsage)
	flag.StringVar(&addMessage, "a", "", addUsage)

	flag.BoolVar(&listLast, "list-last", true, lastListUsage)
	flag.BoolVar(&listLast, "ll", true, lastListUsage)

	flag.BoolVar(&listAll, "list-all", false, fullListUsage)
	flag.BoolVar(&listAll, "la", false, fullListUsage)

	flag.Parse()

	return &addMessage, &listLast, &listAll, nil
}

func PrintLine() {
	logrus.Info("=== === === ===")
}

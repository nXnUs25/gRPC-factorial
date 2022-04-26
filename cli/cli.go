package cli

//package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Port int

type integers struct {
	values []int64
}

func (s *integers) String() string {
	if s != nil {
		return strings.Trim(fmt.Sprint(s.values), "{[]}")
	}
	return ""
}

func (s *integers) Set(v string) error {
	if len(s.values) > 0 {
		return fmt.Errorf("Cannot use [-i|--integers] flag more than once")
	}
	ints := strings.Split(v, ",")
	for _, i := range ints {
		if notOk, err := isNotValidInt(i); notOk || err != nil {
			return err
		}
		i64, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			return err
		}
		s.values = append(s.values, i64)
	}
	return nil
}

func Cmds() []int64 {
	var ints integers
	flag.Var(&ints, "integers", "Numbers to calculate 1,2,3,4,... . (default: 5,20,1000,10000000,999999999999999)")
	flag.Var(&ints, "i", "Numbers to calculate 1,2,3,4,... . (default: 5,20,1000,10000000,999999999999999)")
	var help bool
	flag.BoolVar(&help, "help", false, "Prints usage information.")
	flag.BoolVar(&help, "h", false, "\nPrints usage information.")
	flag.IntVar(&Port, "port", 5100, "Port to connect to Factorial Server.")
	flag.IntVar(&Port, "p", 5100, "Port to connect to Factorial Server.")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if len(flag.Args()) == 0 {
		ints.Set("5,20,1000,10000000,999999999999999")
	}
	return ints.values
}

// return true if wrong number not positive int
func isNotValidInt(i string) (bool, error) {
	num, err := strconv.Atoi(i)
	if num < 0 {
		err = fmt.Errorf("Accepting only positive integer numbers")
	}
	return err != nil, err
}

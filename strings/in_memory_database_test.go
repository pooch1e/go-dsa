package strings

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	gostrings "strings"
)

/*
TestInMemoryDictionary tests solution(s) with the following signature and problem description:

	func RunDBCommand(cmd string) string

Write an in memory database that stores string key-value pairs and supports SET, GET, EXISTS,
and UNSET commands. It should also allow transactions with BEGIN, COMMIT and ROLLBACK commands.

For example:

	GET A // outputs nil
	BEGIN
	SET A 1
	GET A // outputs 1
	COMMIT
	GET A // outputs 1

At the first GET A, nil is returned because it has never been set. The second and third
GET A will output 1 because the value of A was set as 1.

BEGIN, ROLLBACK and COMMIT are referred to as transactions in databases. A transaction is
started by BEGIN. The commands that are followed by a BEGIN are completely ignored if a ROLLBACK
command is given and actually applied only when COMMIT command is given.

For example:

	BEGIN
	SET A 1
	COMMIT
	BEGIN
	SET A 2
	ROLLBACK
	GET A // outputs 1

The output is 1 because SET A 2 was never committed so the value of A remains 1.
*/
var _ = DescribeTable("InMemoryDictionary",
	func(input string, allOutputs string) {
		got := ""
		dbs = make([]map[string]string, 0)
		dbs = append(dbs, make(map[string]string))
		for _, cmd := range gostrings.Split(input, "\n") {
			output := RunDBCommand(cmd)
			if output != "" {
				got += " " + output
			}
		}
		if len(got) > 0 {
			got = got[1:]
		}
		if got != allOutputs {
			Expect(got).To(Equal(allOutputs))
		}
	},
	Entry("InMemoryDictionary #1", "EXISTS A\nSET A 1\nGET A", "false 1"),
	Entry("InMemoryDictionary #2", "EXISTS A\nSET A 1\nGET A\nEXISTS A\nUNSET A\nGET A\nEXISTS A", "false 1 true nil false"),
	Entry("InMemoryDictionary #3", "GET A\nBEGIN\nSET A 1\nGET A\nROLLBACK\nGET A", "nil 1 nil"),
	Entry("InMemoryDictionary #4", "GET A\nBEGIN\nSET A 1\nGET A\nCOMMIT\nGET A", "nil 1 1"),
	Entry("InMemoryDictionary #5", "BEGIN\nSET A 1\nCOMMIT\nBEGIN\nSET A 2\nROLLBACK\nGET A", "1"),
	Entry("InMemoryDictionary #6", "SET A 1\nGET A\nBEGIN\nSET A 2\nGET A\nBEGIN\nUNSET A\nGET A\nCOMMIT\nROLLBACK\nGET A", "1 2 nil 1"),
	Entry("InMemoryDictionary #7", "SET A 2\nGET A\nBEGIN\nSET A 1\nGET A\nCOMMIT\nGET A\nBEGIN\nSET A 2\nGET A\nROLLBACK\nGET A", "2 1 1 2 1"),
)

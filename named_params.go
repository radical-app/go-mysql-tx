package mysql

import (
	"errors"
	"fmt"
	"regexp"
)

const regexwords  = `\:([[:word:]]*)`

func NamedParameters(query string, args map[string]interface{}) (str string, ordArgs []interface{}, err error){

	re := regexp.MustCompile(regexwords)
	// [[14 19 15 19] [25 30 26 30] [37 39 38 39] [42 49 43 49]]
	ix := re.FindAllStringSubmatchIndex(query, -1)
	if len(ix) <= 0 && len(args) > 0 {
		return str, ordArgs, errors.New(fmt.Sprintf("Mismatch on indexed words found on query expected %d found 0" ,len(args)))
	}
	str, err = queryReplaceNamedWithQuestionMark(ix, query)
	if err != nil {
		return str, ordArgs, err
	}
	// [[:supp supp] [:derx derx] [:1 1] [:ciccio ciccio]]
	orderedKeys := re.FindAllStringSubmatch(query, -1)
	if len(ix) != len(orderedKeys) {
		return str, ordArgs, errors.New(fmt.Sprintf("problem on finding strings on query %d != %d", len(args), len(orderedKeys)))
	}

	ordArgs, err = createOrderedArgs(orderedKeys, args)
	if err != nil {
		return str, ordArgs, err
	}

	return str, ordArgs, nil
}

func queryReplaceNamedWithQuestionMark(ix [][]int,  query string) (str string, err error) {
	str = ""
	lastIteration := 0
	for _, block := range ix {
		if len(block) <= 0 {
			return str, errors.New("Problem finding correctly all the :keys")
		}

		first := block[0]
		last := block[len(block)-1]

		str = str + query[lastIteration:first] + "?"
		lastIteration = last
	}
	str = str + query[lastIteration:]
	return str, nil
}

func createOrderedArgs(orderedKeys [][]string, mapArgs map[string]interface{}) (ordArgs []interface{}, err error){

	ordArgs = []interface{}{}
	for _, block := range orderedKeys {
		if len(block) != 2 {
			continue
		}
		key := block[1]
		val, ok := mapArgs[key]
		if !ok {
			return ordArgs, errors.New(fmt.Sprintf("Not found key: `%s` on arguments", val))
		}
		ordArgs = append(ordArgs, val)
	}

	return ordArgs, nil
}
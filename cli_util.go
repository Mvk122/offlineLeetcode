package main

import "strconv"

func ConvertArgsToInts(args []string) ([]int, error) {
    problems := make([]int, len(args))
    for i, problem := range args {
        num, err := strconv.Atoi(problem)
        if err != nil {
            return nil, err
        }
        problems[i] = num
    }
    return problems, nil
}
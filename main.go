package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	for _, file := range getChangedGitFiles(`^pkg/.*\.go$`) {
		println(file)
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, file, nil, parser.AllErrors)
		if err != nil {
			panic(err)
		}
		ast.Print(fset, f)
	}
}

func getChangedGitFiles(regexPattern string) []string {
	gitStatus := exec.Command("git", "status", "--porcelain")
	gitStatusOut, err := gitStatus.Output()
	files := []string{}
	if err != nil {
		panic(err)
	}
	gitStatusOutStr := string(gitStatusOut)
	for _, file := range strings.Split(gitStatusOutStr, "\n") {
		if len(strings.Split(file, " ")) < 2 {
			continue
		}
		fileName := strings.Split(file, " ")[1]
		matched, err := regexp.Match(regexPattern, []byte(fileName))
		if err != nil {
			panic(err)
		}
		if matched {
			files = append(files, fileName)
		}
	}
	return files
}

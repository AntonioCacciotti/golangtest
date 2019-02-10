# golangtest
## current branch is features/cobra
## beta version
=======

go run .\launcher\main.go // run the server

project structure :
  * golangtest/golangtest-cobra contains cobra library with https://github.com/go-resty/resty to call the server
  * launcher starts the server built with gin and load questions and answer from a file
  * model objects
  
TODO:
 * <del>debug http://localhost:8000/v1/quiz/questions?next=1</del>
 * <del>debug cobra command golangtest-cobra questions</del>
 * <del>debug golangtest-cobra answer</del>
 * <del>user story 3 and 4</del>
 * more test with different users
 * a little refactoring
 
 Cobra command sequence:
 * golangtest-cobra addNickname antonio
 * golangtest-cobra golangtest-cobra questions
 * golangtest-cobra questions
 * golangtest-cobra answer 1,1,antonio (questionID,answerId,nickname)

Example sequence of commands:
 * golangtest-cobra addNickname antonio
 * golangtest-cobra questions
 * golangtest-cobra questions 1
 * golangtest-cobra answer 1,2,antonio
 * golangtest-cobra actions result,flavia {"message":"Total correct answers:33.000000"} TODO: should 33%
 * golangtest-cobra actions ranking example list {"1":["andrea","flavia"],"3":["antonio"]} Andrea and Flavia score only 1 point because they gave two wrong questions, antonio score 3/3 question

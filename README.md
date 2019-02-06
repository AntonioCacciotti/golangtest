# golangtest
## current branch is features/cobra
## beta version
=======

project structure :
  * golangtest/golangtest-cobra contains cobra library with https://github.com/go-resty/resty to call the server
  * launcher starts the server built with gin and load questions and answer from a file
  * model objects
  
TODO:
 * <del>debug http://localhost:8000/v1/quiz/questions?next=1</del>
 * debug cobra command golangtest-cobra questions
 * debug golangtest-cobra answer
 * user story 3 and 4
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
 * golangtest-cobra actions result,antonio
 * golangtest-cobra actions ranking

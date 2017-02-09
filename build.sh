GOOS=windows GOARCH=386 go build -o ./social-todo-selenium-tests.exe
mv ./social-todo-selenium-tests.exe tmp/windows
GOOS=linux GOARCH=386 go build -o ./social-todo-selenium-tests
mv ./social-todo-selenium-tests tmp/linux
go build -o social-todo-selenium-tests
mv ./social-todo-selenium-tests tmp/mac

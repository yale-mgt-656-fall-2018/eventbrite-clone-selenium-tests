GOOS=windows GOARCH=386 go build -o ./eventbrite-clone-selenium-tests.exe
mv ./eventbrite-clone-selenium-tests.exe tmp/windows
GOOS=linux GOARCH=386 go build -o ./eventbrite-clone-selenium-tests
mv ./eventbrite-clone-selenium-tests tmp/linux
go build -o eventbrite-clone-selenium-tests
mv ./eventbrite-clone-selenium-tests tmp/mac

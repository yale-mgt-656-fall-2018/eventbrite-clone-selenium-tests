#git tag $1

GOOS=windows GOARCH=386 go build -o ./eventbrite-clone-selenium-tests-windows-$1.exe
mv ./eventbrite-clone-selenium-tests-windows-$1.exe releases

GOOS=linux GOARCH=386 go build -o ./eventbrite-clone-selenium-tests-linux-$1
mv ./eventbrite-clone-selenium-tests-linux-$1 releases

go build -o eventbrite-clone-selenium-tests-mac-$1
mv ./eventbrite-clone-selenium-tests-mac-$1 releases

#aws s3 sync --acl public-read releases s3://files.656.mba --exclude=".git"

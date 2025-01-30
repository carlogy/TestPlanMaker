run:
	go run .

compileRelease:

	go build -o BuildTestPlan$(ver)

	GOOS=windows GOARCH=amd64 go build -o BuildTestPlan$(ver).exe

compile:

	go build -o BuildTestPlan

run:

	go build -o TestPlanMaker && ./TestPlanMaker

compile:

	go build -o TestPlanMaker

	GOOS=windows GOARCH=amd64 go build -o TestPlanMaker.exe

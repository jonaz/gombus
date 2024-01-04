export CGO_ENABLED=0
GOOS=linux GOARCH=arm GOARM=7 go build -o mbusctl-linux-arm

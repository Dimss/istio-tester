curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh;
cd /builds/dimss/my-test-proj
dep ensure;
export CGO_ENABLED=0 GOOS=linux GOARCH=amd64;
go build -o istio-tester;
ls -all

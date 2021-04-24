source config.env
echo "Testing: $(basename $PWD)"
go run main.go -FLAG_B=overridden_by_cli

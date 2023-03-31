# use all folders in cmd instead of search
install:
	find . -name main.go -execdir go install . \;

uninstall:
	find . -name main.go -execdir rm $$GOPATH/bin/{} . \;

command-runner:
	find . -name main.go -execdir 
docs: godocdown
	for pkg in $$(find . -maxdepth 1 -type d -not -name "Godeps" -and -not -name ".git" -and -not -name "." -and -not -name "docs"); \
	do \
		godocdown -output="docs/$$pkg.md" $$pkg; \
	done

godocdown:
	go get -u -v github.com/robertkrimen/godocdown/godocdown

.PHONY: docs

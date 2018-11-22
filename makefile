export VERSION=$(shell gogitver)
ARTIFACT_PATH=./artifacts/hacienda

.PHONY: build buildarm build-debian-package clean package

build:
	go build -o $(ARTIFACT_PATH) cmd/hacienda/main.go

buildarm:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o $(ARTIFACT_PATH) cmd/hacienda/main.go

clean:
	rm -Rf ./artifacts

build-debian-package: clean buildarm
	mkdir -p artifacts/debian/DEBIAN
	mkdir -p artifacts/debian/usr/bin/
	mkdir -p artifacts/debian/lib/systemd/system

	cat ./build/debian/control.template | envsubst > ./artifacts/debian/DEBIAN/control
	cp ./build/debian/postinst ./artifacts/debian/DEBIAN
	cp ./build/debian/prerm ./artifacts/debian/DEBIAN
	cp $(ARTIFACT_PATH) artifacts/debian/usr/bin/hacienda
	cp build/debian/hacienda.service ./artifacts/debian/lib/systemd/system
	dpkg-deb --build ./artifacts/debian ./artifacts/hacienda.deb
	rm -R ./artifacts/debian

package: build-debian-package

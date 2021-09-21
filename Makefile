GO = go

TARGET = midget

# change this if you are using a diferrent os and architecture
# more info in the README.md
ARCH = amd64
OS = linux

MODS_FOLDER = $(HOME)/Funkin/mods
GH_REPO = MatusOllah/midget-pkgs
MIDGETRC = $(HOME)/.midgetrc.toml


all:
	mkdir bin

	$(GO) mod tidy

	GO111MODULE=on GOARCH=$(ARCH) GOOS=$(OS) $(GO) build -o bin/$(TARGET) -v

install:
	install -p -m 0755 bin/$(TARGET) /usr/bin/$(TARGET)

	echo "mods_folder = $(MODS_FOLDER)\ngh_repo = $(GH_REPO)" >> $(MIDGETRC)

uninstall:
	rm /usr/bin/$(TARGET) 

	rm MIDGETRC

clean:
	rm -rf bin
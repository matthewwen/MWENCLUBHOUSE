LIB=-lwebsockets -lcrypto -lsqlite3 -lpython3.8 -lunqlite -lczmq -lzmq
CFLAG=-O6 -std=c99 -Wshadow -Wvla -pedantic -Wall
INCLUDE=-I header -I /usr/include/python3.8

current_dir := $(shell pwd)

service-deploy: 
	rm -rf server.service
	@echo "[Unit] " >> server.service 
	@echo "Description=Server Running Backend for Matthew Wen's mwenclubhouse" >> server.service 
	@echo "[Service]" >> server.service 
	@echo  "Type = simple" >> server.service
	@echo  "ExecStart = ${current_dir}/server dep" >> server.service

build:
	mkdir -p object
	$(CC) -fPIC -c csrc/oauth.c -I cheader -o object/oauth.o -ljwt
	$(CC) -fPIC -c csrc/apipageview.c -I cheader -o object/apipageview.o
	$(CC) -shared -o object/libclubhouse.so object/oauth.o object/apipageview.o -lsqlite3 -ljansson

install:
	make build
	sudo mv object/libclubhouse.so /usr/local/lib/
	sudo ldconfig
	make service-deploy
	sudo cp server.service /etc/systemd/system/
	sudo chmod 644 /etc/systemd/system/server.service

uninstall:
	sudo rm -rf /usr/local/lib/libclubhouse.so /etc/systemd/system/server.service
	ldconfig

clean:
	rm -rf main $(SQLITENAME) object server

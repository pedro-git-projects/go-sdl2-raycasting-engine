OBJ_NAME = 
LDFLAGS = 
install:
	$(eval OBJ_NAME += raycasting)
	$(eval LDFLAGS += "-w -s")
	cd ./cmd/; go build -v -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../bin 
run:
	./bin/raycasting

doc:
	cd ./cmd/; godoc -http=:6060

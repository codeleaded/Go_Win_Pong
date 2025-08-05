GOC = go
SRC_DIR = src
BUILD_DIR = build

TARGET = $(BUILD_DIR)/Main

# @ Removes the Output of execution
all:
	$(GOC) build -o $(TARGET) $(SRC_DIR)/Main.go
	chmod 777 $(TARGET)

exe:
	./$(TARGET)

clean:
	rm -rf $(BUILD_DIR)/*

do: clean all exe
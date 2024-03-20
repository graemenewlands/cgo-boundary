CC=gcc
CFLAGS=-Wall -fPIC
TARGET_LIB_STATIC=pkg/cgoboundary/libtest.a
TARGET_LIB_SHARED=pkg/cgoboundary/libtest.so
OBJ=test.o

all: $(TARGET_LIB_STATIC) $(TARGET_LIB_SHARED)

$(OBJ): pkg/cgoboundary/test.c pkg/cgoboundary/test.h
	$(CC) $(CFLAGS) -c pkg/cgoboundary/test.c

$(TARGET_LIB_STATIC): $(OBJ)
	ar rcs $(TARGET_LIB_STATIC) $(OBJ)

$(TARGET_LIB_SHARED): $(OBJ)
	$(CC) -shared -o $(TARGET_LIB_SHARED) $(OBJ)

clean:
	rm -f $(TARGET_LIB_STATIC) $(TARGET_LIB_SHARED) $(OBJ)

.PHONY: all clean

KDIR := /lib/modules/$(shell uname -r)/build
PATH_TO_MAIN := $(shell dirname $(realpath $(CURDIR)))/main

obj-m := hello.o

default:
	make -C $(KDIR) M=$(CURDIR) src=$(CURDIR) modules EXTRA_CFLAGS='-DPATH_TO_MAIN=\"$(PATH_TO_MAIN)\"'

clean:
	make -C $(KDIR) M=$(CURDIR) clean

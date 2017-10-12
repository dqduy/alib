from subprocess import call

call(["g++", "main.cpp", "../../lib/ArrayList.cpp", "../../lib/util/ListUtil.cpp", "-omain", "--std=c++11"])

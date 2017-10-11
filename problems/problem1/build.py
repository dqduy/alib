from subprocess import call

call(["g++", "--std=c++11", "main.cpp", "../../lib/ArrayList.cpp", "../../lib/util/ListUtil.cpp"])

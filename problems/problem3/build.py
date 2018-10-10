from subprocess import call

call(["g++", "main.cpp", "tinyxml2.cpp", "-omain", "--std=c++11"])

from subprocess import call

call(["g++", "CodePoint.cpp", "main.cpp", "-omain", "--std=c++11"])

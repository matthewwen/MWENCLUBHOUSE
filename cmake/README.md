# MWENCLUBHOUSE C Lang SDK
The mwenclubhouse refers to the SDK for the server. There are 3 parts to the sdk for mwenclubhouse. The first part is GoLang SDK that communicates directly with the GoLang Application. The second part is an sdk written in C, which uses CGO to communicate with both the sdk and the application. This separates the sdk depending on the programming language. The third part of the SDK an SDK written in Python. 

## mwenclubhouse-clang
CGO does allow programmers to directly put their C Program into GoLang as an imported statement. However, it isn't efficient. As a result, mwenclubhouse-clang creates a library on linux machine, and the sdk refers to that library to compile and run. Note, that the mwenclubhouse-clang program must be installed on the computer before using the golang sdk.

### Part 1: CMAKE
The program uses CMAKE to build the library. CMakeList.txt is the first iteration of the effort.
from os import system
from sys import argv, stderr
from platform import system as os_name

def build():
    system("mkdir -p build")
    system("git clone https://github.com/nix-enthusiast/unildd.git build/unildd")
    system("cargo build --release --manifest-path=build/unildd/Cargo.toml")
    system("mkdir lib")
    system("mkdir include")

    # I know this logic is dodgy
    match os_name():
        case "Windows":
            system("cp build\\unildd\\target\\release\\libunildd.dll lib")
        case "Darwin":
            system("cp build/unildd/target/release/libunildd.dylib lib")
        case _:
            system("cp build/unildd/target/release/libunildd.so lib")

    system("cp build/unildd/header/unildd.h include")


match argv[1]:
    case "--build" | "-b":
        build()
        system("go build -o build/objdetect")

    case "--run" | "-r":
        build()
        system("go run main.go " + ' '.join(argv[2:]))

    case "--clean" | "-c":
        system("rm -rf build")

    case f:
        print("Invalid flag '" + f + "'", file=stderr)

# OBJDetect

OBJDetect is a program made in Go as an example of how to use [unildd](https://github.com/nix-enthusiast/unildd) library

## Requirements
- Python >= 3.10
- Cargo
- Go >= 1.17

## Installation
- Clone this git repository:
  
  ```git clone https://github.com/nix-enthusiast/objdetect```

- Go into the repository:

  ```cd objdetect```

- Run the `build.py` file:

  ```
  python3 build.py --build   # To build the program
  python3 build.py --run     # To directly run the program 
  python3 build.py --clean   # To remove the build directory
  ```

- Take the file named `objdetect` (or `objdetect.exe` in Windows) from the directory named `build` and put it anywhere you want!

## ⚠️ A Small Warning

Since it uses [unildd](https://github.com/nix-enthusiast/unildd) library, the library has to be accessible by the program. To do it on

### Windows:
  
  You can put the library in any folder which is in the `%PATH` variable or put them in the same place

  ```
  cp build\target\release\unildd.dll \the\folder\in\the\path\var
  #or
  cp build\target\release\unildd.dll \the\folder\which\includes\objdetect
  ```

### Linux and macOS:
  The same thing as what we do in Windows but the variable is:
  
  - `LD_LIBRARY_PATH` for Linux[^1]
  - `DYLD_LIBRARY_PATH` for macOS[^2]


  ```
  cp build/target/release/unildd.dll /the/folder/in/the/variable
  #or
  cp build/target/release/unildd.dll /the/folder/which/includes/objdetect
  ```

### Other OSes
  If your OS is not listed on here, please take a look at the documentation of your OS to find the path and do the same (or similar since OSes work different) thing as what we did above 

## License
This library is licensed under [BSD-3 Clause License](https://choosealicense.com/licenses/bsd-3-clause/)

[^1]: https://man7.org/linux/man-pages/man8/ld.so.8.html#:~:text=LD_LIBRARY_PATH%0A%20%20%20%20%20%20%20%20%20%20%20%20%20%20A%20list%20of%20directories%20in%20which%20to%20search%20for%20ELF%20libraries%0A%20%20%20%20%20%20%20%20%20%20%20%20%20%20at%20execution%20time.
[^2]: https://developer.apple.com/library/archive/documentation/DeveloperTools/Conceptual/DynamicLibraries/100-Articles/UsingDynamicLibraries.html#:~:text=You%20may%20also,DYLD_FALLBACK_LIBRARY_PATH

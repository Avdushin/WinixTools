### WINIXTOOLS

#### List of unix utilities for Windows

![logo](https://i.pinimg.com/564x/b0/16/be/b016bece7f010a7483548d36f43ddf60.jpg)

#### Utilities

| NAME        | Version    |  Tested     |
| ----------- | -----------| ----- |
| **ls**      | **1.0**    |  :heavy_check_mark:
| **ll**      | **1.0**    |  :heavy_check_mark:
| **cat**     | **1.0**    |  :heavy_check_mark:
| **clear**   | **1.0**    |  :heavy_check_mark:
| **[make](https://www.gnu.org/software/make/)**   | **1.0**    |  :heavy_check_mark:

### How to install

1) Open CMD at the **application folder**

2) Put this script to CMD
```powershell
mkdir %AppData%\WinixTools 
copy Tools\* %AppData%\WinixTools
SET path=%AppData%\WinixTools;%path%
```

or start `start.bat` file 

3) **CMD => Always run as Administrator**
4) Changing System Environment Variables &rarr; Environment settings &rarr; System variables &rarr; Path &rarr; Edit... &rarr; Create &rarr; `%AppData%\WinixTools`

   
  P.s `%AppData%` it's = `C:\Users\%username%\AppData\Roaming`

**YOUR_USER** is your system username

1) Restart CMD
2) All ready to go!

### How to use

* `ls` &rarr; **show files in the directory**
* `ll` &rarr; **show files in the directory & their rights**
* `cat` &rarr; **read file from CMD**
* `clear` &rarr; **clear CMD screen**
* `make` &rarr; **start Makefile scripts**

<p align="center">2022 © <a href="https://github.com/Avdushin" target="_blank">AVDUSHIN</a></p>
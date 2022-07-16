### WINIXTOOLS

#### List of unix utilities for Windows

#### Utilities

| NAME        | Version    |  Tested     |
| ----------- | -----------| ----- |
| **ls**      | **0.1**    |  :heavy_check_mark:
| **cat**     | **0.1**    |  :heavy_check_mark:
| **clear**   | **0.1**    |  :heavy_check_mark:

### How to install

1) Open CMD at the ***application folder***

2) Put this script to CMD
```powershell
copy Tools\* %homepath%
SET path=C:%homepath%\ls.exe;%path%
SET path=C:%homepath%\clear.exe;%path%
SET path=C:%homepath%\cat.exe;%path%
```

or start `start.bat` file 

3) **CMD => Always run as Administrator**
4) Changing System Environment Variables &rarr; Environment settings &rarr; System variables &rarr; Path &rarr; Edit... &rarr; Create &rarr; `C:\User\YOUR_USER`

**YOUR_USER** is your system username

5) Restart CMD
6) All ready to go!

### How to use

* `ls` &rarr; **show files in the directory**
* `cat` &rarr; **read file from CMD**
* `clear` &rarr; **clear CMD screen**


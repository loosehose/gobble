# gobble
<p align="center">
  <img src="https://user-images.githubusercontent.com/75705022/212232137-4dd329e8-fab6-4cc9-94d9-a5bd32ca8588.png" />
</p>

Gobble is a combination of Golang projects, featuring the skeleton of Uru and meat consisting of either original content or inspiration from Tim Whitez's Doge-Gabh. To learn more about creating malware quickly and effectively, please check out both the Uru and Doge-Gabh repositories.

## Installation

```
git clone https://github.com/loosehose/gobble.git
cd gobble
go build
```



## DIY

Assume your goal is to add injection methods to this project. This can easily be accomplished by using the following steps:

To add injection methods to this project, follow these steps:

1. Go to the data/injector/windows/[bananaphone/native]/local/ directory.

2. Create a new folder with the name of the malware (e.g. EarlyBird).

3. Within the new folder, create two files: functions.go.tmpl and instanciation.go.tmpl. The functions.go.tmpl file will contain the actual malware code, while the instanciation.go.tmpl file specifies where the code execution begins.

4. Once the malware code is complete, navigate to pkg/injector/[bananaphone/native]/ and create a new file with the same name as the malware folder (e.g. EarlyBird.go).

5. Populate this file by copying and modifying an existing template in this directory, including changing function names and importing necessary packages.

6. In the pkg folder, locate the injectorFactory.go file and add the following code:

   ```
   if injectorType == "windows/native/local/EarlyBird" {
   		return native.NewEarlyBird(), nil
   }
   ```

7. Save all changes and rebuild the project using `go build` before testing the newly added code.

## Common Errors

To solve issues on non-Windows operating systems, use the following syntax: 

`GOOS=windows ./gobble generate ...` 

If a package is causing problems, such as this error message: 

`FTL Error during build: error exit status 1: go list error: exit status 1: go build github.com/timwhitez/Doge-Gabh/pkg/Gabh: no Go files in` 

A possible solution is to manually copy the generated file from the output message: `INF Payload file has been written path=/path/to/file.go` and then build it using the following command: `GOOS=windows go build /path/to/file.go` 

This workaround can help to build the package successfully.
